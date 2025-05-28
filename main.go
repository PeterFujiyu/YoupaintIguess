package main

import (
    "io/fs"
    "log"
    "net/http"
)

func main() {
    fsys, err := fs.Sub(content, "public") // content 来自 static.go 的 embed
    if err != nil {
        log.Fatal(err)
    }
    http.Handle("/", http.FileServer(http.FS(fsys)))

    // static JSON like themes/words.json
    http.HandleFunc("/themes/words.json", func(w http.ResponseWriter, r *http.Request) {
        data, err := content.ReadFile("themes/words.json")
        if err != nil {
            http.NotFound(w, r)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        w.Write(data)
    })

    log.Println("Server running at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}