package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {
	// 1. Load config variables
	viper.AddConfigPath("config")
	viper.SetConfigName("test")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	port := viper.GetString("server.port")

	// 2. Declare API methods and handlers
	http.HandleFunc("/helloworld", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello, World!2345")
	})

	fmt.Printf("Server running (port=%s), route: http://localhost:8080/helloworld\n", port)

	// 3. Start server
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

}
