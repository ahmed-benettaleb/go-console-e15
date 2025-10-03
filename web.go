package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
)

//go:embed static/*
var staticFiles embed.FS

// runWeb starts a simple HTTP server that serves a static page and an API endpoint
func runWeb(apiKey string, port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := staticFiles.ReadFile("static/index.html")
		if err != nil {
			http.Error(w, "index not found", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(data)
	})

	http.HandleFunc("/api/weather", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")
		if q == "" {
			http.Error(w, "missing q parameter", http.StatusBadRequest)
			return
		}
		// sanitize city
		city := strings.TrimSpace(q)
		res, err := getWeather(city, apiKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		resp := map[string]string{"result": res}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	if port == "" {
		port = "8080"
	}
	// special-case port "0" => ask OS for any free port
	if port == "0" {
		ln, err := net.Listen("tcp", ":0")
		if err != nil {
			log.Fatalf("unable to bind to a free port: %v", err)
		}
		actual := ln.Addr().(*net.TCPAddr).Port
		fmt.Printf("Starting web server at http://localhost:%d\n", actual)
		log.Fatal(http.Serve(ln, nil))
	}

	p, err := strconv.Atoi(port)
	if err != nil || p <= 0 {
		p = 8080
	}

	// Try requested port, then next 10 ports if occupied
	var ln net.Listener
	var try int
	for i := 0; i <= 10; i++ {
		try = p + i
		addr := ":" + strconv.Itoa(try)
		l, err := net.Listen("tcp", addr)
		if err != nil {
			// continue to next port
			continue
		}
		ln = l
		fmt.Printf("Starting web server at http://localhost:%d\n", try)
		break
	}
	if ln == nil {
		log.Fatalf("unable to bind to ports %d-%d", p, p+10)
	}
	log.Fatal(http.Serve(ln, nil))
}
