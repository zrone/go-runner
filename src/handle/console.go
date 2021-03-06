package handle

import (
	"awesome-runner/src/sql"
	"awesome-runner/types"
	"github.com/kataras/iris/v12"
	"sync"
)

// 自动化部署任务列表
func ConsoleList(ctx iris.Context) {
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
		list    []types.TaskLog
		total   int64
		taskLog types.TaskLog
	)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		sql.GetLiteInstance().Where("type = ?", 1).Order("id DESC").Offset((page - 1) * size).Limit(size).Find(&list)
	}()
	go func() {
		defer wg.Done()
		sql.GetLiteInstance().Where("type = ?", 1).Model(&taskLog).Count(&total)
	}()
	wg.Wait()

	ctx.JSON(map[string]interface{}{
		"current":  page,
		"data":     list,
		"pageSize": size,
		"success":  true,
		"total":    total,
	})
}

// 发布任务列表
func ReleaseList(ctx iris.Context) {
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
		list    []types.TaskLog
		total   int64
		taskLog types.TaskLog
	)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		sql.GetLiteInstance().Where("type = ?", 2).Order("id DESC").Offset((page - 1) * size).Limit(size).Find(&list)
	}()
	go func() {
		defer wg.Done()
		sql.GetLiteInstance().Where("type = ?", 2).Model(&taskLog).Count(&total)
	}()
	wg.Wait()

	ctx.JSON(map[string]interface{}{
		"current":  page,
		"data":     list,
		"pageSize": size,
		"success":  true,
		"total":    total,
	})
}
