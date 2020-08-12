package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hanjin7278/gin-base-starter/config"
	"github.com/hanjin7278/gin-base-starter/logs"
	"github.com/hanjin7278/gin-base-starter/utils/result"
)

type blogArticle struct {
	Id         int    `json:"id" gorose:"id"`
	TagId      int    `json:"tag_id" gorose:"tag_id"`
	Title      string `json:"title" gorose:"title"`
	Des        string `json:"des" gorose:"des"`
	Content    string `json:"content" gorose:"content"`
	CreatedBy  string `json:"createdBy" gorose:"created_by"`
	ModifiedBy string `json:"modifiedBy" gorose:"modified_by"`
	DeletedOn  int    `json:"deletedOn" gorose:"deleted_on"`
	State      int    `json:"state" gorose:"state"`
}

func (u *blogArticle) TableName() string {
	return "blog_article"
}

// @Summary 获取文章列表
// @Description 描述信息
// @Success 200 {object} result.R {"code":200,"data":null,"msg":"",time:"2020年08月11日17:58:03"}
// @Router /articlec/list [get]
func ArticlecList(r *gin.Context) {
	var as []blogArticle
	err := config.DB().Table(&as).Limit(10).Select()
	if err != nil {
		logs.Error("读取数据库错误", err)
	}
	r.JSON(200, result.Ok(as))
}
