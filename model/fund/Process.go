package fund

import (
	"fmt"
	"github.com/pquerna/ffjson/ffjson"
)

type fundInfo struct {
	Name   string  `json:"name"`
	PE     float32 `json:"pe"`
}

func Process() {
	data := "{\"name\":\"Bonus Fund\",\"pe\":7.94}"
	decodedData := &fundInfo{}
	err := ffjson.Unmarshal([]byte(data), decodedData)
	if err != nil {
		fmt.Printf("process failed. err:%v", err)
	} else {
		fmt.Printf("process succ. fund name:%s | fund PE:%.2f \n", decodedData.Name, decodedData.PE)
	}
}
