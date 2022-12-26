package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Err struct {
	Message string `json:"message"`
}

var users = []User{
	{ID: 1, Name: "kridsada", Age: 26},
	{ID: 2, Name: "kridsada", Age: 26},
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		var u User
		err = json.Unmarshal(body, &u)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		users = append(users, u)

		w.Write([]byte(`{"name": "kridsada", "method": "POST"}`))
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}
