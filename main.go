package main

import (
	"github.com/gin-gonic/gin"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

var (
	DB *gorm.DB
)

type Todo struct {
	ID     int    "json:'id'"
	Title  string "json:'title'"
	Status bool   "json:'status'"
}

/*
	func initMySQL() (err error) {
		dsn := "go_admin:123456@tcp(127.0.0.1:3306)/go_demo?charset=utf8mb4&parseTime=True&loc=Local"
		DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		//DB, err := gorm.Open(mysql.Open(dsn))
		//DB, err = gorm.Open("mysql", dsn)
		if err != nil {
			return
		}
		// 测试DB连通性
		return DB.DB().ping()

}
*/
func main() {
	// 创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库
	dsn := "go_admin:123456@tcp(127.0.0.1:3306)/go_demo?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//err := initMySQL()
	if err != nil {
		panic(err)
	}
	//defer DB.Close()

	// 绑定模型
	DB.AutoMigrate(&Todo{})

	r := gin.Default()
	// 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// 定义v1 api
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", func(c *gin.Context) {

		})
		// 查看所有待办事项
		v1Group.GET("/todo", func(c *gin.Context) {

		})
		// 查看某一个代办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		// 修改某一个待办事项
		v1Group.PUT("/todo/:id", func(c *gin.Context) {

		})
		// 删除某一个待办事项
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {

		})
	}

	r.Run()

}
