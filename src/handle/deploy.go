package handle

import (
	"awesome-runner/src/logic"
	"awesome-runner/src/logr"
	"awesome-runner/src/sql"
	"awesome-runner/types"
	"github.com/kataras/iris/v12"
	"strconv"
	"strings"
)

// 发送部署任务
func DeployHandle(ctx iris.Context) {

	var (
		cryptDataConfig types.CryptDataConfig
		message         types.Message
		crypt           types.AbstractCrypt
		symbol          string
	)
	symbolInt, _ := ctx.URLParamInt("symbol")
	symbol = strconv.Itoa(symbolInt)

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
	sql.GetLiteInstance().Take(&internalDeloy, "symbol = ?", symbol)
	if internalDeloy == (types.InternalDeploy{}) {
		ctx.JSON(types.Response{
			400,
			"Unknown symbol",
			nil,
		})
		return
	}

	body, _ := ctx.GetBody()
	logr.JSON.Unmarshal(body, &message)
	cryptDataConfig = types.CryptDataConfig{
		Symbol:  symbol,
		Message: message,
		Headers: ctx.Request().Header,
		Project: types.Project{
			Secret: internalDeloy.Secret,
			Path:   internalDeloy.Path,
			Auth:   internalDeloy.Auth,
		},
		InternalDeloy: internalDeloy,
		Payload:       string(body),
	}

	htmlUrl := message.Repository.HtmlUrl
	if htmlUrl == "" {
		crypt = types.DiscoverGitlabCrypt(cryptDataConfig)
	} else {
		if strings.Contains(htmlUrl, "https://github.com") {
			crypt = types.DiscoverGithubCrypt(cryptDataConfig)
		} else {
			crypt = types.DiscoverGiteeCrypt(cryptDataConfig)
		}
	}

	logic.SignatureVerification(ctx, crypt)
}
