# 使用Golang的[Gin](https://gin-gonic.com/zh-cn/)框架搭建一个简易的setuAPI


使用的第三方库:
  
    gin         go get -u github.com/gin-gonic/gin
    sqlite3     go get -u github.com/mattn/go-sqlite3
    
设置路由以及端口部分, 当然想改可以任意
```golang
r.GET("/setu", setuApi) // 设置路由
r.Run(":23856")         // 端口
```

响应内容中每份setu的返回内容
```golang
type Setu struct {
	Pid    int      `json:"pid"`
	P      int      `json:"p"`
	Uid    int      `json:"uid"`
	Title  string   `json:"title"`
	Author string   `json:"author"`
	R18    bool     `json:"r18"`
	Tags   []string `json:"tags"`
	Ext    string   `json:"ext"`
	Urls   string   `json:"urls"`
}
```


api的返回内容
```golang
type SetuList struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []Setu `json:"data"`
}
```
