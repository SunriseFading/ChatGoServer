package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Message struct {
	Username  string
	Message   string
	Timestamp string
}

func main() {
	r := chi.NewRouter()
	fmt.Println("Hello, world!")
	messages := make([]Message, 0)
	r.Post("/api/Messanger", func(w http.ResponseWriter, r *http.Request) {
		mes := Message{}
		raw, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(raw, &mes)
		fmt.Println(mes)
		messages = append(messages, mes)
	})
	r.Get("/api/Messanger", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		id_num, _ := strconv.Atoi(id)
		if id_num >= 0 && id_num < len(messages) {
			msg := messages[id_num]
			data, _ := json.Marshal(msg)
			w.Write(data)
		}
	})
	http.ListenAndServe(":5000", r)
}
