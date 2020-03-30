package urlimagediff

import (
	"image"
	"io/ioutil"
	"os"
	"path/filepath"

	diffimage "github.com/murooka/go-diff-image"
	"github.com/sclevine/agouti"
)

// Diff 二つのURLの画面イメージが同じかどうかを確認する
func Diff(url1, url2 string) (image.Image, bool, error) {
	dir, err := ioutil.TempDir("./", "urlimagediff_images")
	if err != nil {
		return nil, false, err
	}
	defer os.RemoveAll(dir)
	options := agouti.ChromeOptions(
		"args", []string{
			"--headless",
			"--disable-gpu", // 暫定的に必要らしいです。
			"--lang=ja",
		})
	driver := agouti.ChromeDriver(options)
	driver.Start()
	defer driver.Stop()
	page, err := driver.NewPage()
	if err != nil {
		return nil, false, err
	}

	if err := page.Navigate(url1); err != nil {
		return nil, false, err
	}
	var width int
	var height int
	page.RunScript(`return document.body.scrollWidth`, nil, &width)
	page.RunScript(`return document.body.scrollHeight`, nil, &height)
	page.Size(width, height)
	page.Screenshot(filepath.Join(dir, "1.png"))
	if err := page.Navigate(url2); err != nil {
		return nil, false, err
	}
	page.RunScript(`return document.body.scrollHeight`, nil, &height)
	page.Size(width, height)
	page.Screenshot(filepath.Join(dir, "2.png"))
	f, _ := os.Open(filepath.Join(dir, "1.png"))
	img1, _, err := image.Decode(f)
	f.Close()
	f, _ = os.Open(filepath.Join(dir, "2.png"))
	img2, _, err := image.Decode(f)
	f.Close()
	diff := diffimage.DiffImage(img1, img2)
	if img1.Bounds().Dx() == diff.Bounds().Dx() && img1.Bounds().Dy() == diff.Bounds().Dy() &&
		img2.Bounds().Dx() == diff.Bounds().Dx() && img2.Bounds().Dy() == diff.Bounds().Dy() {
		return diff, true, nil
	}
	return diff, false, nil
}
