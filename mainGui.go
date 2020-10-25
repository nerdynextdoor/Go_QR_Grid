package main

import (
	"strconv"

	l "./libs"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"

	"log"
)

func main() {
	var numberOfQRcodes, QRWidth, QRHeight, bgWidth, bgHeight *walk.TextEdit

	//txtBoxMinWidth := 100
	//txtBoxMinHeight := 20
//
	//txtBoxMaxWidth := 100
	//txtBoxMaxHeight := 20

	MainWindow{
		Title:   "QRProject",
		MinSize:  Size{200, 200},
		Size:     Size{200, 200},
		Layout:   VBox{MarginsZero: true},
		Children: []Widget{
			VSplitter{
				Children: []Widget{
					HSplitter{
						MaxSize: Size{200,20}, MinSize: Size{200,20},
							Children: []Widget{ 
								TextLabel{Text: "  number Of QR codes",TextAlignment: AlignHNearVCenter, MaxSize: Size{130,15}, MinSize: Size{130,15}},
								TextEdit{AssignTo: &numberOfQRcodes , MaxSize: Size{50,15}, MinSize: Size{50,15}},
							},
					
					},

					HSplitter{
						MaxSize: Size{200,20}, MinSize: Size{200,20},
							Children: []Widget{ 
								TextLabel{Text: "  QR Width", TextAlignment: AlignHNearVCenter, MaxSize: Size{130,15}, MinSize: Size{130,15}},
								TextEdit{AssignTo: &QRWidth , MaxSize: Size{50,15}, MinSize: Size{50,15}},
							},
					
					},

					HSplitter{
						MaxSize: Size{200,20}, MinSize: Size{200,20},
							Children: []Widget{ 
								TextLabel{Text: "  QR Height", TextAlignment: AlignHNearVCenter, MaxSize: Size{130,15}, MinSize: Size{130,15}},
								TextEdit{AssignTo: &QRHeight , MaxSize: Size{50,15}, MinSize: Size{50,15}},
							},
					
					},

					HSplitter{
						MaxSize: Size{200,20}, MinSize: Size{200,20},
							Children: []Widget{ 
								TextLabel{Text: "  Background Image Width", TextAlignment: AlignHNearVCenter, MaxSize: Size{130,20}, MinSize: Size{130,20}},
								TextEdit{AssignTo: &bgWidth , MaxSize: Size{50,15}, MinSize: Size{50,15}},
							},
					
					},

					HSplitter{
						MaxSize: Size{200,20}, MinSize: Size{200,20},
							Children: []Widget{ 
								TextLabel{Text: "  Background Image Height",TextAlignment: AlignHNearVCenter, MaxSize: Size{130,20}, MinSize: Size{130,20}},
								TextEdit{AssignTo: &bgHeight , MaxSize: Size{50,15}, MinSize: Size{50,15}},
							},
					
					},
					
					
					
					//TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			PushButton{
				Text: "Create QR Grid",
				OnClicked: func() {
					numQR, err := strconv.Atoi(numberOfQRcodes.Text())
					if err != nil {
						log.Fatal(err)
					}
					QRw, err := strconv.Atoi(QRWidth.Text())
					if err != nil {
						log.Fatal(err)
					}
					QRh, err := strconv.Atoi(QRHeight.Text())
					if err != nil {
						log.Fatal(err)
					}
					BGw, err := strconv.Atoi(bgWidth.Text())
					if err != nil {
						log.Fatal(err)
					}
					BGh, err := strconv.Atoi(bgHeight.Text())
					if err != nil {
						log.Fatal(err)
					}

					l.CreatingQRgrid(numQR ,QRw ,QRh , BGw, BGh)
				//	outTE.SetText(strings.ToUpper(inTE.Text()))
				},
			},
		},
	}.Run()
}