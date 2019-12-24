package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"strings"
	"time"
)

var colors = []color.Attribute{color.FgBlue, color.FgGreen, color.FgRed, color.FgYellow, color.FgMagenta}

func tree(colorStart int, input string) string{
	runes := []rune(input)
	middle := len(runes) / 2
	out := fmt.Sprintf(strings.Repeat(" ", middle),
		"\u2605\n",
		strings.Repeat(" ", middle),
		"|\n")
	offset := middle
	for i := 1; i <= len(runes); i +=2 {
		out += strings.Repeat(" ", offset)
		offset -= 1
		changeColor := color.New(colors[colorStart]).SprintfFunc()
		if colorStart + 1 >= len(colors) {
			colorStart = 0
		} else {
			colorStart++
		}
		out += changeColor(string(runes[:i]))
		out += "\n"
	}
	for i := 0; i < 3; i++ {
		out += strings.Repeat(" ", middle - 1)
		out += "|||\n"
	}
	return out
}

func treeToConsole(input string) {
	colorStart := 0
	out := ""
	for {
		out = tree(colorStart, input)
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Printf("%s", out)
		time.Sleep(800 * time.Millisecond)
		if colorStart + 1 >= len(colors) {
			colorStart = 0
		} else {
			colorStart++
		}
		out = ""
	}
}

func main() {
	treeToConsole("Димон сосок верблюда!")
}