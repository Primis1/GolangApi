// package main

// import (
// 	"log/slog"
// 	"os"
// 	"url-shortener/internal/config"
// )

// const (
// 	envLocal = "local"
// 	envDev   = "dev"
// 	envProd  = "prod"
// )

// func main() {
//     cfg := config.MustLoad()

//     log := setupLogger(cfg.Env)
//     log = log.With(slog.String("env", cfg.Env)) // к каждому сообщению будет добавляться поле с информацией о текущем окружении

//     log.Info("initializing server", slog.String("address", cfg.Address)) // Помимо сообщения выведем параметр с адресом
//     log.Debug("logger debug mode enabled")
// }

// func setupLogger(env string) *slog.Logger {
//     var log *slog.Logger

//     switch env {
//     case envLocal:
//         log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
//     case envDev:
//         log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
//     case envProd:
//         log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
//     }

//     return log
// }

package main

import (
	"log"
	"net/http"
) 

func home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/" {
        http.NotFound(w, r)
        return 
    }
    w.Write([]byte("Hello my server"))
}

func snippet(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Snippet"))
}

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/", home)
    mux.HandleFunc("/snip", snippet)

    log.Print("Starting server on :4000")

    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}