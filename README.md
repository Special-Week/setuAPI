# 使用Golang的[Gin](https://gin-gonic.com/zh-cn/)框架搭建一个简易的setuAPI


## 使用的第三方库
  
    gin         go get -u github.com/gin-gonic/gin
    sqlite3     go get -u github.com/mattn/go-sqlite3
    
## 设置路由以及端口部分, 当然想改可以任意
```golang
r.GET("/setu", setuApi) // 设置路由
r.Run(":23856")         // 端口
```
可以携带的访问参数: tag&num&r18
tag默认空, num默认1, r18默认false
```golang
keyword := c.DefaultQuery("tag", "")     // 获取tag参数,默认为空
argum := c.DefaultQuery("num", "1")      // 获取num参数,默认为1
argr18 := c.DefaultQuery("r18", "false") // 获取r18参数,默认为false
```
如果你想要白丝10张r18, 你应该这样访问:
> host:port/setu?num=10&tag=白丝&r18=true


## 响应内容中每份setu的返回内容
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


## api的返回内容
```golang
type SetuList struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []Setu `json:"data"`
}
```



## 实例
```json
{
    "code": 200,
    "message": "success",
    "data": [
        {
            "pid": 97413757,
            "p": 0,
            "uid": 4698160,
            "title": "20220405 Kyoka",
            "author": "レオナト",
            "r18": false,
            "tags": [
                "R-18",
                "ロリ",
                "萝莉",
                "loli",
                "プリンセスコネクト!Re:Dive",
                "公主连结Re:Dive",
                "プリコネR",
                "公主连结",
                "キョウカ(プリコネ)",
                "镜华（公主连结）",
                "ふんどし",
                "兜裆布",
                "イカ腹",
                "乌贼肚",
                "裸足",
                "赤脚"
            ],
            "ext": "jpg",
            "urls": "https://i.pixiv.re/img-original/img/2022/04/05/00/58/10/97413757_p0.jpg"
        },
        {
            "pid": 94654438,
            "p": 0,
            "uid": 801146,
            "title": "BOOTH通販のお知らせ",
            "author": "イコモチ",
            "r18": false,
            "tags": [
                "タペストリー",
                "挂毯",
                "沙花叉クロヱ",
                "沙花叉库洛艾",
                "非公式",
                "ホロライブ",
                "Hololive",
                "バーチャルYouTuber",
                "虚拟Youtuber",
                "holoX",
                "ピアス",
                "穿洞",
                "さかまた飼育日記",
                "Orca Raising Diary",
                "かきあげ",
                "撩头发"
            ],
            "ext": "png",
            "urls": "https://i.pixiv.re/img-original/img/2021/12/08/23/46/12/94654438_p0.png"
        }
    ]
}
```
