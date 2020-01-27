package tui

import (
	"fmt"
	"github.com/rivo/tview"
	"github.com/zuolicong/lucky-cat/model"
)

func ShowPE(indexInfo *model.IndexInfo) {
	desc := fmt.Sprintf("Process succ. index name:%s | date:%s | index PE:%.2f \n", indexInfo.Name, indexInfo.Date, indexInfo.PE)
	app := tview.NewApplication()
	list := tview.NewList().
		AddItem("List item 1", desc, 'a', nil).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})
	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}
