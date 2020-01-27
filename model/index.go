package model

import (
	"github.com/pquerna/ffjson/ffjson"
	"io/ioutil"
	"net/http"
)

// 蛋卷指数列表：https://danjuanapp.com/djmodule/value-center
const danjuanIndexDetailUrl = "https://danjuanapp.com/djapi/index_eva/detail/"
const danjuanIndexCode = "SH000922"    // 中证红利指数

type IndexInfo struct {
	ID        int      `json:"id"`
	IndexCode string   `json:"index_code"`
	Name      string   `json:"name"`
	PE        float32  `json:"pe"`
	Date      string   `json:"date"`
}

type danjuanResponse struct {
	ResultCode int      `json:"result_code"`
	Data       *IndexInfo `json:"data"`
}

func GetIndexInfo() (*IndexInfo, error) {
	url := getIndexDetailUrl(danjuanIndexCode)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	/**
	if resp.StatusCode != 200 {
		fmt.Printf("Http response error. status code:%d", resp.StatusCode)
	}
	*/

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &danjuanResponse{}
	err = ffjson.Unmarshal(result, ret)
	if err != nil {
		return nil, err
	}

	return ret.Data, err
}

func getIndexDetailUrl(indexCode string) string {
	return danjuanIndexDetailUrl + indexCode
}
