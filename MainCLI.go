package main

import (
	"flag"
	"fmt"
	"os"

	l "./libs"
)

func main() {

	numberOfQRcodes := flag.Int("QRnum", 0, "Number of QRcodes")
	QRWidthAndHeight := flag.Int("QRsize", 0, "Size of QRcodes Width and Height in Pixels")
	bgWidthAndHeight := flag.Int("BGsize", 0, "Size of Background Image Width and Height in Pixels")
	OutputPath := flag.String("Output", "./", "This is output path of the Image, Default path is the program folder")

	flag.Parse()

	if *numberOfQRcodes == 0 || *QRWidthAndHeight == 0 || *bgWidthAndHeight == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *QRWidthAndHeight < 30 || *bgWidthAndHeight < 30 {
		fmt.Printf("the value of background or QRcode size can not be less than 30 pixels \n")
		os.Exit(1)
	}

	l.CreatingQRgrid(*numberOfQRcodes, *QRWidthAndHeight, *QRWidthAndHeight, *bgWidthAndHeight, *bgWidthAndHeight, *OutputPath)
}
