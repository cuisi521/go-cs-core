// package cs
// @Author cuisi
// @Date 2024/1/3 11:12:00
// @Desc
package cs

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/cuisi521/go-cs-core/sys/clog"
)

func TestLog(t *testing.T) {
	clog.SetSaveMod(false)
	// 设置日志输出的文件名以及行号
	clog.SetLineNumber(true)
	clog.SetLogPath("./log")
	clog.Install()
	Log().Infof("sssssss%s", "vvvv")
}

func TestGetType(t *testing.T) {
	var ts interface{}
	ts = []byte("sssss")
	vs := tStruct{Id: "888888"}
	v1 := GetType(ts)
	v2 := GetType(vs)
	fmt.Println(v1)
	fmt.Println(v2)
}

type tStruct struct {
	Id string
}

// 测试http get请求
// 请求方式：go test -v -run="TestHttpGet"
func TestHttpSendData(t *testing.T) {
	var wg sync.WaitGroup
	for j := 0; j < 10; j++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				c := Client()
				c.SetHeader("Authorization", "test")
				dataReq := map[string]string{"key": fmt.Sprintf("根据问题找答案%v", i),
					"value":  fmt.Sprintf("大模型是怎样学习的任务派遣规则%v", i),
					"name":   fmt.Sprintf("任务派遣规则，任务派遣规则根据工作流找答案%v", i),
					"remark": fmt.Sprintf("任务派遣规则任务派遣规则，根据工作流找答案%v", i),
					"code":   fmt.Sprintf("任务派遣规则，任务派遣规则根据工作流找答案%v", i),
					"uid":    fmt.Sprintf("任务派遣规则，任务派遣规则根据工作流找答案%v", i),
					"face":   fmt.Sprintf("任务派遣规则任务派遣规则，根据工作流找答案%v", i),
					"bz1":    fmt.Sprintf("任务派遣规则任务派遣规则，根据工作流找答案%v", i),
					"bz2":    fmt.Sprintf("任务派遣规则任务派遣规则，根据工作流找答案%v", i),
					"bz3":    fmt.Sprintf("任务派遣规则任务派遣规则，根据工作流找答案%v", i),
					"bz4":    fmt.Sprintf("任务派遣规则任务派遣规则，根据工作流找答案%v", i),
					"bz5":    fmt.Sprintf("任务派遣规则，任务派遣规则根据工作流找答案%v", i)}
				dataReqByte, err := json.Marshal(&dataReq)
				if err != nil {
					fmt.Println(err.Error())
				}
				data := &DataCenterReq{
					Producer: "330100",
					Consumer: "330000",
					Data:     dataReqByte,
					SendAt:   time.Now().Unix(),
				}
				dataByte, err := json.Marshal(data)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				resp, err := c.Post("http://localhost:8001/api/v1/dc/dataCenter/sendData", dataByte)
				if err != nil {
					fmt.Println(err.Error())
				} else {
					result := resp.ReadAllString()
					fmt.Println("result:", result)
				}
				defer func() {
					resp.Close()
				}()
			}
		}(j)

	}
	wg.Wait()
}

type DataCenterReq struct {
	Producer string `form:"producer" json:"producer" description:"生产"`
	Consumer string `form:"consumer" json:"consumer" description:"消费"`
	Data     []byte `form:"data" json:"data" description:"数据"`
	SendAt   int64  `form:"sendAt" json:"sendAt,string" description:"发送时间"`
}
