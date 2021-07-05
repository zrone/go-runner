package main

import (
	"bufio"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"log"
	"sync"
	"time"
)

func main() {
	sshHost := "127.0.0.1"
	sshUser := "zrone"
	sshPassword := "bluestone"
	sshType := "publickey"                   //password 或者 key
	sshKeyPath := "/Users/zrone/.ssh/id_rsa" //ssh id_rsa.id 路径"
	sshPort := 22

	//创建sshp登陆配置
	config := &ssh.ClientConfig{
		Timeout:         time.Second, //ssh 连接time out 时间一秒钟, 如果ssh验证错误 会在一秒内返回
		User:            sshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以， 但是不够安全
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}
	if sshType == "password" {
		config.Auth = []ssh.AuthMethod{ssh.Password(sshPassword)}
	} else {
		config.Auth = []ssh.AuthMethod{publicKeyAuthFunc1(sshKeyPath)}
	}

	//dial 获取ssh client
	addr := fmt.Sprintf("%s:%d", sshHost, sshPort)
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatal("创建ssh client 失败", err)
	}
	defer sshClient.Close()

	//创建ssh-session
	session, err := sshClient.NewSession()
	if err != nil {
		log.Fatal("创建ssh session 失败", err)
	}
	defer session.Close()
	////执行远程命令
	//combo,err := session.CombinedOutput("ping www.baidu.com")
	//if err != nil {
	//	logr.Fatal("远程执行cmd 失败",err)
	//}
	//
	//logr.Println("命令输出:",string(combo))

	stdout, err := session.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		reader := bufio.NewReader(stdout)
		for {
			readString, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				return
			}
			fmt.Print(readString)
		}
	}()

	err = session.Run("cd /Users/zrone && ls -la && ping www.baidu.com")
	wg.Wait()

	if err != nil {
		log.Fatal("远程执行cmd 失败", err)
	}
}

func publicKeyAuthFunc1(kPath string) ssh.AuthMethod {
	keyPath, err := homedir.Expand(kPath)
	if err != nil {
		log.Fatal("find key's home dir failed", err)
	}
	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		log.Fatal("ssh key file read failed", err)
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatal("ssh key signer failed", err)
	}
	return ssh.PublicKeys(signer)
}
