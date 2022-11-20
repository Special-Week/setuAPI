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

func main() {
	r := gin.Default()
	r.GET("/setu", setuApi)
	r.Run(":23856")
}

func setuApi(c *gin.Context) {
	keyword := c.DefaultQuery("tag", "")
	argum := c.DefaultQuery("num", "1")
	argr18 := c.DefaultQuery("r18", "false")
	argumInt, err := strconv.Atoi(argum)
	if err != nil {
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

	argr18Bool, err := strconv.ParseBool(argr18)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "r18参数错误",
		})
		return
	}
	c.JSONP(http.StatusOK, selectSql(keyword, argumInt, argr18Bool))
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func selectSql(keyword string, argum int, argr18 bool) SetuList {
	db, err := sql.Open("sqlite3", "./database/setu.db")
	checkErr(err)
	defer db.Close()
	r18String := "False"
	numString := fmt.Sprintf("%d", argum)
	if argr18 {
		r18String = "True"
	} else {
		r18String = "False"
	}
	sqlCommand := "select * from main where (tags like '%" + keyword + "%' or title like '%" + keyword + "%' or author like '%" + keyword + "%')and r18 = '" + r18String + "' order by random() limit " + numString
	rows, err := db.Query(sqlCommand)
	checkErr(err)
	data := make([]Setu, 0)
	for rows.Next() {
		var pid, p, uid int
		var title, author, r18, tags, ext, urls string
		err = rows.Scan(&pid, &p, &uid, &title, &author, &r18, &tags, &ext, &urls)
		checkErr(err)
		tagsArray := strings.Split(tags, ",")
		newTagsArray := make([]string, 0)
		for _, tag := range tagsArray {
			tag = strings.Trim(tag, " ")
			tag = strings.Trim(tag, "\"")
			newTagsArray = append(newTagsArray, tag)
		}
		r18Bool := false
		if r18 == "True" {
			r18Bool = true
		} else {
			r18Bool = false
		}
		setu := NewSetu(pid, p, uid, title, author, r18Bool, newTagsArray, ext, urls)
		data = append(data, setu)
	}
	return NewSetuList(200, "success", data)
}
