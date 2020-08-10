package main

import (
	"fmt"
	"net/http"
	"github.com/goburrow/modbus"
)

func main() {

	client := modbus.TCPClient("localhost:502")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		results, err := client.WriteMultipleRegisters(1, 2, []byte{0, 3, 0, 4})
		if err != nil {
			http.Error(w, "There was some error", http.StatusInternalServerError)
			return
		}
		w.Write(results)
	})

	fmt.Println("Starting modbus server...")
	http.ListenAndServe(":8080", nil)
}
