package main

import (
	"fmt"
	"os"

	"github.com/branogarbo/imgcli/util"
)

var (
	pixelString string
	err         error
)

func main() {
	options := util.OutputConfig{
		Src:          "https://www.researchgate.net/profile/John-Hubbard-6/publication/228540477/figure/fig1/AS:669410264088623@1536611252979/The-Mandelbrot-set-M-The-two-main-conjectures-are-the-following-MLC-The-set-M-is.png",
		Dst:          "./prints/mandelbrotHuge.txt",
		OutputMode:   "ascii",
		AsciiPattern: " .:-=+*#%@",
		OutputWidth:  2000,
		IsUseWeb:     true,
		IsInverted:   true,
		IsPrinted:    false,
		IsSaved:      true,
		IsQuiet:      false,
	}

	pixelString, err = util.OutputImage(options)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

//	fmt.Println(pixelString)
}
