package main

import (
    "fmt"
    "net/http"
    "os/exec"
    "log"
    "github.com/rs/cors"
    "strings"
)
var id = 1
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
    w.Header().Set("Pragma", "no-cache")
    w.Header().Set("Expires", "0")
    if (len(string(r.URL.Path[1:])) == 11 && string(r.URL.Path[1:]) != "favicon.ico" ){
        log.Printf("Got request for : %s", r.URL.Path[1:])
        err := exec.Command("youtube-dl", "--extract-audio", "--audio-format", "mp3", "--output", "%(title)s.%(ext)s", "--restrict-filenames", r.URL.Path[1:]).Run()
        if (err != nil) {
            log.Printf("Error occurred processing URL : %s", r.URL.Path[1:])
            fmt.Fprintf(w, "Error")
        }else{
            out, err2 := exec.Command("youtube-dl", "--extract-audio", "--audio-format", "mp3", "--output", "%(title)s.%(ext)s", "--restrict-filenames", "--get-filename", r.URL.Path[1:]).Output()
          if (err2 != nil){
            log.Panic(err2)
            fmt.Fprintf(w, "Error Type 2")
          } else{
            filename := strings.Replace(strings.Replace(strings.TrimSuffix(string(out),"\n"),".webm", ".mp3",1),".m4a", ".mp3",1)
            w.Header().Set("Content-Disposition", "attachment; filename=" + filename)
            w.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
	        w.Header().Set("Content-Type", "audio/mp3")
            fmt.Println(filename)
            http.ServeFile(w, r, filename)
            id++
        }
          }
    } else{
        log.Printf("Bad URL : %s", r.URL.Path[1:])
        fmt.Fprintf(w, "Bad URL")
    }
}
func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", handler)
    log.Printf("Youtube MP3 Download Backend Server Started")
    handler := cors.Default().Handler(mux)
    http.ListenAndServe(":3000", handler)
}
