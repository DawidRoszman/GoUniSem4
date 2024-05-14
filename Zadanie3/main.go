package main

import (
	"net/http"
)

type Record struct {
	date     string
	country  string
	activity string
	name     string
	injury   string
	species  string
}

func main() {
	http.HandleFunc("/attacks/get", handleGetAttacks())
	http.HandleFunc("/attacks/post", handlePostAttack())
	http.HandleFunc("/attacks/update", handleUpdateAttack())
	http.HandleFunc("/attacks/delete", handleDeleteAttack())
}
