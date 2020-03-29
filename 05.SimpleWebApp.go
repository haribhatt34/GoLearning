package main

import ("fmt"
        "net"/"http")

//w -> writer, can name it anything.
func index_handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is my Home Page")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "The author is Hari Shankar Bhatt !!!")
}

func main() {
    // below HandleFunc handles the mapping i.e. for which path which method has to be called.
    http.HandleFunc("/", index_handler)
    http.HandleFunc("/about/", about_handler)
    http.ListenAndServe(":8000", nil);
}
