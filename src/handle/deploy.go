package handle

import (
	"awesome-runner/src/logr"
	"awesome-runner/src/sql"
	"awesome-runner/types"
	"fmt"
	"github.com/kataras/iris/v12"
	"strings"
)

// 发送部署任务
func DeployHandle(ctx iris.Context) {
	symbol := ctx.URLParam("symbol")

	if symbol == "" {
		ctx.JSON(types.Response{
			400,
			"Unknown symbol",
			nil,
		})
		return
	}

	// 读取项目配置
	var internalDeloy types.InternalDeploy
	sql.GetLiteInstance().First(&internalDeloy, "symbol = ?", symbol)
	if internalDeloy == (types.InternalDeploy{}) {
		ctx.JSON(types.Response{
			400,
			"Unknown symbol",
			nil,
		})
		return
	}

	var (
		cryptDataConfig types.CryptDataConfig
		message         types.Message
	)

	body, _ := ctx.GetBody()
	logr.JSON.Unmarshal(body, &message)
	cryptDataConfig = types.CryptDataConfig{
		Symbol:  symbol,
		Message: message,
		Headers: ctx.Request().Header,
		Project: types.Project{
			Secret: internalDeloy.Secret,
			Path:   internalDeloy.Path,
		},
		Payload: string(body),
	}

	var (
		giteeCrypt  types.GiteeCrypt
		githubCrypt types.GithubCrypt
		gitlabCrypt types.GitlabCrypt
		result      bool
	)

	htmlUrl := message.Repository.HtmlUrl
	if htmlUrl == "" {
		gitlabCrypt = types.GitlabCrypt{
			cryptDataConfig,
		}
		types.HandleBuild(&gitlabCrypt, cryptDataConfig)
		types.HandleBuildPrefixCryptSign(&gitlabCrypt)
		result = types.HandleCompare(&gitlabCrypt)
	} else {
		if strings.Contains(htmlUrl, "https://github.com") {
			githubCrypt = types.GithubCrypt{
				cryptDataConfig,
			}
			types.HandleBuild(&githubCrypt, cryptDataConfig)
			types.HandleBuildPrefixCryptSign(&githubCrypt)
			result = types.HandleCompare(&githubCrypt)
		} else {
			giteeCrypt = types.GiteeCrypt{
				CryptDataConfig: cryptDataConfig,
			}
			types.HandleBuild(&giteeCrypt, cryptDataConfig)
			types.HandleBuildPrefixCryptSign(&giteeCrypt)
			result = types.HandleCompare(&giteeCrypt)
		}
	}

	fmt.Println(result)
	// 签名验证
	// 发送部署任务
	// 输出
}
