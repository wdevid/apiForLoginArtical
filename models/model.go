package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

var (
	globalnum int
)

func init() {
	globalnum = 100000000
}

type User struct {
	UserName string
	PassWord string `orm:"size(500)"`
	Id       int64 `auto:"true" index:"pk"`
	Customer    *Customer   `orm:"rel(one)"` // OneToOne relation

}

type Store struct {
	Id              int64 `auto:"true" index:"pk"`
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

type Customer struct {
	Id              int64 `auto:"true" index:"pk"`
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	User        *User   `orm:"reverse(one)"` // 设置一对一反向关系(可选)
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
	Public          int64
}

func RegisterDB() {
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册 model
	orm.RegisterModel(new(User), new(Store), new(Customer))
	//注册默认数据库
	orm.RegisterDataBase("default", "mysql", "root:root@/myapp?charset=utf8", 30, 30) //密码为空格式
}

func Generate() (tiny string) {
	globalnum++
	num := globalnum
	fmt.Println(num)
	alpha := merge(getRange(48, 57), getRange(65, 90))
	alpha = merge(alpha, getRange(97, 122))
	if num < 62 {
		tiny = string(alpha[num])
		return tiny
	} else {
		var runes []rune
		runes = append(runes, alpha[num%62])
		num = num / 62
		for num >= 1 {
			if num < 62 {
				runes = append(runes, alpha[num-1])
			} else {
				runes = append(runes, alpha[num%62])
			}
			num = num / 62

		}
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		tiny = string(runes)
		return tiny
	}
	return tiny
}

func getRange(start, end rune) (ran []rune) {
	for i := start; i <= end; i++ {
		ran = append(ran, i)
	}
	return ran
}

func merge(a, b []rune) []rune {
	c := make([]rune, len(a)+len(b))
	copy(c, a)
	copy(c[len(a):], b)
	return c
}
