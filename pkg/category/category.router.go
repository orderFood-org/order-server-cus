package category

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册分类相关路由
func RegisterRoutes(r *gin.RouterGroup) {
	categoryGroup := r.Group("/categories")
	{
		categoryGroup.GET("", listCategories)
		categoryGroup.GET("/active", getActiveCategories)
		categoryGroup.GET("/:id", getCategory)
		categoryGroup.POST("", createCategory)
		categoryGroup.PUT("/:id", updateCategory)
		categoryGroup.DELETE("/:id", deleteCategory)
	}
}

// listCategories 获取分类列表
func listCategories(c *gin.Context) {
	// 获取所有分类
	c.JSON(200, gin.H{"message": "获取所有分类"})
}

// getActiveCategories 获取活跃分类
func getActiveCategories(c *gin.Context) {
	// 获取活跃分类
	c.JSON(200, gin.H{"message": "获取活跃分类"})
}

// getCategory 获取单个分类
func getCategory(c *gin.Context) {
	// 获取分类详情
	c.JSON(200, gin.H{"message": "获取分类详情"})
}

// createCategory 创建分类
func createCategory(c *gin.Context) {
	// 创建新分类
	c.JSON(201, gin.H{"message": "创建分类成功"})
}

// updateCategory 更新分类
func updateCategory(c *gin.Context) {
	// 更新分类信息
	c.JSON(200, gin.H{"message": "更新分类成功"})
}

// deleteCategory 删除分类
func deleteCategory(c *gin.Context) {
	// 删除分类
	c.JSON(200, gin.H{"message": "删除分类成功"})
}
