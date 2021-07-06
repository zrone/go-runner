package handle

import (
	"awesome-runner/src/logr"
	"awesome-runner/src/sql"
	"awesome-runner/types"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/kataras/iris/v12"
	"sync"
)

type InternalDeploy struct {
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
	Secret string `json:"secret"`
	Path   string `json:"path"`
	User   string `json:"user"`
	Host   string `json:"host"`
	Port   int    `json:"port"`
	Pwd    string `json:"pwd,omitempty"`
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
		data = append(data, InternalDeploy{
			Symbol: item.Symbol,
			Name:   item.Name,
			Secret: item.Secret,
			Path:   item.Path,
			User:   item.Auth.User,
			Host:   item.Auth.Host,
			Port:   item.Auth.Port,
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

	internalDeploy := &types.InternalDeploy{
		Symbol: logr.SnowFlakeId(),
		Secret: body.Secret,
		Path:   body.Path,
		Auth: types.Authentication{
			1,
			body.User,
			body.Host,
			body.Port,
			body.Pwd,
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

	symbol := ctx.Params().Get("symbol")
	var ext types.InternalDeploy
	sql.GetLiteInstance().Where("symbol = ?", symbol).Take(&ext)
	if ext == (types.InternalDeploy{}) {
		ctx.JSON(types.Response{
			Code:    400,
			Message: "fail",
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
		Auth: types.Authentication{
			1,
			body.User,
			body.Host,
			body.Port,
			func(pwd string) string {
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
