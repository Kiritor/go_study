package main

import (
	"log"
	"fmt"
	"os"
	"runtime"
	"path"
	"strings"
)


func main(){


	logFile, err := os.OpenFile(CurrentName(), os.O_RDWR | os.O_CREATE |os.O_APPEND, 0777)
    if err != nil {
        fmt.Printf("open file error=%s\r\n", err.Error())
        os.Exit(-1)
    }

    defer logFile.Close()
    logger:=log.New(logFile,"\r\n", log.Ldate | log.Ltime | log.Llongfile)
    logger.Println("normal log 1")
    logger.Println("normal log 2")
}

/**
   得到当前执行程序的名称
*/

func CurrentName() string{
	//获取调用者本身的名称(包含路径)
	_,fullFileName,_,_ := runtime.Caller(0)
	fileNameWithSuffix:=path.Base(fullFileName)    //得到文件名+后缀
	suffix:=path.Ext(fileNameWithSuffix)           //得到扩展名也就是后缀
	fileName :=strings.TrimSuffix(fileNameWithSuffix,suffix)  //得到文件名
	return fileName+".log"
}
