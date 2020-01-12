package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zuolicong/lucky-cat/model/fund"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	fmt.Println("Use lucky cat and get your fortune")
	fund.Process()
}
