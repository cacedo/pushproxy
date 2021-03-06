package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "fmt"
    "github.com/mitsuse/pushbullet-go"
	"github.com/mitsuse/pushbullet-go/requests"
    "os"
)

var token = os.Getenv("PB_TOKEN")

type res struct {
  Message string
}

func Alert(w http.ResponseWriter, r *http.Request) {

  fmt.Println("Alert!!")
  decoder := json.NewDecoder(r.Body)
  var rr res
  err := decoder.Decode(&rr)

  if err != nil {
		panic(err)
  }

  pb := pushbullet.New(token)
	// Create a push. The following codes create a note, which is one of push types.
	n := requests.NewNote()
	n.Title = "Alert!"
	n.Body = rr.Message

	if _, err := pb.PostPushesNote(n); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		return
	}


  fmt.Println(rr.Message)
}

// our main function
func main() {
  if token == "" {
		panic("PB_TOKEN env is empty!")
  }

    router := mux.NewRouter()
    router.HandleFunc("/alert", Alert).Methods("POST")
    log.Fatal(http.ListenAndServe(":8008", router))
}
