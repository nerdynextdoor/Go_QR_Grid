package libs

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"strconv"

	gim "github.com/ozankasikci/go-image-merge"
)

func CreatingQRgrid(numberOfQRcodes, QRWidth, QRHeight, bgWidth, bgHeight int, OutputPath string) {

	num := numberOfQRcodes
	CreatingBackground(bgWidth, bgHeight)
	QRcode(num, QRWidth, QRHeight)

	counter := num * num
	grids := []*gim.Grid{}
	backgroundimage := "./out/image.png"
	var QRimage string

	setXoffset, setYoffset := settingOffset(QRWidth, QRHeight, bgWidth, bgHeight)

	fmt.Println("\nDrawing Grid")
	for x := 0; x < counter; x++ {

		QRimage = "./out/test" + strconv.Itoa(x) + ".png"
		in := []*gim.Grid{

			{
				ImageFilePath: backgroundimage,
				//BackgroundColor: color.White,
				// these grids will be drawn on top of the first grid
				Grids: []*gim.Grid{
					{
						ImageFilePath: QRimage,
						OffsetX:       setXoffset, OffsetY: setYoffset,
					},
				},
			},
		}
		grids = append(grids, in...)

		//fmt.Println(fn)
		fmt.Print(".")

	}

	fmt.Println("\nCreating Grid......")
	rgba, err := gim.New(grids, num, num,
		gim.OptBaseDir("./"),
	).Merge()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Done!! \nSaving Image......")
	// save the output to jpg or png
	file, err := os.Create(OutputPath + "/QRgrid.png")
	err = jpeg.Encode(file, rgba, &jpeg.Options{Quality: 80})
	err = png.Encode(file, rgba)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//removing Temp Output folder
	err = os.RemoveAll("out")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Output image has been saved but program could not delete the Output temp folder")
		os.Exit(1)
	}
	fmt.Println("Done!! \nYou can now open the new image called: QRgrid.png")

}

func settingOffset(QRWidth, QRHeight, bgWidth, bgHeight int) (int, int) {
	x := 0
	y := 0
	if (bgWidth - QRWidth) < 0 {
		x = 0
	} else {
		x = (bgWidth - QRWidth) / 2
	}

	if (bgHeight - QRHeight) < 0 {
		y = 0
	} else {
		y = (bgHeight - QRHeight) / 2
	}

	return x, y
}
