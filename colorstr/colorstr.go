package colorstr

// the color format information comes from https://en.wikipedia.org/wiki/ANSI_escape_code

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var color = map[string]string{
	"BlackFg":         "30",
	"RedFg":           "31",
	"GreenFg":         "32",
	"YellowFg":        "33",
	"BlueFg":          "34",
	"MagentaFg":       "35",
	"CyanFg":          "36",
	"WhiteFg":         "37",
	"BrightBlackFg":   "90",
	"BrightRedFg":     "91",
	"BrightGreenFg":   "92",
	"BrightYellowFg":  "93",
	"BrightBlueFg":    "94",
	"BrightMagentaFg": "95",
	"BrightCyanFg":    "96",
	"BrightWhiteFg":   "97",
	"End":             "0",
}

// can receive more than one color, foreground color and background color
// cause color format depends on the color number, so the order is not important
// \033[30;45m is the same as \033[45;30m
func Colorize(colorname []string, text string) string {
	var colortext string

	colorcode1, exists := color[colorname[0]]
	if !exists {
		io.WriteString(os.Stdout, "color name is not right\n")
		os.Exit(1)
	}
	colortext = fmt.Sprintf("\033[%sm%s\033[0m", colorcode1, text)

	if len(colorname) == 2 {
		colorcode2, exists := color[colorname[1]]
		if !exists {
			io.WriteString(os.Stdout, "color name is not right\n")
			os.Exit(1)
		}

		colortext = fmt.Sprintf("\033[%s;%sm%s\033[0m", colorcode1, colorcode2, text)
	}

	return colortext
}

func ColorizeRgbFg(rgb, text string) string {
	if len(rgb) != 7 || !strings.HasPrefix(rgb, "#") {
		io.WriteString(os.Stdout, "rgb color not right\n")
		os.Exit(1)
	}

	r, g, b := rgb[1:3], rgb[3:5], rgb[5:7]
	numr, _ := strconv.ParseUint(r, 16, 8)
	numg, _ := strconv.ParseUint(g, 16, 8)
	numb, _ := strconv.ParseUint(b, 16, 8)

	colorizeText := fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", numr, numg, numb, text)
	return colorizeText
}
