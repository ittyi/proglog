package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func NewHTTPServer(addr string) *http.Server {
	println("addr:", addr)
	httpsrv := newhttpServer()
	println("httpsrv:", httpsrv)
	if httpsrv != nil {
		println("httpsrv.Log:", httpsrv.Log)
		println("httpsrv.Log.records:", httpsrv.Log.records)
	}

	// NewRouter は新しいルーター インスタンスを返します。
	r := mux.NewRouter()
	println("r:", r)
	// HandleFunc は、URL パスのマッチャーを使用して新しいルートを登録します。
	// Route.Path() と Route.HandlerFunc() を参照してください。
	r.HandleFunc("/", httpsrv.handleProduce).Methods("POST")
	r.HandleFunc("/", httpsrv.handleConsume).Methods("GET")
	// Methods は HTTP メソッドのマッチャーを追加します。
	// 一致する 1 つ以上のメソッドのシーケンスを受け入れます。例:
	// 「GET」、「POST」、「PUT」。

	// サーバーは、HTTP サーバーを実行するためのパラメータを定義します。
	// Server の値 0 は有効な構成です。
	return &http.Server{
		// Addr はオプションで、リッスンするサーバーの TCP アドレスを指定します。
		// 「ホスト:ポート」の形式。 空の場合、「:http」(ポート 80) が使用されます。
		// サービス名は RFC 6335 で定義され、IANA によって割り当てられます。
		// アドレス形式の詳細については、net.Dial を参照してください。
		Addr:    addr,
		Handler: r, // 呼び出すハンドラー、nil の場合は http.DefaultServeMux
	}
}

type httpServer struct {
	Log *Log
}

func newhttpServer() *httpServer {
	return &httpServer{
		Log: NewLog(),
	}
}

type ProduceRequest struct {
	Record Record `json:"record"`
}

type ProduceResponse struct {
	Offset uint64 `json:"offset"`
}

type ConsumeRequest struct {
	Offset uint64 `json:"offset"`
}

type ConsumeResponse struct {
	Record Record `json:"record"`
}

func (s *httpServer) handleProduce(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req ProduceRequest
	// Body はリクエストの本体です。
	//
	// クライアントリクエストの場合、nil ボディはリクエストに何も含まれていないことを意味します。
	// 本文（GET リクエストなど）。 HTTP クライアントのトランスポート
	// は Close メソッドを呼び出す役割を果たします。
	//
	// サーバーリクエストの場合、リクエストボディは常に非 nil です
	// ただし、ボディが存在しない場合はすぐに EOF を返します。
	// サーバーはリクエスト本文を閉じます。 ServeHTTP
	// ハンドラーはその必要はありません。
	//
	// 本体では、Close と同時に Read を呼び出すことができる必要があります。
	// 特に、Close を呼び出すと、待機中の読み取りのブロックが解除されるはずです
	// 入力用。
	println("r.Body:", r.Body)
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	println("req.Record.Offset:", req.Record.Offset)
	println("req.Record.Value:", string(req.Record.Value))

	off, err := s.Log.Append(req.Record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	println("off:", off)

	res := ProduceResponse{Offset: off}
	println("res.Offset:", res.Offset)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s httpServer) handleConsume(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req ConsumeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	record, err := s.Log.Read(req.Offset)
	if err == ErrOffsetNotFound {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := ConsumeResponse{Record: record}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
