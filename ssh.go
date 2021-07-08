package main

import (
	"awesome-runner/src/logr"
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"time"
)

func main() {
	//创建sshp登陆配置
	config := &ssh.ClientConfig{
		Timeout:         time.Second,
		User:            "root",
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	config.Auth = []ssh.AuthMethod{ssh.Password("123456")}

	//dial 获取ssh client
	addr := fmt.Sprintf("%s:%d", "10.211.55.3", 22)
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		fmt.Printf("创建ssh client 失败", err)
	}
	defer sshClient.Close()

	//创建ssh-session
	session, err := sshClient.NewSession()
	if err != nil {
		logr.Clog.Errorf("创建ssh session 失败", err)
	}
	defer session.Close()

	//err = session.Setenv("WORKDIR", "demo")
	//if err != nil {
	//	fmt.Println(err)
	//}

	stdout, _ := session.StdoutPipe()
	go func() {
		reader := bufio.NewReader(stdout)
		for {
			readString, _ := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				return
			}
			fmt.Println(readString)
		}
	}()

	stdinBuf, _ := session.StdinPipe()

	session.Shell()
	stdinBuf.Write([]byte("WORKDIR=a\n"))
	stdinBuf.Write([]byte("echo $WORKDIR\n"))

	defer session.Close()
	fmt.Println("ok")
}
