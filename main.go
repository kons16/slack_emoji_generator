package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/fogleman/gg"
)

func main() {
	rgbs := [][]int{
		{236, 113, 161},
		{0, 187, 0},
		{27, 170, 198},
		{5, 201, 139},
		{231, 159, 12},
	}

	file, err := os.Open("data/test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	size := 128

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		moji := strings.Split(line, " ")[0]
		pngName := strings.Split(line, " ")[1]

		lineMojis := strings.Split(moji, "/")
		lineCnt := len(lineMojis)                   // 行数
		cnt := utf8.RuneCountInString(lineMojis[0]) // 1行にある文字数

		dc := gg.NewContext(size, size)

		// 文字色を乱数で選択
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(5)

		dc.SetRGB(float64(rgbs[n][0]), float64(rgbs[n][1]), float64(rgbs[n][2]))

		if err := dc.LoadFontFace("rounded-mplus-1c-black.ttf", float64(size/cnt)); err != nil {
			fmt.Println("load font error")
		}

		for i := 0; i < lineCnt; i++ {
			strImg := lineMojis[i]
			dc.DrawStringAnchored(strImg, float64(size), float64((size/cnt)*(i+1)), 1, 0)
		}

		dc.SavePNG("out/" + pngName + ".png")
	}

	if err = scanner.Err(); err != nil {
		fmt.Println(err)
	}

}
