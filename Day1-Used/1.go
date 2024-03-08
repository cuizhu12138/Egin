// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func indexHandle(w http.ResponseWriter, req *http.Request) {
// 	fmt.Fprintf(w, "URL.path = %q\n", req.URL.Path)
// }
// func helloHandle(w http.ResponseWriter, req *http.Request) {
// 	for k, v := range req.Header {
// 		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
// 	}
// }

// func main() {
// 	http.HandleFunc("/", indexHandle)
// 	http.HandleFunc("/hello", helloHandle)
// 	log.Fatal(http.ListenAndServe(":9999", nil))
// }
