package task

import (
	"awesome-runner/src/logr"
	"awesome-runner/src/sql"
	interactive "awesome-runner/src/ssh"
	"awesome-runner/types"
	"context"
	"errors"
	"fmt"
	taskLogrus "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"os"
	"time"
)

func Deliver(UUID string, Symbol string, Branch string, BeforeScript []string, Script []string, AfterScript []string) error {
	var (
		internalDeloy types.InternalDeploy
		taskRecord    types.TaskLog
	)

	sql.GetLiteInstance().Take(&internalDeloy, "symbol = ?", Symbol)
	if internalDeloy == (types.InternalDeploy{}) {
		logr.Logrus.Errorln("Unknown symbol")
		return errors.New(types.NOTIFICATION_WORK_SERVER)
	}

	sql.GetLiteInstance().Model(&taskRecord).Where("uuid = ?", UUID).Update("State", `RUNNING`)

	// 链接SSH
	sshHost := internalDeloy.Auth.Host
	sshUser := internalDeloy.Auth.User
	sshPassword := internalDeloy.Auth.Pwd
	sshType := "password"       //password 或者 key
	sshKeyPath := "id_rsa path" //ssh id_rsa.id 路径"
	sshPort := internalDeloy.Auth.Port

	//创建sshp登陆配置
	config := &ssh.ClientConfig{
		Timeout:         time.Second,
		User:            sshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	if sshType == "password" {
		config.Auth = []ssh.AuthMethod{ssh.Password(sshPassword)}
	} else {
		config.Auth = []ssh.AuthMethod{interactive.PublicKeyAuthFunc(sshKeyPath)}
	}

	//dial 获取ssh client
	addr := fmt.Sprintf("%s:%d", sshHost, sshPort)
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		taskLogrus.Errorf("创建ssh client 失败", err)
		logr.Logrus.Errorf("创建ssh client 失败", err)
	}
	defer sshClient.Close()

	ctx := context.Background()
	taskLog, _ := os.OpenFile("runtime/task/"+UUID+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer taskLog.Close()
	taskLogrus.SetOutput(taskLog)
	tl := taskLogrus.WithContext(ctx)

	tl.Println("--------- Before script ---------")
	// 执行 before 脚本 *File
	for _, b := range BeforeScript {
		err := interactive.Send(sshClient, b, tl)
		if err != nil {
			sql.GetLiteInstance().Model(&taskRecord).Where("uuid = ?", UUID).Update("State", `FAILURE`)
			return nil
		}
	}
	tl.Println("")
	tl.Println("--------- Deploy script ---------")
	// 执行 main 脚本
	for _, s := range Script {
		err := interactive.Send(sshClient, s, tl)
		if err != nil {
			sql.GetLiteInstance().Model(&taskRecord).Where("uuid = ?", UUID).Update("State", `FAILURE`)
			return nil
		}
	}
	tl.Println("")
	tl.Println("--------- After script  ---------")
	// 执行 after 脚本
	for _, a := range AfterScript {
		err := interactive.Send(sshClient, a, tl)
		if err != nil {
			sql.GetLiteInstance().Model(&taskRecord).Where("uuid = ?", UUID).Update("State", `FAILURE`)
			return nil
		}
	}
	tl.Println("")
	tl.Println("----------- SUCCESS -------------")

	sql.GetLiteInstance().Model(&taskRecord).Where("uuid = ?", UUID).Update("State", `SUCCESS`)
	return nil
}
