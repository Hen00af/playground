package main 
import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-type", "application/json")

	source := map[string]string {
		"Hello": "World",
	}

	// 作成したjsonの情報をLogとしてSTDOUTに出力する。
	io.MultiWriter()

	// json encoder

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}