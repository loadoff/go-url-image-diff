# go-url-image-diff
二つのURLのスクリーンショットを撮って比較し、差異がある場合は表示する

# Demo
```bash
urlimagediff -o output.png https://github.com/loadoff/go-url-image-diff/blob/develop/README.md https://github.com/loadoff/go-url-image-diff/blob/master/README.md
```
![output](https://user-images.githubusercontent.com/22957487/77889060-35265c00-72a8-11ea-8601-198d6c9cd3cd.png)

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
