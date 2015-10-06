package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "time"
    "strconv"

    "github.com/gorilla/mux"
)

func handle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    n, err := strconv.Atoi(vars["wait"])
    if err != nil {
        log.Fatal(err)
    }
    time.Sleep(time.Duration(n) * time.Second)

    switch vars["status"] {
    case "200":
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(http.StatusText(http.StatusOK)))
    case "300":
        w.WriteHeader(http.StatusMultipleChoices)
        w.Write([]byte(http.StatusText(http.StatusMultipleChoices)))
    case "400":
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(http.StatusText(http.StatusBadRequest)))
    case "500":
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
    }
}

func main() {
    r := mux.NewRouter().StrictSlash(true)
    r.HandleFunc("/{status}/{wait}", handle).Methods("GET")

    http.Handle("/", r)
    envPort := os.Getenv("PORT")
    if len(envPort) == 0 {
        envPort = "11011"
    }
    listenOn := fmt.Sprintf(":%s", envPort)
    log.Printf("listening on: %s", listenOn)
    log.Fatal(http.ListenAndServe(listenOn, nil))
}
