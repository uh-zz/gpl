package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}

	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		// FYI: 多くのプログラミング言語の実装は、固定長の関数呼び出しスタックを使っている。
		// スタックの大きさは64KBから2MBまでが普通
		// 固定長スタックは、再帰呼び出しの深さに制限を課す
		// 大きなデータ構造を再帰的に操作する場合にはスタックオーバーフローを避けるために注意深くなる必要あり
		//
		// 対象的に、代表的なGoの実装は、最初は小さく、そして必要に応じてギガバイトの大きさへと拡張される可変長スタックを使っています。
		// これにより、オーバーフローを心配することなく再帰を使える
		// https://ymotongpoo.hatenablog.com/entry/2015/02/23/165341#f-75dbf815
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
