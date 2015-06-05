package main

import (
    "github.com/lxn/walk"
    . "github.com/lxn/walk/declarative"
	"strconv"
)

func main() {
    var inTE, outTE *walk.TextEdit

    MainWindow{
        Title:   "时间戳转换器",
        MinSize: Size{600, 400},
        Layout:  VBox{},
        Children: []Widget{
            HSplitter{
                Children: []Widget{
                    TextEdit{AssignTo: &inTE},
                    TextEdit{AssignTo: &outTE, ReadOnly: true},
                },
            },
            PushButton{
                Text: "转换",
                OnClicked: func() {
					num,_ := strconv.ParseInt(inTE.Text(),0,64)
					outTE.SetText(ToDate(num))
                },
            },
        },
    }.Run()
}
