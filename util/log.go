package log

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	LOGPATH  = "log/"
	FORMAT   = "20200102"
	LINEFEED = "\r\n"
	INFO     = "[INFO]: "
)

var path = LOGPATH + time.Now().Format(FORMAT) + "/"

//WriteLog return error
func WriteLog(fileName, msg string) error {
	if !ifexist(path) {
		return createdir(path)
	}
	f, err := os.OpenFile(path+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	_, err = io.WriteString(f, LINEFEED+msg)

	defer f.Close()
	return err
}

func WriteLogf(fileName, format string, arg ...interface{}) error {
	if !ifexist(path) {
		return createdir(path)
	}
	f, err := os.OpenFile(path+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	msg := fmt.Sprintf(format, arg...)

	_, err = io.WriteString(f, LINEFEED+INFO+msg)

	defer f.Close()
	return err
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
