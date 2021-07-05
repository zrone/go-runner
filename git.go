package main

import (
	conf "awesome-runner/src/config"
	"awesome-runner/src/logr"
	"awesome-runner/src/sql"
	interactive "awesome-runner/src/ssh"
	"awesome-runner/types"
	"bufio"
	"errors"
	"fmt"
	examples "github.com/go-git/go-git/v5/_examples"
	"github.com/golang-module/carbon"
	jsoniter "github.com/json-iterator/go"
	"github.com/mitchellh/go-homedir"
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

//func main() {
//	branch := "origin master"
//	var internalDeloy types.InternalDeploy
//	sql.GetLiteInstance().First(&internalDeloy, "symbol = ?", "demo")
//
//	examples.Info(fmt.Sprintf("git pull %s", branch))
//
//	_, err := git.PlainClone("/tmp/foo", false, &git.CloneOptions{
//		URL:      "https://gitee.com/marksirl/demo",
//		Progress: os.Stdout,
//	})
//
//	examples.CheckIfError(err)
//}

func main() {
	JSON := jsoniter.ConfigCompatibleWithStandardLibrary

	closeChan := make(chan byte)

	time.AfterFunc(time.Second*5, func() {
		fmt.Println("after")
		close(closeChan)
	})

	//for {
	//	select {
	//	// 读等待
	//	case <-closeChan:
	//		fmt.Println("demo")
	//		return
	//	default:
	//	}
	//}

	var (
		wg  sync.WaitGroup
		cmd = exec.Command("/bin/bash", "-c", fmt.Sprintf("tail -n +0 -f runtime/task/418423262502981632.log"))
	)

	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	wg.Add(1)
	go func() {
		defer wg.Done()
		reader := bufio.NewReader(stdout)
		for {
			select {
			// 读等待
			case <-closeChan:
				fmt.Println("demo")
				return
			default:
			}

			readString, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				return
			}
			var msg types.LogFormat
			err = JSON.Unmarshal([]byte(readString), &msg)
			if err != nil {
				return
			}
			fmt.Println(msg)
		}
	}()

	wg.Wait()
	cmd.Wait()
}

func main12() {
	var internalDeloy types.InternalDeploy
	sql.GetLiteInstance().First(&internalDeloy, "symbol = ?", "demo")

	// 链接ssh
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
		config.Auth = []ssh.AuthMethod{publicKeyAuthFunc(sshKeyPath)}
	}

	//dial 获取ssh client
	addr := fmt.Sprintf("%s:%d", sshHost, sshPort)
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatal("创建ssh client 失败", err)
	}
	defer sshClient.Close()

	snowflake := carbon.Now().Format("Ymd") + logr.SnowFlakeId()
	tempDir := `~/.runner/demo/` + snowflake
	if err = interactive.Send(sshClient, fmt.Sprintf(`rm -rf %s && mkdir -p %s && cd %s && git init && git remote add origin %s && git config core.sparsecheckout true && echo .runner-ci.yml >> .git/info/sparse-checkout && git pull origin master`, tempDir, tempDir, tempDir, "git@gitee.com:marksirl/demo.git"), ""); err != nil {
		fmt.Println(err)
		return
	}

	dir, _ := os.Getwd()
	cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("scp %s@%s:%s/.runner-ci.yml %s/runtime/%s.runner-ci.yml", "zrone", "localhost", tempDir, dir, snowflake))
	cmd.Start()
	cmd.Wait()

	var (
		runnerCi types.RunnerCi
	)

	conf.ParseYaml(fmt.Sprintf("runtime/%s.runner-ci.yml", snowflake), &runnerCi)
	var isContain bool = false
	for _, b := range runnerCi.Only {
		if strings.Compare(b, "branch") == -1 {
			isContain = true
			break
		}
	}
	fmt.Println(isContain)
	// 执行脚本
}

func send(sshClient *ssh.Client, direct string) error {
	fmt.Println(direct)
	//创建ssh-session
	session, err := sshClient.NewSession()
	if err != nil {
		//log.Fatal("创建ssh session 失败", err)
		return errors.New("创建ssh session 失败")
	}
	defer session.Close()

	stdout, err := session.StdoutPipe()
	if err != nil {
		return err
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

	// 打印
	err = session.Run(direct)
	wg.Wait()
	if err != nil {
		return err
	}
	defer session.Close()
	return nil
}

func main1() {
	branch := "origin master"
	direct := fmt.Sprintf("git pull %s", branch)

	var internalDeloy types.InternalDeploy
	sql.GetLiteInstance().First(&internalDeloy, "symbol = ?", "demo")

	// 链接ssh
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
		config.Auth = []ssh.AuthMethod{publicKeyAuthFunc(sshKeyPath)}
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

	// 打印
	//examples.Info(fmt.Sprintf("cd %s", internalDeloy.Path))
	examples.Info(fmt.Sprintf(`cd %s && %s`, internalDeloy.Path, direct))
	err = session.Run(fmt.Sprintf(`cd %s && pwd && %s`, internalDeloy.Path, direct))
	//err = session.Run(direct)
	wg.Wait()
	if err != nil {
		fmt.Println(err)
	}

	// 解析ci文件
	// 执行脚本
}

func publicKeyAuthFunc(kPath string) ssh.AuthMethod {
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
