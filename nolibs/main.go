package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"sync"
)

var (
	listUsersRe  = regexp.MustCompile(`^\/users[\/]*$`)
	getUserRe    = regexp.MustCompile(`^\/users\/(\d+)$`)
	CreateUserRe = regexp.MustCompile(`^\/users[\/]*$`)
)

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type datastore struct {
	m map[string]user
	*sync.RWMutex
}

type userHandler struct {
	store *datastore
}

func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet && listUsersRe.MatchString(r.URL.Path):
		h.List(w, r)
		return
	case r.Method == http.MethodGet && getUserRe.MatchString(r.URL.Path):
		h.Get(w, r)
		return
	case r.Method == http.MethodPost && CreateUserRe.MatchString(r.URL.Path):
		h.Create(w, r)
		return
	default:
		notFound(w, r)
		return
	}
}

func (h *userHandler) List(w http.ResponseWriter, r *http.Request) {
	users := make([]user, 0, len(h.store.m))
	h.store.RLock()
	for _, u := range h.store.m {
		users = append(users, u)
	}
	h.store.RUnlock()
	jsonBytes, err := json.Marshal(users)
	if err != nil {
		// log.Fatal("Error")
		internalServerError(w, r)
	}
	fmt.Println(string(jsonBytes))
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *userHandler) Get(w http.ResponseWriter, r *http.Request) {
	// điều đầu tiên cần phải làm là lấy id ra khỏi url
	matchs := getUserRe.FindStringSubmatch(r.URL.Path)
	if len(matchs) < 2 {
		notFound(w, r)
		return
	}
	h.store.RLock()
	user, ok := h.store.m[matchs[1]]
	if !ok {
		notFound(w, r)
		return
	}
	jsonByte, err := json.Marshal(user)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonByte)
}

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	u := &user{}
	// var u *user
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		badRequest(w, r)
		return
	}
	fmt.Println(*u)
	h.store.Lock()
	h.store.m[u.ID] = *u
	h.store.Unlock()
	jsonBytes, err := json.Marshal(u)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonBytes)
}

func badRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(`{"error": "Bad request!"}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "Not Found!"}`))
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(`{"error": "Internal Server Error!!"}`))
}

func main() {

	//test read file pem
	// Load PEM
	// pemfile, err := os.Open("./private.pem")

	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// // need to convert pemfile to []byte for decoding

	// pemfileinfo, _ := pemfile.Stat()
	// var size int64 = pemfileinfo.Size()
	// pembytes := make([]byte, size)

	// // read pemfile content into pembytes
	// buffer := bufio.NewReader(pemfile)
	// _, err = buffer.Read(pembytes)

	// // proper decoding now
	// data, _ := pem.Decode([]byte(pembytes))

	// pemfile.Close()

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	//end -test
	mux := http.NewServeMux()
	userH := &userHandler{
		store: &datastore{
			m: map[string]user{
				"1": {ID: "1", Name: "Bob"},
				"2": {ID: "2", Name: "KhaBD"},
			}, RWMutex: &sync.RWMutex{},
		},
	}
	mux.Handle("/users/", userH)
	// mux.Handle("/users", userH)
	http.ListenAndServe("localhost:8080", mux)
}
