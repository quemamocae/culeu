package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Links</title>
		</head>
		<body>
			<h1>Welcome to My Website</h1>
			<p>Here are some useful links:</p>
			<ul>
				<li><a href="https://www.example.com">example</a></li>
				<li><a href="https://www.example.com">example</a></li>
				<li><a href="https://www.example.com">example</a></li>
			</ul>
		</body>
		</html>`
	fmt.Fprint(w, html)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
