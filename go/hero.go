package main

import (
        "encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// PROBLEM DESCRIPTION:
// Your goal here is to design an API that allows for hero tracking, much like the Vue problem
// You are to implement an API (for which the skeleton already exists) that has the following capabilities
// - Get      : return a JSON representation of the hero with the name supplied
// - Make     : create a superhero according to the JSON body supplied
// - Calamity : a calamity of the supplied level requires heroes with an equivalent combined powerlevel to address it.
//              Takes a calamity with powerlevel and at least 1 hero. On success return a 200 with json response indicating the calamity has been resolved.
//              Otherwise return a response indicating that the heroes were not up to the task. Addressing a calamity adds 1 point of exhaustion.
// - Rest     : recover 1 point of exhaustion
// - Retire   : retire a superhero, someone may take up his name for the future passing on the title
// - Kill     : a superhero has passed away, his name may not be taken up again.

// On success all endpoints should return a status code 200.

// If a hero reaches an exhaustion level of maxExhaustion then they die.

// You are free to decide what your API endpoints should be called and what shape they should take. You can modify any code in this file however you'd like.

// NOTE: you may want to install postman or another request generating software of your choosing to make testing easier. (api is running on localhost port 8081)

// NOTE the second: the API is receiving asynchronous requests to manage our super friends. As such, your hero access should be thread-safe.
// Even if the operations are extremely lightweight we want to make our application scalable.
// Your multithreaded protection should allow the API to still be performant even if these requests took a reasonably long time. This means that
// a global lock/mutex that makes the API only handle 1 request at a time is not an acceptable solution to this problem

// NOTE the third: There are many ways to make whatever package-level tracking you implement thread-safe, feel free to change anything about this file (without changing the functionality of the program) to do so.
// i.e. add package-level maps, add functions that take the hero struct as reference, modify the hero struct, creating access control paradigms
// I highly recommend looking into channels, mutexes, and other golang memory management patterns and pick whatever you're most comfortable with.
// For mad props: a timeout on the memory management process which returns a resource not available.

// Bonus: If you're having fun (this is by no means necessary) you can make the calamity hold the heroes up for a time and delay their unlocking in a go-routine
// example:
// go func(h *hero) {
//     time.Sleep(calamityTime)
//     // release lock on hero
// }(heroPtr)

var maxExhaustion = 10

type hero struct {
	PowerLevel int    `json:"PowerLevel"`
	Exhaustion int    `json:"Exhaustion"`
	Name       string `json:"Name"`
	// TODO: changeme?
}

type calamity struct {
       Level int
       Heros [] string
}

type health struct {
     Status string
}

var healthData health

// TODO: add storage and memory management

// TODO: add more routines

func healthGet(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "HELLO:  ");
	fmt.Fprintf(w, healthData.Status);
}

func healthPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var t health
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	healthData = t
        w.WriteHeader(http.StatusOK)
	return
}

var heroData hero

var heroMap = make(map[string]hero)

func heroGet(w http.ResponseWriter, r *http.Request) {
        var name = mux.Vars(r)["name"]
	hero, present := heroMap[name]
	if(present) {
	       json.NewEncoder(w).Encode(hero)
	} else {
	       json.NewEncoder(w).Encode(present)
	}
	return
}

func heroMake(w http.ResponseWriter, r *http.Request) {
        decoder := json.NewDecoder(r.Body)
	var t hero
	err := decoder.Decode(&t)
	if(err != nil) {
	       panic(err)
	}
	heroData = t
	heroMap[t.Name] = t
	w.WriteHeader(http.StatusOK)
	return
}

func heroCalamity(w http.ResponseWriter, r *http.Request) {
     decoder := json.NewDecoder(r.Body)
     var t calamity
     err := decoder.Decode(&t)
     if(err != nil) {
     	    panic(err)
     }
     var sum int
     for i:=0 ; i < len(t.Heros); i++ {
     	 sum += heroMap[t.Heros[i]].PowerLevel
     }
     delta := sum < t.Level
     if(delta) {
     	    for i:=0 ; i < len(t.Heros); i ++ {
	    	k := t.Heros[i];
	    	h := heroMap[k]
		h.Exhaustion = h.Exhaustion + 1
		heroMap[k] = h
	    }
     }
     json.NewEncoder(w).Encode(delta)
     return
}

func heroRest(w http.ResponseWriter , r *http.Request) {
     var name = mux.Vars(r)["name"]
     h := heroMap[name]
     h.Exhaustion = h.Exhaustion - 1
     heroMap[name] = h
     w.WriteHeader(http.StatusOK)
     return
}

func linkRoutes(r *mux.Router) {
        r.HandleFunc("/health/check", healthGet).Methods("GET")
	r.HandleFunc("/health/check", healthPost).Methods("POST")
	r.HandleFunc("/hero", heroMake).Methods("POST")
	r.HandleFunc("/hero/calamity", heroCalamity).Methods("POST")
	r.HandleFunc("/hero/{name}", heroGet).Methods("GET")
	r.HandleFunc("/hero/rest/{name}", heroRest).Methods("POST")
	// TODO: add more routes
}

func main() {
	// create a router
	router := mux.NewRouter()
	// and a server to listen on port 8081
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8081),
		Handler: router,
	}
	// link the supplied routes
	linkRoutes(router)
	// wait for requests
	log.Fatal(server.ListenAndServe())
}
