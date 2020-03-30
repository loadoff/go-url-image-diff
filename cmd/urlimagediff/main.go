package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"

	urlimagediff "github.com/loadoff/go-url-image-diff"
)

func main() {
	var output string
	flag.StringVar(&output, "output", "diff.png", "output filename")
	flag.StringVar(&output, "o", "diff.png", "output filename")
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("usage: imagediff [<option>...] <url1> <url2>")
		os.Exit(1)
	}
	url1 := args[0]
	url2 := args[1]
	diff, isSame, err := urlimagediff.Diff(url1, url2)
	if err != nil {
		fmt.Printf("\x1b[31mエラーが発生しました。[%v]\x1b[0m\n", err)
		return
	}

	f, _ := os.OpenFile(output, os.O_WRONLY|os.O_CREATE, 0644)
	png.Encode(f, diff)
	f.Close()

	if isSame {
		fmt.Printf("\x1b[31m[%s]と[%s]は同じ画面です。\x1b[0m\n", url1, url2)
		return
	}
	fmt.Printf("\x1b[31m[%s]と[%s]は差異があります。\x1b[0m\n", url1, url2)
	return
}
