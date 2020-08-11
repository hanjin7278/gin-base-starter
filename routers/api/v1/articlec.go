package v1

import "github.com/gin-gonic/gin"

// @Summary 获取文章列表
// @Description 描述信息
// @Success 200 {object} gin.H
// @Router /articlec/list [get]
func ArticlecList(r *gin.Context) {
	r.JSON(200, gin.H{
		"data": "我是文章列表内容",
	})
}
