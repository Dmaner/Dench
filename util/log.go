package log

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	LOGPATH    = "log/"
	TIMEFORMAT = "20200102"
	LINEFEED   = "\r\n"
	INFOLOG    = "[INFO]: "
	ERRORLOG   = "[ERROR]: "
)

var path = LOGPATH + time.Now().Format(TIMEFORMAT) + "/"

//WriteLog return error
func WriteLog(fileName string, args ...interface{}) error {
	if !ifexist(path) {
		return createdir(path)
	}
	f, err := os.OpenFile(path+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	msg := fmt.Sprint(args...)
	_, err = io.WriteString(f, msg+LINEFEED)

	defer f.Close()
	return err
}

func WriteLogf(fileName, TIMEFORMAT string, arg ...interface{}) error {
	if !ifexist(path) {
		return createdir(path)
	}
	f, err := os.OpenFile(path+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	msg := fmt.Sprintf(TIMEFORMAT, arg...)

	_, err = io.WriteString(f, INFOLOG+msg+LINEFEED)

	defer f.Close()
	return err
}

// log.fatal
func ErrorLog(args ...interface{}) {
	fmt.Print(ERRORLOG)
	fmt.Println(args...)
}

//createdir
func createdir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	os.Chmod(path, os.ModePerm)
	return nil
}

//ifexist
func ifexist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}
