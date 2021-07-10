package handle

import (
	"awesome-runner/src/config"
	"awesome-runner/src/logr"
	"awesome-runner/src/queue"
	"awesome-runner/src/sql"
	interactive "awesome-runner/src/ssh"
	"awesome-runner/types"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/RichardKnop/machinery/v2/tasks"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/golang-module/carbon"
	"github.com/kataras/iris/v12"
	taskLogrus "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

type InternalDeploy struct {
	Symbol        string `json:"symbol"`
	ReleaseSymbol string `json:"release_symbol"`
	Name          string `json:"name"`
	Secret        string `json:"secret"`
	Path          string `json:"path"`
	Option        uint8  `json:"option"`
	User          string `json:"user"`
	Host          string `json:"host"`
	Port          int    `json:"port"`
	Pwd           string `json:"pwd,omitempty"`
}

type ReleaseBody struct {
	Branch  string `json:"branch"`
	Symbol  string `json:"symbol"`
	HtmlUrl string `json:"html_url"`
}

// 项目列表
func ProjList(ctx iris.Context) {
	var (
		page int = 1
		size int = 20
	)

	page, err := ctx.URLParamInt("current")
	if err != nil {
		page = 1
	}
	size, err = ctx.URLParamInt("pageSize")
	if err != nil {
		size = 20
	}

	var (
		list           []types.InternalDeploy
		data           []InternalDeploy
		total          int64
		internalDeploy types.InternalDeploy
	)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		sql.GetLiteInstance().Where("is_delete = ?", 0).Offset((page - 1) * size).Limit(size).Find(&list)
	}()
	go func() {
		defer wg.Done()
		sql.GetLiteInstance().Model(&internalDeploy).Where("is_delete = ?", 0).Count(&total)
	}()
	wg.Wait()

	for _, item := range list {
		var temp string
		if item.Option == 2 {
			temp = ""
		} else {
			temp = config.Cnf.Domain + "?symbol=" + item.Symbol
		}

		data = append(data, InternalDeploy{
			Symbol:        temp,
			ReleaseSymbol: item.Symbol,
			Name:          item.Name,
			Secret:        item.Secret,
			Path:          item.Path,
			Option:        item.Option,
			User:          item.Auth.User,
			Host:          item.Auth.Host,
			Port:          item.Auth.Port,
		})
	}

	ctx.JSON(map[string]interface{}{
		"current":  page,
		"data":     data,
		"pageSize": size,
		"success":  true,
		"total":    total,
	})
}

// 创建项目
func ProjCreate(ctx iris.Context) {
	var body InternalDeploy
	bodyByte, _ := ctx.GetBody()
	logr.JSON.Unmarshal(bodyByte, &body)

	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()

	zh_translations.RegisterDefaultTranslations(validate, trans)

	if body.Option == 1 && body.Secret == "" {
		ctx.JSON(types.Response{
			Code:    400,
			Message: "自动化部署秘钥不能为空",
		})
		return
	}

	internalDeploy := &types.InternalDeploy{
		Symbol: logr.SnowFlakeId(),
		Secret: body.Secret,
		Path:   body.Path,
		Option: body.Option,
		Auth: types.Authentication{
			Scheme: 1,
			User:   body.User,
			Host:   body.Host,
			Port:   body.Port,
			Pwd:    body.Pwd,
		},
		IsDelete: false,
		Name:     body.Name,
	}

	err := validate.Struct(internalDeploy)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		// from here you can create your own error messages in whatever language you wish
		ctx.JSON(types.Response{
			Code: 400,
			Message: func(errs validator.ValidationErrors) string {
				var re string
				for _, e := range errs.Translate(trans) {
					re = e
					break
				}
				return re
			}(errs),
		})
		return
	}

	if err := sql.GetLiteInstance().Create(&internalDeploy).Error; err != nil {
		ctx.JSON(types.Response{
			Code:    400,
			Message: "fail",
		})
	} else {
		ctx.JSON(types.Response{
			Code:    200,
			Message: "success",
		})
	}
}

// 更新项目
func ProjUpdate(ctx iris.Context) {
	var body InternalDeploy
	bodyByte, _ := ctx.GetBody()
	logr.JSON.Unmarshal(bodyByte, &body)

	if body.Option == 1 && body.Secret == "" {
		ctx.JSON(types.Response{
			Code:    400,
			Message: "自动化部署秘钥不能为空",
		})
		return
	}

	symbol := ctx.Params().Get("symbol")
	var ext types.InternalDeploy
	sql.GetLiteInstance().Where("symbol = ?", symbol).Take(&ext)
	if ext == (types.InternalDeploy{}) {
		ctx.JSON(types.Response{
			Code:    400,
			Message: "Unknown symbol",
		})
	}

	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()

	zh_translations.RegisterDefaultTranslations(validate, trans)

	internalDeploy := &types.InternalDeploy{
		Symbol: symbol,
		Secret: body.Secret,
		Path:   body.Path,
		Option: body.Option,
		Auth: types.Authentication{
			Scheme: 1,
			User:   body.User,
			Host:   body.Host,
			Port:   body.Port,
			Pwd: func(pwd string) string {
				if pwd == "" {
					return ext.Auth.Pwd
				}
				return pwd
			}(body.Pwd),
		},
		IsDelete: ext.IsDelete,
		Name:     body.Name,
	}

	err := validate.Struct(internalDeploy)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		// from here you can create your own error messages in whatever language you wish
		ctx.JSON(types.Response{
			Code: 400,
			Message: func(errs validator.ValidationErrors) string {
				var re string
				for _, e := range errs.Translate(trans) {
					re = e
					break
				}
				return re
			}(errs),
		})
		return
	}

	if err := sql.GetLiteInstance().Model(&internalDeploy).Updates(&internalDeploy).Error; err != nil {
		ctx.JSON(types.Response{
			Code:    400,
			Message: "fail",
		})
	} else {
		ctx.JSON(types.Response{
			Code:    200,
			Message: "success",
		})
	}
}

// 删除项目
func ProjDelete(ctx iris.Context) {
	var internalDeploy = types.InternalDeploy{}
	symbol := ctx.Params().Get("symbol")

	if err := sql.GetLiteInstance().Model(&internalDeploy).Where("symbol = ?", symbol).Update("is_delete", true).Error; err != nil {
		ctx.JSON(types.Response{
			Code:    400,
			Message: "fail",
		})
	} else {
		ctx.JSON(types.Response{
			Code:    200,
			Message: "success",
		})
	}
}

// 发布上线单
func TaskPublish(ctx iris.Context) {
	var (
		crypt         types.AbstractCrypt
		body          ReleaseBody
		internalDeloy types.InternalDeploy
	)

	bodyByte, _ := ctx.GetBody()
	logr.JSON.Unmarshal(bodyByte, &body)

	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()

	zh_translations.RegisterDefaultTranslations(validate, trans)

	if body.Symbol == "" && body.Branch == "" {
		ctx.JSON(types.Response{
			Code:    400,
			Message: "参数错误",
		})
		return
	}

	sql.GetLiteInstance().Where("symbol = ?", body.Symbol).Take(&internalDeloy)
	if internalDeloy == (types.InternalDeploy{}) {
		ctx.JSON(types.Response{
			Code:    400,
			Message: "Unknown symbol",
		})
		return
	}

	resArr, err, params := queryRepo(internalDeloy, body)
	if err != nil || len(resArr) != 2 {
		ctx.JSON(types.Response{
			Code:    400,
			Message: "fail",
		})
		return
	}

	cryptDataConfig := types.CryptDataConfig{
		Symbol: body.Symbol,
		Message: types.Message{
			Ref:      fmt.Sprintf(`refs/heads/%s`, body.Branch),
			After:    strings.TrimSpace(string(resArr[0])),
			UserName: strings.TrimSpace(string(resArr[1])),
			Repository: types.Repository{
				HtmlUrl: body.HtmlUrl,
				SshUrl:  body.HtmlUrl,
			},
		},
		Headers: ctx.Request().Header,
		Project: types.Project{
			Secret: internalDeloy.Secret,
			Path:   internalDeloy.Path,
			Auth:   internalDeloy.Auth,
		},
		InternalDeloy: internalDeloy,
		Payload:       "",
	}

	htmlUrl := body.HtmlUrl
	if htmlUrl == "" {
		crypt = types.DiscoverGitlabCrypt(cryptDataConfig)
	} else {
		if strings.Contains(htmlUrl, "https://github.com") {
			crypt = types.DiscoverGithubCrypt(cryptDataConfig)
		} else {
			crypt = types.DiscoverGiteeCrypt(cryptDataConfig)
		}
	}

	delivered(ctx, crypt, body, params)
}

// 投递任务
func delivered(ctx iris.Context, crypt types.AbstractCrypt, body ReleaseBody, params types.TaskParams) {
	logr.Logrus.Printf("当前分支 %s", body.Branch)
	uuid := logr.SnowFlakeId()
	taskLog, _ := os.Create("runtime/task/" + uuid + ".log")
	defer taskLog.Close()

	taskLogrus.SetOutput(taskLog)

	env, _ := logr.JSON.Marshal(params.Environment)

	eta := time.Now().Add(time.Second * 3)
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
			Value: body.Branch,
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
	tl := types.TaskLog{
		Symbol:    crypt.GetCryptDataConfig().InternalDeloy.Symbol,
		Committer: crypt.GetCryptDataConfig().Message.UserName,
		Version:   string([]byte(crypt.GetCryptDataConfig().Message.After)[:7]),
		Uuid:      uuid,
		CreateAt:  carbon.Now().ToTimestamp(),
		Args:      string(argsString),
		Type:      2,
	}

	if _, err := queue.MachineryServer.SendTask(signature); err != nil {
		tl.State = "FAILURE"
		sql.GetLiteInstance().Create(&tl)

		taskLogrus.Errorf("Failed task delivered，%v", err)
		logr.Clog.Errorf("Failed task delivered，%v", err)

		ctx.JSON(types.Response{
			Code:    400,
			Message: "fail",
			Data:    nil,
		})
	} else {
		tl.State = "PENDING"
		sql.GetLiteInstance().Create(&tl)
		ctx.JSON(types.Response{
			Code:    200,
			Message: "Success task " + uuid + " delivered",
			Data:    nil,
		})
	}
}

func queryRepo(internalDeloy types.InternalDeploy, body ReleaseBody) ([]string, error, types.TaskParams) {
	// 链接ssh
	sshHost := internalDeloy.Auth.Host
	sshUser := internalDeloy.Auth.User
	sshPassword := internalDeloy.Auth.Pwd
	sshType := "password"       //password 或者 key
	sshKeyPath := "id_rsa path" //ssh id_rsa.id 路径"
	sshPort := internalDeloy.Auth.Port

	//创建sshp登陆配置
	conf := &ssh.ClientConfig{
		Timeout:         time.Second,
		User:            sshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	if sshType == "password" {
		conf.Auth = []ssh.AuthMethod{ssh.Password(sshPassword)}
	} else {
		conf.Auth = []ssh.AuthMethod{interactive.PublicKeyAuthFunc(sshKeyPath)}
	}

	//dial 获取ssh client
	addr := fmt.Sprintf("%s:%d", sshHost, sshPort)
	sshClient, err := ssh.Dial("tcp", addr, conf)
	if err != nil {
		logr.Clog.Errorf("创建ssh client 失败", err)
		return nil, err, types.TaskParams{}
	}
	defer sshClient.Close()

	snowflake := carbon.Now().Format("Ymd") + logr.SnowFlakeId()
	tempDir := `~/.runner/` + body.Symbol + `/` + snowflake
	if err = interactive.InterSend(sshClient, fmt.Sprintf(`rm -rf %s && mkdir -p %s && cd %s && git init && git checkout -b %s && git remote add origin %s && git config core.sparsecheckout true && echo .runner-ci.yml >> .git/info/sparse-checkout && git pull origin %s`, tempDir, tempDir, tempDir, body.Branch, body.HtmlUrl, body.Branch)); err != nil {
		logr.Clog.Errorf("检查部署配置异常, %v", err)
		return nil, err, types.TaskParams{}
	}

	//创建ssh-session
	session, err := sshClient.NewSession()
	if err != nil {
		logr.Clog.Errorf("创建ssh session 失败, %v", err)
		return nil, err, types.TaskParams{}
	}
	defer session.Close()

	res, err := session.CombinedOutput(`cd ` + tempDir + ` && git log --pretty=format:"%h+-+%an" -1 | awk '{print $0}'`)
	if err != nil {
		logr.Clog.Errorf("%v", err)
		return nil, err, types.TaskParams{}
	}
	resArr := strings.Split(string(res), "+-+")

	dir, _ := os.Getwd()
	cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("scp %s@%s:%s/.runner-ci.yml %s/runtime/git/%s.runner-ci.yml", internalDeloy.Auth.User, internalDeloy.Auth.Host, tempDir, dir, snowflake))
	cmd.Start()
	cmd.Wait()

	var (
		runnerCi types.RunnerCi
	)

	config.ParseYaml(fmt.Sprintf("runtime/git/%s.runner-ci.yml", snowflake), &runnerCi)
	var (
		taskParam types.TaskParams = types.TaskParams{
			Environment:  runnerCi.Environment,
			BeforeScript: runnerCi.Prepare,
			Script:       runnerCi.Script,
			AfterScript:  runnerCi.Release,
		}
	)

	return resArr, nil, taskParam
}
