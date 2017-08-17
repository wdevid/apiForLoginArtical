package mylog

import "github.com/astaxie/beego/logs"

func mylog(level int64,message string){
	log := logs.NewLogger()
	log.SetLogger(logs.AdapterFile,
		`{"filename":"project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	switch level {
	case 1:
		log.Debug(message)
	case 2:
		log.Error(message)
	}
}
/**
debug日志
 */
func Logers(message string)  {
	go mylog(1,message)
}
/**
错误日志
 */
func LogersError(message string)  {
	go mylog(2,message)
}
