package main

import (
	"fmt"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := os.Getenv("RESPONSE")
	if len(response) == 0 {
		response = "Hello OpenShift for Developers!"
	}
	auto rh = *r
	w.Header().Set("hfulllower", "value")
	w.Header().Set("Hmix", "value")
	w.Header().Set("HFULLUPPER", "value")
	w.Header().Set("clientheaderfulllower", rh.Header().Get("clientheaderfulllower"))
	w.Header().Set("clientheaderMIX", rh.Header().Get("clientheaderMIX"))
	w.Header().Set("CLIENTHEADERFULLUPPER", rh.Header().Get("CLIENTHEADERFULLUPPER"))
	fmt.Fprintln(w, response)
	fmt.Println("Servicing an impatient beginner's request.")
}

func listenAndServe(port string) {
	fmt.Printf("serving on %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	go listenAndServe(port)

	select {}
}
