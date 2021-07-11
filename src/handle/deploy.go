package handle

import (
	"awesome-runner/src/logic"
	"awesome-runner/src/logr"
	"awesome-runner/src/sql"
	"awesome-runner/types"
	"strconv"
	"strings"

	"github.com/kataras/iris/v12"
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

	// 验证ssh链接配置
	err := BuildProjConfigure(symbol)
	if err != nil {
		ctx.JSON(types.Response{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	if symbol == "" {
		ctx.JSON(types.Response{
			Code:    400,
			Message: "Unknown symbol",
			Data:    nil,
		})
		return
	}

	// 读取项目配置
	var internalDeloy types.InternalDeploy
	sql.GetLiteInstance().Take(&internalDeloy, "symbol = ?", symbol)
	if internalDeloy == (types.InternalDeploy{}) {
		ctx.JSON(types.Response{
			Code:    400,
			Message: "Unknown symbol",
			Data:    nil,
		})
		return
	}

	if internalDeloy.Option != 1 {
		ctx.JSON(types.Response{
			Code: 400,
			// "Invalid Project Option",
			Message: "当前项目为上线发布类型，不支持自动化部署",
			Data:    nil,
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
