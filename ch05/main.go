package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func main() {

}

// CountWordsAndImages はHTMLドキュメントに対するGETリクエストを行い、
// ドキュメント内に含まれる単語と画像の数を返す
//
// FYI: 名前付きの結果を持つ関数内では、return文のオペランドは省略可能。
func CountWordsAndImages(url string) (words, images int, err error) {

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}

	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) { return }
