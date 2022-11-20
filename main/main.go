package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// main函数
func main() {
	r := gin.Default()      // 创建路由
	r.GET("/setu", setuApi) // 设置路由
	r.Run(":23856")         // 端口
}

// setuApi函数
func setuApi(c *gin.Context) {
	keyword := c.DefaultQuery("tag", "")     // 获取tag参数,默认为空
	argum := c.DefaultQuery("num", "1")      // 获取num参数,默认为1
	argr18 := c.DefaultQuery("r18", "false") // 获取r18参数,默认为false
	argumInt, err := strconv.Atoi(argum)     // 将num参数转换为int
	if err != nil {                          // 错误处理
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "num参数错误",
		})
		return
	}
	// num最大1000
	if argumInt > 1000 {
		argumInt = 1000
	}

	argr18Bool, err := strconv.ParseBool(argr18) // 将r18参数转换为bool
	if err != nil {                              // 错误处理
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "r18参数错误",
		})
		return
	}
	c.JSONP(http.StatusOK, selectSql(keyword, argumInt, argr18Bool)) // 调用selectSql函数, 给路由返回json
}

// 错误处理函数
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// 返回api的内容, SetuList结构体
func selectSql(keyword string, argum int, argr18 bool) SetuList {
	db, err := sql.Open("sqlite3", "./database/setu.db") // 打开数据库
	checkErr(err)                                        // 错误处理
	defer db.Close()                                     // 关闭数据库
	r18String := "False"                                 // r18参数
	numString := fmt.Sprintf("%d", argum)                // 将num参数转换为string
	if argr18 {
		r18String = "True"
	} else {
		r18String = "False"
	}
	keyword = strings.ReplaceAll(keyword, "'", "") // 去除单引号防止sql注入
	// 查询语句
	sqlCommand := "select * from main where (tags like '%" + keyword + "%' or title like '%" + keyword + "%' or author like '%" + keyword + "%')and r18 = '" + r18String + "' order by random() limit " + numString
	rows, err := db.Query(sqlCommand) // 执行查询语句
	checkErr(err)                     // 错误处理
	data := make([]Setu, 0)           // 创建Setu切片
	for rows.Next() {                 // 遍历查询结果
		var pid, p, uid int                                                        // 定义变量
		var title, author, r18, tags, ext, urls string                             // 定义变量
		err = rows.Scan(&pid, &p, &uid, &title, &author, &r18, &tags, &ext, &urls) // 将查询结果赋值给变量
		checkErr(err)                                                              // 错误处理
		tagsArray := strings.Split(tags, ",")                                      // 将tags字段分割为切片
		newTagsArray := make([]string, 0)                                          // 创建新的切片
		for _, tag := range tagsArray {                                            // 遍历tags切片
			tag = strings.Trim(tag, " ")             // 去除空格
			tag = strings.Trim(tag, "\"")            // 去除空格和双引号
			newTagsArray = append(newTagsArray, tag) // 将tags切片添加到新的切片
		}
		r18Bool := false // 定义r18Bool变量
		if r18 == "True" {
			r18Bool = true // 如果r18字段为True,则r18Bool为true
		} else {
			r18Bool = false // 如果r18字段为False,则r18Bool为false
		}
		setu := NewSetu(pid, p, uid, title, author, r18Bool, newTagsArray, ext, urls) // 调用NewSetu函数, 创建Setu结构体
		data = append(data, setu)                                                     // 将Setu结构体添加到Setu切片
	}
	return NewSetuList(200, "success", data) // 调用NewSetuList函数, 创建SetuList结构体, 并且返回
}
