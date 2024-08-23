package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/marktsarkov/notesapp/internal/config"
)

func main() {
    cfg, err := config.LoadConfig("configs/dev.yaml")
    if err != nil {
        log.Fatalf("Ошибка загрузки конфигурации: %v", err)
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "alloha from %s!", cfg.App.Name)
    })

    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.App.Port), nil))
}
