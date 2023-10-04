package main

import (
	"log"

	"github.com/ittyi/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")

	// ListenAndServe は TCP ネットワーク アドレス srv.Addr をリッスンし、その後
	// Serve を呼び出して、受信接続のリクエストを処理します。
	// 受け入れられた接続は、TCP キープアライブを有効にするように構成されます。
	//
	// srv.Addr が空白の場合、「:http」が使用されます。
	//
	// ListenAndServe は常に非 nil エラーを返します。 シャットダウンまたはクローズ後、
	// 返されるエラーは ErrServerClosed です。
	log.Fatal(srv.ListenAndServe())
}
