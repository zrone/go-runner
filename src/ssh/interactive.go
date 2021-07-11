package ssh

import (
	"awesome-runner/src/logr"
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"sync"

	"github.com/golang-module/carbon"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

func Send(sshClient *ssh.Client, direct string, taskLogrus *logrus.Entry, env map[string]string) error {
	taskLogrus.Debug(fmt.Sprintf(`%s %s`, carbon.Now().ToTimeString(), direct))

	//创建ssh-session
	session, err := sshClient.NewSession()
	if err != nil {
		taskLogrus.Errorf(fmt.Sprintf("%s %s:%v", carbon.Now().ToTimeString(), "创建ssh session 失败", err))
		return errors.New("创建ssh session 失败")
	}
	//session.Setenv()
	var prefix string
	for name, value := range env {
		prefix += fmt.Sprintf("%s=%s && ", name, value)
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

			readString = strings.Replace(readString, prefix, "", -1)
			taskLogrus.Println(fmt.Sprintf(`%s %s`, carbon.Now().ToTimeString(), readString))
			// fmt.Print(readString)
		}
	}()

	// 打印
	err = session.Run(fmt.Sprintf("%s%s", prefix, direct))
	wg.Wait()
	if err != nil {
		taskLogrus.Errorln(fmt.Sprintf(`%s %v`, carbon.Now().ToTimeString(), err))
		taskLogrus.Println("")
		return err
	}
	defer session.Close()
	return nil
}

func InterSend(sshClient *ssh.Client, direct string) error {
	//创建ssh-session
	session, err := sshClient.NewSession()
	if err != nil {
		logr.Clog.Errorf("创建ssh session 失败", err)
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
			logr.Clog.Debug(readString)
			//fmt.Println(readString)
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

func PublicKeyAuthFunc(kPath string) ssh.AuthMethod {
	keyPath, err := homedir.Expand(kPath)
	if err != nil {
		logr.Clog.Errorf("find key's home dir failed", err)
	}
	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		logr.Clog.Errorf("ssh key file read failed", err)
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		logr.Clog.Errorf("ssh key signer failed", err)
	}
	return ssh.PublicKeys(signer)
}
