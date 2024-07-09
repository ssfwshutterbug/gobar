package main

import (
	"fmt"
	"gobar/colorstr"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	color, length, ratio := receiveArgs()
	colorlist := []string{color}
	bar := genChar(colorlist, length, ratio)
	io.WriteString(os.Stdout, bar+"\n")
}

func receiveArgs() (string, int64, float64) {
	args := os.Args
	if len(args) < 4 {
		help()
		os.Exit(1)
	}

	color := args[1]

	length, err := strconv.ParseInt(args[2], 10, 64)
	if err != nil {
		fmt.Println("err")
	}

	ratio := args[3]
	var decimal float64
	if strings.HasSuffix(ratio, "%") {
		ratio = strings.TrimRight(ratio, "%")
		decimal, err = strconv.ParseFloat(ratio, 64)
		if err != nil {
			fmt.Println(err)
		}
		decimal = decimal / 100
	} else {
		decimal, err = strconv.ParseFloat(ratio, 64)
		if err != nil {
			fmt.Println(err)
		}
	}

	return color, length, decimal
}

func genChar(barcolor []string, length int64, ratio float64) string {
	var bgcolor = []string{"BlackFg"}
	if ratio > 1 {
		ratio = 1
	}

	firstSeg := int64(float64(length) * ratio)
	secondSeg := length - firstSeg

	var firstchars, colorfirst string
	for i := 0; i < int(firstSeg); i++ {
		firstchars = firstchars + ""
	}

	var secondchars, colorsecond string
	for i := 0; i < int(secondSeg); i++ {
		secondchars = secondchars + ""
	}
	colorsecond = colorstr.Colorize(bgcolor, secondchars)

	percent := " " + strconv.Itoa(int(ratio*100)) + "%"

	// use rgb color or color name
	if strings.HasPrefix(barcolor[0], "#") {
		colorfirst = colorstr.ColorizeRgbFg(barcolor[0], firstchars)
		percent = colorstr.ColorizeRgbFg(barcolor[0], percent)
	} else {
		colorfirst = colorstr.Colorize(barcolor, firstchars)
		percent = colorstr.Colorize(barcolor, percent)
	}

	return colorfirst + colorsecond + percent
}

func help() {
	io.WriteString(os.Stdout, `
gobar <color> <bar length> <ratio>
gobar BrightBlueFg 60 35%
gobar "#982374" 60 0.35

available color:

	BlackFg
	RedFg
	GreenFg
	YellowFg
	BlueFg
	MagentaFg
	CyanFg
	WhiteFg
	BrightBlackFg
	BrightRedFg
	BrightGreenFg
	BrightYellowFg
	BrightBlueFg
	BrightMagentaFg
	BrightCyanFg
	BrightWhiteFg
        `+"\n")
}
