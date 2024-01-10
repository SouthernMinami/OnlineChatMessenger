// server.pyをgoに移行させる

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// サーバーを起動
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	// クライアントからの接続を待ち続ける
	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// クライアントからのメッセージを受信するgoroutine
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	// 標準入力から読み込んだメッセージをクライアントに送信するgoroutine
	for {
		stdin := bufio.NewScanner(os.Stdin)
		stdin.Scan()
		text := stdin.Text()
		fmt.Fprintf(conn, "%s\n", text)
		if strings.Contains(text, "bye") {
			fmt.Println("bye!")
			break
		}
	}
}