package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zuolicong/lucky-cat/model"
	"github.com/zuolicong/lucky-cat/tui"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	fmt.Println("Use lucky cat and make your fortune!")
	indexInfo, err := model.GetIndexInfo()
	if err != nil {
		fmt.Printf("Failed to get index info. err:%s", err.Error())
	}

	tui.ShowPE(indexInfo)
}
