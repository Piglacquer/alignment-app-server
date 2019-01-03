package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type data struct {
	stuff     string
	stuffer   string
	something int
}

func ShopsHandler(w http.ResponseWriter, r *http.Request) {
	p := data{stuff: "howdy", stuffer: "hello", something: 3}
	fmt.Printf("%+v", p)
	p1, _ := json.Marshal(p)
	json.NewEncoder(w).Encode(p1)
}
