package main

import (
	"fmt"
	"github.com/whutwxn/dxf-go/document"
	"github.com/whutwxn/dxf-go/entities"
	"log"
	"os"
)

func main() {

	file, err := os.Open("E:/whut_wxn_doc/Work/项目/跨平台/dxf-go/example/南刘庄疏浚区背景图(1).dxf")
	if err != nil {
		log.Fatal(err)
	}

	doc, err := document.DxfDocumentFromStream(file)
	if err != nil {
		log.Fatal(err)
	}
	//entities.Polyline
	for _, block := range doc.Blocks {
		for _, entity := range block.Entities {
			//if polyline, ok := entity.(*entities.Polyline); ok {
			//	//fmt.Println(polyline.ColorName)
			//	fmt.Println(polyline.TrueColor)
			//	// process polyline here...
			//} else if lwpolyline, ok := entity.(*entities.LWPolyline); ok {
			//	// process lwpolyline here...
			//	//fmt.Println(lwpolyline.ColorName)
			//	fmt.Println(lwpolyline.TrueColor)
			//} else
			if text, ok := entity.(*entities.Text); ok {
				if text.Value == "分区A" {
					fmt.Println(text.Value, text.TrueColor)
				}
				//fmt.Println(text.Value, text.TrueColor)
			}
			//...
		}
	}
}
