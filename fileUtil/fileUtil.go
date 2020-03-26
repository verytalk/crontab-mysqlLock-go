package fileUtil

import (
	"crontask/configUtil"
	"fmt"
	"os"
	"strings"
	"time"
)

func Tracefile(str_content string)  {
	fmt.Println(configUtil.Config.LogFile.FileName)
	fd,_:=os.OpenFile(configUtil.Config.LogFile.FileName,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	fd_time:=time.Now().Format("2006-01-02 15:04:05");
	fd_content:=strings.Join([]string{"======",fd_time,"=====\n",str_content,"\n=======end=======\n"},"")
	buf:=[]byte(fd_content)
	fd.Write(buf)
	fd.Close()
}
