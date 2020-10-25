package libs

import (
	"fmt"
	"image/png"
	"os"
	"strconv"

	rn "github.com/asggo/random"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)


func QRcode(numberOfQrcodes,QRwidth,QRheight int) {
	
	
	num := numberOfQrcodes*numberOfQrcodes
	for x := 0; x < num; x++ {
	n,err := rn.Uint16()
	if err != nil {
			fmt.Println(err)
		}

	// Create the barcode
	qrCode, _ := qr.Encode(strconv.Itoa(int(n)), qr.M, qr.Auto)
	fmt.Println(strconv.Itoa(int(n)))

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, QRwidth, QRheight)


	 

	// create the output file
	
	file, _ := os.Create("./out/test"+strconv.Itoa(x)+".png")
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, qrCode)
	}
}