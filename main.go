package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fogleman/gg"
)

func main() {
	// rgbs := [][]int{
	// 	{236, 113, 161},
	// 	{0, 187, 0},
	// 	{27, 170, 198},
	// 	{5, 201, 139},
	// 	{231, 159, 12},
	// }

	file, err := os.Open("data/test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		moji := strings.Split(line, " ")[0]
		//pngName := strings.Split(line, " ")[1]

		lineMojis := strings.Split(moji, "/")
		lineCnt := len(lineMojis)
		fmt.Println(lineCnt)

		dc := gg.NewContext(128, 128)
		dc.SetRGB(1, 1, 1)

		for _, str := range lineMojis {
			fmt.Println(str)
			dc.DrawStringAnchored(str, float64(size*len([]rune(message))/2), float64(size/2), 0.5, 0.5)
		}

		// dc.SavePNG("out/" + pngName + ".png")

	}

	if err = scanner.Err(); err != nil {
		fmt.Println(err)
	}

}
