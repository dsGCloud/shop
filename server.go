package main
import (
 "fmt"
 "net/http"
 "os"
)
func handler(w http.ResponseWriter, r *http.Request) {
 fmt.Fprintf(w, "<html><head><title>Cloud Run Demo</title></head><body><h1>Greetings from Cloud Run !!!</h1></body></html>\n")
}
func main() {
 port := os.Getenv("PORT")
 if port == "" {
  port = "8080"
 }
 http.HandleFunc("/", handler)
 http.ListenAndServe(":"+port, nil)
}
