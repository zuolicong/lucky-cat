package fund

import (
	"fmt"
	"github.com/pquerna/ffjson/ffjson"
	"io/ioutil"
	"net/http"
)

const danjuanDetailUrl = "https://danjuanapp.com/djapi/index_eva/detail/"
const danjuanIndexID = "SH000922"    // 中证红利指数

type fundInfo struct {
	ID        int      `json:"id"`
	IndexCode string   `json:"index_code"`
	Name      string   `json:"name"`
	PE        float32  `json:"pe"`
	Date      string   `json:"date"`
}

type danjuanResponse struct {
	ResultCode int      `json:"result_code"`
	Data       fundInfo `json:"data"`
}

func Process() {
	url := getUrl(danjuanIndexID)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Http request error:%s", err.Error())
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("Http response error. status code:%d", resp.StatusCode)
	}

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Get response body failed. err:%s", err.Error())
		return
	}
	ret := &danjuanResponse{}
	err = ffjson.Unmarshal(result, ret)
	if err != nil || ret.ResultCode != 0 {
		fmt.Printf("Get fund info failed. err:%v | result code:%d", err, ret.ResultCode)
	} else {
		fmt.Printf("Process succ. fund name:%s | date:%s | fund PE:%.2f \n", ret.Data.Name, ret.Data.Date, ret.Data.PE)
	}
}

func getUrl(indexID string) string {
	return danjuanDetailUrl + indexID
}
