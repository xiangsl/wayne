package images

import (
	"strconv"

	"github.com/Qihoo360/wayne/src/backend/controllers/base"
	"github.com/Qihoo360/wayne/src/backend/models"
	"github.com/Qihoo360/wayne/src/backend/util/logs"
)

// 操作App相关资源
type ImagesController struct {
	base.APIController
}

func (c *ImagesController) URLMapping() {
	c.Mapping("List", c.List)
}

// @Title List/
// @Description get all app
// @Param	starred		query 	bool	false		"is starred app.default not star"
// @Param	pageNo		query 	int	false		"the page current no"
// @Param	pageSize		query 	int	false		"the page size"
// @Param	name		query 	string	false		"name filter"
// @Param	namespaceId		query 	int	false		"namespace id"
// @Param	deleted		query 	bool	false		"is deleted, default list all."
// @Success 200 {object} []models.App success
// @router / [get]
func (c *ImagesController) List() {
	param := c.BuildQueryParam()
	pageNo := c.Input().Get("pageNo")
	pageSize := c.Input().Get("pageSize")
	versionName := c.Input().Get("versionName")
	intPageNo, err := strconv.Atoi(pageNo)
	intPageSize, err := strconv.Atoi(pageSize)
	startNum := strconv.Itoa((intPageNo - 1) * intPageSize)
	countNum, apps, err := models.ImagesModel.ListImages(startNum, pageSize, versionName)
	if err != nil {
		logs.Error("list by param (%s) error. %v", param, err)
		c.HandleError(err)
		return
	}

	c.Success(param.NewPage(countNum, apps))
	return
}
