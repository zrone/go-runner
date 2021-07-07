package handle

import (
	"awesome-runner/src/logr"
	"awesome-runner/src/sql"
	"awesome-runner/types"
	"github.com/kataras/iris/v12"
)

type LoginForm struct {
	UserName string `json:"userName"`
	Password string `json:"password,omitempty"`
}

func LoginAccount(ctx iris.Context) {
	var (
		manager types.Manager
		body    LoginForm
	)
	bodyByte, _ := ctx.GetBody()
	logr.JSON.Unmarshal(bodyByte, &body)

	err := sql.GetLiteInstance().Take(&manager, "user_name = ?", body.UserName).Error
	if err != nil || manager.Password != body.Password {
		ctx.JSON(map[string]interface{}{
			"currentAuthority": "guest",
			"status":           "error",
			"type":             "account",
		})
		return
	}
	ctx.JSON(map[string]interface{}{
		"currentAuthority": "admin",
		"status":           "ok",
		"type":             "account",
	})
}

func CurrentUser(ctx iris.Context) {
	ctx.JSON(map[string]interface{}{
		"name":   "admin",
		"userid": "00000001",
		"avatar": "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
	})
}
