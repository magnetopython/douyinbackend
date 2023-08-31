package main

import (
	"douyinbackend/config"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	var err error
	//初始化config
	err = config.ConfigInit()
	if err != nil {
		fmt.Println("配置导入失败")
		return
	}
	//初始化mysql
	err = config.InitDB()
	if err != nil {
		fmt.Println("数据库导入失败")
		return
	}
	//初始化gin
	r := gin.Default()

	initRouter(r)

	err = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
