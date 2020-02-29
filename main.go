package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hello Jake")
	fmt.Println("I'm sleepy...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte(`
		<html>
		<body>hello from main.go</body>
		</html>
		`)); err != nil {
			panic(err)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
