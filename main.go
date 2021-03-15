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
		Src:          "./images/portrait.jpg",
		Dst:          "./prints/portrait.txt",
		OutputMode:   "ascii",
		AsciiPattern: " .:-=+*#%@",
		OutputWidth:  150,
		IsUseWeb:     false,
		IsInverted:   false,
		IsPrinted:    false,
		IsSaved:      true,
		IsQuiet:      true,
	}

	pixelString, err = util.OutputImage(options)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(pixelString)
}
