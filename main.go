package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Server struct {
	Address string `json:"Address"`
	SslGrade string `json:"SslGrade"`
	Country string `json:"Country"`
	Owner string `json:"Owner"`
}

type Servers []Server

type Truora struct {
	Servers Servers 
	ServersChanged string `json:"ServersChanged"`
	SslGrade string `json:"SslGrade"`
	PreviousSslGrade string `json:"PreviousSslGrade"`
	Logo string `json:"Logo"`	
	Title string `json:"Title"`
	IsDown string `json:"IsDown"`
}

type Search struct {
	Url string 
}

var Searchs []Search

func allServers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	servers := Servers {
		Server { Address: "server1", SslGrade: "B", Country: "US", Owner: "Amazon.com, Inc."  },
		Server { Address: "server2", SslGrade: "A+", Country: "US", Owner: "Amazon.com, Inc."  },
		Server { Address: "server3", SslGrade: "A", Country: "US", Owner: "Amazon.com, Inc."  },
	} 

	truora := Truora {
		ServersChanged:    "true", 
		SslGrade:    "B", 
		PreviousSslGrade:    "A", 
		Logo:    "https://server.com/icon.png",
		Title:    "Title of the page",
		IsDown:    "false",
		Servers: servers,
	}

	search := Search { Url: r.URL.Path}
	Searchs = append(Searchs, search)
	fmt.Println("truora.com endpoint")
	json.NewEncoder(w).Encode(truora)

}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Fprint(w, "Homepage Endpoint was triggered " , r.URL.Path)
	search := Search { Url: r.URL.Path}
	if(r.URL.Path != "/favicon.ico" && r.URL.Path != "/") {
		Searchs = append(Searchs, search)
	}
}

func allSearchs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(Searchs)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/truora.com", allServers)
	http.HandleFunc("/searchs", allSearchs)
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func main() {
	handleRequests()
}
