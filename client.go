// client.pyのコードをGoに移行させる

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// サーバーに接続
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// ユーザー名を入力
	fmt.Print("ユーザー名を入力してください: ")
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	username := stdin.Text()

	// ユーザー名をサーバーに送信
	fmt.Fprintf(conn, "%s\n", username)

	// サーバーからのメッセージを受信するgoroutine
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	// 標準入力から読み込んだメッセージをサーバーに送信するgoroutine
	for {
		stdin.Scan()
		text := stdin.Text()
		fmt.Fprintf(conn, "%s\n", text)
		if strings.Contains(text, "bye") {
			fmt.Println("bye!")
			break
		}
	}
}