package ssh

import (
	"bufio"
	"errors"
	"github.com/mitchellh/go-homedir"
	taskLogrus "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

func Send(sshClient *ssh.Client, direct string, uuid string) error {
	taskLog, _ := os.OpenFile("runtime/task/"+uuid+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer taskLog.Close()

	taskLogrus.SetOutput(taskLog)
	taskLogrus.Println(direct)

	//创建ssh-session
	session, err := sshClient.NewSession()
	if err != nil {
		taskLogrus.Errorln("创建ssh session 失败", err)
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
			taskLogrus.Println(readString)
			// fmt.Print(readString)
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
