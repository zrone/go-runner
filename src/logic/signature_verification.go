package logic

import (
	conf "awesome-runner/src/config"
	"awesome-runner/src/logr"
	"awesome-runner/src/queue"
	"awesome-runner/src/sql"
	interactive "awesome-runner/src/ssh"
	"awesome-runner/types"
	"fmt"
	"github.com/RichardKnop/machinery/v2/tasks"
	"github.com/golang-module/carbon"
	"github.com/kataras/iris/v12"
	taskLogrus "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

// 签名验证
func SignatureVerification(ctx iris.Context, crypt types.AbstractCrypt) (int, error) {
	// 签名验证
	crypt.Build()
	crypt.BuildPrefixCryptSign()
	if crypt.Compare() {
		re := regexp.MustCompile(`refs/heads/(.*)`)
		match := re.FindStringSubmatch(crypt.GetCryptDataConfig().Message.Ref)
		var ref string = ""
		if len(match) > 1 {
			ref = match[1]
		}

		logr.Logrus.Printf("当前分支 %s", ref)
		isAllowBranch, _, params := isAllowBranch(crypt.GetCryptDataConfig(), ref)

		if len(match) > 1 && isAllowBranch {
			uuid := logr.SnowFlakeId()
			taskLog, _ := os.Create("runtime/task/" + uuid + ".log")
			defer taskLog.Close()

			taskLogrus.SetOutput(taskLog)

			env, _ := logr.JSON.Marshal(params.Environment)

			eta := time.Now().Add(time.Second * 5)
			args := []tasks.Arg{
				{
					Name:  "UUID",
					Type:  "string",
					Value: uuid,
				},
				{
					Name:  "Symbol",
					Type:  "string",
					Value: crypt.GetCryptDataConfig().Symbol,
				},
				{
					Name:  "Branch",
					Type:  "string",
					Value: match[1],
				},
				{
					Name:  "Env",
					Type:  "string",
					Value: string(env),
				},
				{
					Name:  "BeforeScript",
					Type:  "[]string",
					Value: params.BeforeScript,
				},
				{
					Name:  "Script",
					Type:  "[]string",
					Value: params.Script,
				},
				{
					Name:  "AfterScript",
					Type:  "[]string",
					Value: params.AfterScript,
				},
			}
			// 发送部署任务
			signature := &tasks.Signature{
				UUID: uuid,
				Name: "call",
				Args: args,
				ETA:  &eta,
				// 重试次数和斐波那契间隔
				// RetryCount:   3,
				// RetryTimeout: 15,
			}

			argsString, _ := logr.JSON.Marshal(args)

			if _, err := queue.MachineryServer.SendTask(signature); err != nil {
				sql.GetLiteInstance().Create(&types.TaskLog{
					Symbol:   crypt.GetCryptDataConfig().InternalDeloy.Symbol,
					Uuid:     uuid,
					CreateAt: carbon.Now().Format("Y-m-d H:i:s"),
					State:    "FAILURE",
					Args:     string(argsString),
				})

				taskLogrus.Errorf("Failed task delivered，%v", err)
				logr.Clog.Errorf("Failed task delivered，%v", err)
			} else {
				//
				sql.GetLiteInstance().Create(&types.TaskLog{
					Symbol:   crypt.GetCryptDataConfig().InternalDeloy.Symbol,
					Uuid:     uuid,
					CreateAt: carbon.Now().Format("Y-m-d H:i:s"),
					State:    "PENDING",
					Args:     string(argsString),
				})

				return ctx.JSON(types.Response{
					200,
					"Success task " + uuid + " delivered",
					nil,
				})
			}
		}

		if !isAllowBranch {
			return ctx.JSON(types.Response{
				200,
				"Unnecessary deployment",
				nil,
			})
		}
	}

	return ctx.JSON(types.Response{
		400,
		"Invalid signature",
		nil,
	})
}

// 判断是否分支是否需要发布
func isAllowBranch(crypt types.CryptDataConfig, ref string) (bool, error, types.TaskParams) {

	// 链接ssh
	sshHost := crypt.Project.Auth.Host
	sshUser := crypt.Project.Auth.User
	sshPassword := crypt.Project.Auth.Pwd
	sshType := "password"       //password 或者 key
	sshKeyPath := "id_rsa path" //ssh id_rsa.id 路径"
	sshPort := crypt.Project.Auth.Port

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
		logr.Clog.Errorf("创建ssh client 失败", err)
		return false, err, types.TaskParams{}
	}
	defer sshClient.Close()

	snowflake := carbon.Now().Format("Ymd") + logr.SnowFlakeId()
	tempDir := `~/.runner/` + crypt.Symbol + `/` + snowflake
	if err = interactive.InterSend(sshClient, fmt.Sprintf(`rm -rf %s && mkdir -p %s && cd %s && git init && git checkout -b %s && git remote add origin %s && git config core.sparsecheckout true && echo .runner-ci.yml >> .git/info/sparse-checkout && git pull origin %s`, tempDir, tempDir, tempDir, ref, crypt.Message.Repository.SshUrl, ref)); err != nil {
		logr.Clog.Errorf("检查部署配置异常, %v", err)
		return false, err, types.TaskParams{}
	}

	dir, _ := os.Getwd()
	cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("scp %s@%s:%s/.runner-ci.yml %s/runtime/git/%s.runner-ci.yml", crypt.Project.Auth.User, crypt.Project.Auth.Host, tempDir, dir, snowflake))
	cmd.Start()
	cmd.Wait()

	var (
		runnerCi types.RunnerCi
	)

	conf.ParseYaml(fmt.Sprintf("runtime/git/%s.runner-ci.yml", snowflake), &runnerCi)
	var (
		isContain bool             = false
		taskParam types.TaskParams = types.TaskParams{
			runnerCi.Environment,
			runnerCi.Prepare,
			runnerCi.Script,
			runnerCi.Release,
		}
	)

	for _, b := range runnerCi.Only {
		if strings.Compare(b, ref) == 0 {
			isContain = true
			break
		}
	}

	return isContain, nil, taskParam
}
