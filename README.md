# go-url-image-diff
二つのURLのスクリーンショットを撮って比較し、差異がある場合は表示する

# Installation

Macの場合
1. SeleniumとChromedriverを入れる
```bash
brew install phantomjs
brew install chromedriver
brew install selenium-server-standalone
```

2. バイナリを取得
```bash
go get -u github.com/loadoff/go-url-image-diff/...
```
 
# Usage

```bash
urlimagediff -o output.png http://sample1.com http://sample2.com
```

# Author

* Yu wasaki
* loadoff inc.
* yuiwasaki@loadoff.jp

# License

MIT
