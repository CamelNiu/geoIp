package LogFmt

import (
	"log"
	"os"
	"strings"
	"time"
)

var (
	ERR_MSG = "-err.log"
	VISIT_MSG = "-visit.log"
)

func SetErrLog(msg string) {
	path := getParentDir()
	toDay := time.Now().Format("20060102")
	logPath := path+"/logfile/"
	createLogPath(logPath)
	file := logPath+toDay+ERR_MSG
	logFile, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[err]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	log.Println(msg)
}

func SetVisitLog(msg string) {
	path := getParentDir()
	toDay := time.Now().Format("20060102")
	logPath := path+"/logfile/"
	createLogPath(logPath)
	file := logPath+toDay+VISIT_MSG
	logFile, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[visit]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	log.Println(msg)
}

func getParentDir() string {
	path, _ := os.Getwd()
	itemPath := getParentDirectory(path)
	return itemPath
}


func createLogPath(filePath string){
	file,err := os.Open(filePath)
	defer func(){file.Close()}()
	if err != nil && os.IsNotExist(err) {
		os.MkdirAll(filePath,os.ModePerm)
	}
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func getParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}
