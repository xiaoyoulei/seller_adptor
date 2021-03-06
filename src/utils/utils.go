package utils

//package main

import (
	"crypto/sha1"
	"fmt"
	"os"
	//	"github.com/op/go-logging"
	"io"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type LogLevel uint32

const (
	NoticeLevel  LogLevel = 1
	FatalLevel   LogLevel = 2
	WarningLevel LogLevel = 3
	DebugLevel   LogLevel = 4
)

type LogControl struct {
	TimeGap   int64  //间隔时间，单位秒
	FileName  string //日志文件名
	FilePath  string //日志路径
	FileOut   *os.File
	FileMutex sync.Mutex //日志锁
	LogLevel  LogLevel   //当前日志级别
	LogFormat string
}

var DebugLog *LogControl
var FatalLog *LogControl
var WarningLog *LogControl
var NoticeLog *LogControl
var GlobalLogLevel LogLevel

// 传入timegap 单位为分钟
func (this *LogControl) Init(timegap int64, filename string, filepath string, loglevel LogLevel) (err error) {
	// 内部转化为秒
	this.TimeGap = timegap * 60
	this.FileName = filename
	this.FilePath = filepath
	this.LogLevel = loglevel
	switch loglevel {
	case NoticeLevel:
		this.LogFormat = "NOTICE: "
	case FatalLevel:
		this.LogFormat = "FATAL: "
	case WarningLevel:
		this.LogFormat = "WARNING: "
	case DebugLevel:
		this.LogFormat = "DEBUG: "
	}
	if this.LogLevel > GlobalLogLevel {
		return
	}
	err = this.open_file()
	if err != nil {
		return
	}
	go this.LogCut()
	return
}

func (this *LogControl) Write(format string, args ...interface{}) (err error) {
	if this.LogLevel > GlobalLogLevel {
		return
	}
	this.FileMutex.Lock()
	defer this.FileMutex.Unlock()
	err = this.check_valid()
	if err != nil {
		return err
	}
	var body string
	head := fmt.Sprintf("%s %s * ", this.LogFormat, time.Now().Format("2006-01-02 15:04:05"))
	if args != nil {
		body = fmt.Sprintf(format, args...)
	} else {
		body = format
	}
	_, err = this.FileOut.Write([]byte(head + body + "\n"))
	return
}

func (this *LogControl) open_file() (err error) {
	this.FileOut, err = os.OpenFile(this.FilePath+"/"+this.FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	return
}

func (this *LogControl) check_valid() (err error) {
	//这部分代码有待商榷。。。时间开销
	_, err = os.Stat(this.FilePath + "/" + this.FileName)
	if err != nil {
		err = this.open_file()
		if err != nil {
			return
		}
	}
	return
}

func (this *LogControl) LogCut() {
	var err error
	for {
		nowtime := time.Now().Unix()
		nexttime := int64(nowtime/this.TimeGap+1) * this.TimeGap
		var delta time.Duration
		delta = time.Duration(nexttime - nowtime)
		time.Sleep(time.Second * delta)
		this.FileMutex.Lock()
		//		date_format := time.Now().Truncate(time.Duration(this.TimeGap * time.Second)).Format("200601021504")
		date_now := time.Now().Unix() - this.TimeGap
		date_format := time.Unix(date_now, 0).Format("200601021504")
		err = this.check_valid()
		if err != nil {
			log.Printf("check log file fail. err[%s]", err.Error())
			os.Exit(-1)
		}
		this.FileOut.Close()
		os.Rename(this.FilePath+this.FileName, this.FilePath+this.FileName+"."+date_format+"00")
		err = this.open_file()
		if err != nil {
			os.Exit(-1)
		}
		this.FileMutex.Unlock()
	}
	return
}

func GenSearchid(imei string) (searchid string) {

	var tmp string
	tmp = imei
	tmp += time.Now().String()
	tmp += strconv.Itoa(rand.Int())
	DebugLog.Write("searchid is %s", tmp)
	sha1_t := sha1.New()
	io.WriteString(sha1_t, tmp)
	searchid = fmt.Sprintf("%x", sha1_t.Sum(nil))
	return
}
