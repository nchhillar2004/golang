package main

import(
    "log"
    "net/http"
)

// sarting a live server in golang and serving a .html file
func web(){
    log.Println("server stated on port :8080")
    var fs = http.FileServer(http.Dir("./src/"))
    http.Handle("/", fs)

    err := http.ListenAndServe(":8080", nil)
    if err!=nil {
        log.Fatal(err)
    }
}
