package handle

import (
	"awesome-runner/src/logr"
	"awesome-runner/src/sql"
	"awesome-runner/types"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"sync"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/golang-module/carbon"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

type LoginForm struct {
	UserName string `json:"userName"`
	Password string `json:"password,omitempty"`
}

type UpdateUserPasswordForm struct {
	Password string `json:"password" validate:"required,min=6"`
}

// 登录
func LoginAccount(ctx iris.Context) {
	var (
		manager types.Manager
		body    LoginForm
	)
	bodyByte, _ := ctx.GetBody()
	logr.JSON.Unmarshal(bodyByte, &body)

	err := sql.GetLiteInstance().Take(&manager, "is_delete = 0 AND user_name = ?", body.UserName).Error

	if err != nil || !isPasswordValid(body.Password, manager.Salt, manager.Password) {
		ctx.JSON(map[string]interface{}{
			"currentAuthority": "guest",
			"status":           "error",
			"type":             "account",
		})
		return
	}

	// session 登录信息
	session := sessions.Get(ctx)
	session.Set("Runner:Login", true)
	session.Set("Runner:User:ID", manager.ID)

	ctx.JSON(map[string]interface{}{
		"currentAuthority": manager.UserName,
		"status":           "ok",
		"type":             "account",
	})
}

// 当前用户信息
func CurrentUser(ctx iris.Context) {
	var manager types.Manager
	session := sessions.Get(ctx)

	id, err := session.GetInt64("Runner:User:ID")
	if err != nil {
		ctx.JSON(map[string]interface{}{})
		return
	}

	err = sql.GetLiteInstance().Take(&manager, "is_delete = 0 AND id = ?", id).Error
	if err != nil {
		ctx.JSON(map[string]interface{}{})
		return
	}

	ctx.JSON(map[string]interface{}{
		"name":   manager.UserName,
		"userid": manager.ID,
		"avatar": "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
	})
}

// 退出登录
func Logout(ctx iris.Context) {
	session := sessions.Get(ctx)
	session.Delete("Runner:Login")
	session.Delete("Runner:User:ID")
}

// 用户列表
func UserList(ctx iris.Context) {
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
		list    []types.Manager
		data    []types.Manager
		total   int64
		manager types.Manager
	)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		sql.GetLiteInstance().Where("is_delete = ?", 0).Offset((page - 1) * size).Limit(size).Find(&list)
	}()
	go func() {
		defer wg.Done()
		sql.GetLiteInstance().Model(&manager).Where("is_delete = ?", 0).Count(&total)
	}()
	wg.Wait()

	for _, v := range list {
		data = append(data, types.Manager{
			ID:       v.ID,
			UserName: v.UserName,
			CreateAt: v.CreateAt,
			UpdateAt: v.CreateAt,
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

// 添加用户
func UserCreate(ctx iris.Context) {
	var body types.Manager
	bodyByte, _ := ctx.GetBody()
	logr.JSON.Unmarshal(bodyByte, &body)

	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()

	zh_translations.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(body)
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

	body.Salt = logr.SnowFlakeId()
	body.Password = encodePassword(body.Password, body.Salt)
	body.IsDelete = false
	body.CreateAt = carbon.Now().ToTimestamp()
	body.UpdateAt = carbon.Now().ToTimestamp()

	if err := sql.GetLiteInstance().Create(&body).Error; err != nil {
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

// 修改密码
func UserUpdate(ctx iris.Context) {
	var body UpdateUserPasswordForm
	bodyByte, _ := ctx.GetBody()
	logr.JSON.Unmarshal(bodyByte, &body)

	id, _ := ctx.Params().GetInt64("id")
	var ext types.Manager
	sql.GetLiteInstance().Where("id = ?", id).Take(&ext)
	if ext == (types.Manager{}) {
		ctx.JSON(types.Response{
			Code:    400,
			Message: "Unknown params",
		})
	}

	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()

	zh_translations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(body)
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

	salt := logr.SnowFlakeId()
	if err := sql.GetLiteInstance().Model(&ext).Updates(map[string]interface{}{
		"salt":      salt,
		"password":  encodePassword(body.Password, salt),
		"update_at": carbon.Now().ToTimestamp(),
	}).Error; err != nil {
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

// 删除用户
func UserDelete(ctx iris.Context) {
	var manager = types.Manager{}
	id, err := ctx.Params().GetInt64("id")
	if err != nil {
		ctx.JSON(types.Response{
			Code:    200,
			Message: "Invalid param",
			Data:    nil,
		})
		return
	}

	if err := sql.GetLiteInstance().Model(&manager).Where("is_delete = 0 AND id = ?", id).Update("is_delete", true).Error; err != nil {
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

// 加密
func encodePassword(password string, salt string) string {
	m := hmac.New(sha256.New, []byte(salt))
	m.Write([]byte(password))

	return hex.EncodeToString(m.Sum(nil))
}

// 校验
func isPasswordValid(password string, salt string, raw string) bool {
	return encodePassword(password, salt) == raw
}
