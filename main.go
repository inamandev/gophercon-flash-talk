package main

import (
	"embed"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/webview/webview"
)

type IncrementResult struct {
	Count uint `json:"count"`
}

//go:embed static
var static embed.FS

const PORT = 3000

func main() {
	// var count uint = 0
	mutex := http.NewServeMux()
	mutex.Handle("/", http.FileServer(http.FS(static)))
	go runServer(mutex)
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("SiteLish")
	w.SetSize(1080, 780, webview.HintNone)
	w.Navigate(fmt.Sprintf(`http://localhost:%d/static`, PORT))
	/* w.Bind("increment", func() IncrementResult {
		count++
		return IncrementResult{Count: count}
	})
	w.SetHtml(html) */
	w.Run()
}

func runServer(app *http.ServeMux) error {
	log.Printf(`Service Will be running on port %d `, PORT)
	if ok := http.ListenAndServe(fmt.Sprintf(`:%d`, PORT), app); ok != nil {
		log.Println(ok)
		return ok
	}
	return errors.New("web app terminated")
}
