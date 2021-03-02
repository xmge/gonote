package main

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type App struct {
	Name string
	SSHNames []string
	CurrentSSHName string
	LogPath string
	ErrorLineThreshold int64
	SamplingDuration time.Duration
}

func main() {
	river := App{}
	river.Name="runfast"
	river.SSHNames=[]string{"runfast01"}
	river.LogPath="/mnt/log/runfast/app.log"
	river.ErrorLineThreshold=50
	river.SamplingDuration=10
	err := river.restart()
	if err != nil {
		log.Panic(err)
	}
}


func (app *App)restart() error  {
	for _, sshName := range app.SSHNames {
		app.CurrentSSHName=sshName
		if err := app.restartOne();err != nil{
			errMsg := fmt.Sprintf("sshName:%s,error:%v",sshName,err)
			return errors.New(errMsg)
		}
	}
	return nil
}

func (app *App)restartOne() error {

	// 重启并记录当前日志打印行数
	startLogCount,err := app.restartServerAndGetLogLineCount()
	if err != nil {
		return err
	}
	fmt.Printf("%s,startLogCount:%d\n",app.CurrentSSHName,startLogCount)
	time.Sleep(time.Second * app.SamplingDuration)

	endLogCount,err := app.getLogLineCount()
	if err != nil {
		return err
	}
	fmt.Printf("%s,endLogCount:%d\n",app.CurrentSSHName,endLogCount)

	tailFileNumber := endLogCount-startLogCount
	if endLogCount < startLogCount {
		tailFileNumber = endLogCount
	}

	fmt.Printf("%s,tailFileNumber:%d\n",app.CurrentSSHName,tailFileNumber)
	errorLineCount,err := app.countErrorLogLineNum(tailFileNumber)
	if err != nil {
		return err
	}

	fmt.Printf("%s,errorLineCount:%d\n",app.CurrentSSHName,errorLineCount)
	if errorLineCount >= app.ErrorLineThreshold/int64(len(app.SSHNames)) {
		errMsg := fmt.Sprintf("error line num is over ErrorLineThreshold,errorLineCount:%d",errorLineCount)
		return errors.New(errMsg)
	}
	return nil
}

func (app *App)countErrorLogLineNum(tailFileNumber int64) (int64,error) {
	cmdStr := fmt.Sprintf("tail -n %d %s | grep ERROR | wc -l",tailFileNumber,app.LogPath)
	outputBytes,err := exec.Command("ssh",app.CurrentSSHName,cmdStr).Output()
	if err != nil {
		return -1,err
	}
	output := strings.ReplaceAll(string(outputBytes),"\n","")
	return strconv.ParseInt(output,10,64)
}

func (app *App)getLogLineCount() (int64,error) {
	cmdStr := fmt.Sprintf("wc -l %s | awk '{print $1}'",app.LogPath)
	outputBytes,err := exec.Command("ssh", app.CurrentSSHName, cmdStr).Output()
	if err != nil {
		panic(err)
	}
	output := strings.ReplaceAll(string(outputBytes),"\n","")
	return strconv.ParseInt(output,10,64)
}

func (app *App)restartServerAndGetLogLineCount() (int64,error) {
	cmdStr := fmt.Sprintf("sudo supervisorctl restart river >/dev/null 2>&1 && sudo supervisorctl status %s|awk '{print $2}' && wc -l %s | awk '{print $1}'",app.Name,app.LogPath)
	outputBytes,err := exec.Command("ssh", app.CurrentSSHName,cmdStr).Output()

	if err != nil {
		return -1,err
	}
	outputArray := strings.Split(string(outputBytes),"\n")
	if len(outputBytes) < 2 {
		errMsg := fmt.Sprintf("ssh return result. error:%s",string(outputBytes))
		return -1,errors.New(errMsg)
	}

	if outputArray[0] != "RUNNING" {
		errMsg := fmt.Sprintf("restart error, status:%s",string(outputBytes))
		return -1,errors.New(errMsg)
	}
	countStr := strings.ReplaceAll(outputArray[1],"\n","")
	count,err := strconv.ParseInt(countStr,10,64)
	if err != nil {
		return -1,err
	}
	return count,nil
}
