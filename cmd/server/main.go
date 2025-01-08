package main

import (
	"encoding/json"
	"hajin-chung/deps.me/internal/generate"
	"hajin-chung/deps.me/internal/upload"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type PostList struct {
	Posts []string `json:"posts"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/list", HandleList)
	mux.HandleFunc("/read", HandleRead)
	mux.HandleFunc("/write", HandleWrite)
	mux.HandleFunc("/delete", HandleDelete)
	mux.HandleFunc("/publish", HandlePublish)
	log.Fatal(http.ListenAndServe(":80", mux))
}

func HandleList(w http.ResponseWriter, r *http.Request) {
	entries, err := os.ReadDir("posts")
	if err != nil {
		log.Printf("error on reading dir ./posts\n%s\n", err)
		w.WriteHeader(500)
		return
	}

	posts := []string{}
	for _, entry := range entries {
		posts = append(posts, entry.Name())
	}

	post_list := PostList{Posts: posts}
	res, err := json.Marshal(post_list)
	if err != nil {
		log.Printf("error on encoding post list json\n%s\n", err)
		w.WriteHeader(500)
		return
	}
	_, err = w.Write(res)
	if err != nil {
		log.Printf("error on Writing response\n%s\n", err)
		w.WriteHeader(500)
		return
	}
}

func HandleRead(w http.ResponseWriter, r *http.Request) {
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		log.Printf("error on parsing read query\n%s\n", err)
		w.WriteHeader(500)
		return
	}

	filename := query.Get("file")
	content, err := os.ReadFile("posts/" + filename)
	if err != nil {
		log.Printf("error on reading file %s\n%s\n", filename, err)
		w.WriteHeader(500)
		return
	}
	_, err = w.Write(content)
	if err != nil {
		log.Printf("error on Writing response\n%s\n", err)
		w.WriteHeader(500)
		return
	}
}

func HandleWrite(w http.ResponseWriter, r *http.Request) {
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		log.Printf("error on parsing read query\n%s\n", err)
		w.WriteHeader(500)
		return
	}
	filename := query.Get("file")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error on reading request body\n%s\n", err)
		w.WriteHeader(500)
		return
	}

	err = os.WriteFile("posts/"+filename, body, os.ModePerm)
	if err != nil {
		log.Printf("error on writing new file %s\n%s\n", filename, err)
		w.WriteHeader(500)
		return
	}

	go Publish()
}

func HandleDelete(w http.ResponseWriter, r *http.Request) {
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		log.Printf("error on parsing read query\n%s\n", err)
		w.WriteHeader(500)
		return
	}
	filename := query.Get("file")

	err = os.Remove("posts/" + filename)
	if err != nil {
		log.Printf("error on removing file %s\n%s\n", filename, err)
		w.WriteHeader(500)
		return
	}
}

func HandlePublish(w http.ResponseWriter, r *http.Request) {
	err := Publish()
	if err != nil {
		w.WriteHeader(500)
		return
	}
}

func Publish() error {
	err := generate.GenereatePosts()
	if err != nil {
		log.Printf("error on generating posts\n%s\n", err)
		return err
	}
	err = upload.UploadDirectory("deps.me", "out", "ap-northeast-2")
	if err != nil {
		log.Printf("error on uploading directory\n%s\n", err)
		return err
	}
	return nil
}
