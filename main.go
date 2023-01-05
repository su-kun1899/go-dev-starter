package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	err := http.ListenAndServe(
		fmt.Sprintf(":%s", getEnv("PORT", "18080")),
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello, world!")
		}))

	if err != nil {
		fmt.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}

// getEnv は指定した key の環境変数を取得します
// 環境変数が未設定の場合 fallback の値を返します
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
