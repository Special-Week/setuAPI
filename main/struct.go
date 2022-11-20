package main

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

type SetuList struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []Setu `json:"data"`
}

func NewSetu(pid int, p int, uid int, title string, author string, r18 bool, tags []string, ext string, urls string) Setu {
	return Setu{
		Pid:    pid,
		P:      p,
		Uid:    uid,
		Title:  title,
		Author: author,
		R18:    r18,
		Tags:   tags,
		Ext:    ext,
		Urls:   urls,
	}
}

func NewSetuList(code int, message string, data []Setu) SetuList {
	return SetuList{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
