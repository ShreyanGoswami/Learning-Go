package main

import "net/http"

func main() {
	// Note: defers are executed before panics and then the function returns
	// This would allow you to close resources even if panics occur
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error()) // Its a panic for developers because application didn't start
	}
}
