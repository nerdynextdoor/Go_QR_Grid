package libs

import (
	"image"
	//"image/color"
	"image/png"
	"log"
	"os"
)

func CreatingBackground(bgWidth, bgHeight int) { //Creates background image for the grid cell
	//const width, height = 500, 500

	// Create a colored image of the given width and height.
	img := image.NewNRGBA(image.Rect(0, 0, bgWidth, bgHeight))

	if _, err := os.Stat("./out"); os.IsNotExist(err) {
		os.Mkdir("./out", os.ModeDir)
	}
	f, err := os.Create("./out/image.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}