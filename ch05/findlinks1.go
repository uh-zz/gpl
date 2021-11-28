// findlinks1 は標準入力から読み込まれたHTMLドキュメント内のリンクを表示する
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}

}

// visit は、n内で見つかったリンクを一つ一つlinksへ追加し、その結果を返す
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	// ノードnに対するツリーを加工するために、visitはFirstChildリンクリストに保持されているnのChildに対して、自身を再帰的に呼び出す
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
