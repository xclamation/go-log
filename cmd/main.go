package main

import (
	"errors"
	"fmt"
	"github/xclamation/go-log/log"
	"net/http"
	"os"
	"time"
)

type User struct {
	ID   int
	Name string
}

func ServeHTTP(w http.ResponseWriter, r *http.Request, logger log.Logger) {
	fmt.Println("ServeHTTP is being called")
	lg := logger.Begin()
	defer lg.End()

	// Log the receipt of an HTTP request
	lg.Log("Receive HTTP Request\n")
	lg.Trace("HTTP request: ", r, "\n")

	// Simulate processing the request
	lg.Log("Processing request...\n")
	time.Sleep(1 * time.Second)

	// Simulate inserting data into a database
	lg.Log("Tried to INSERT into database.\n")
	err := InsertIntoDatabase()
	if err != nil {
		lg.Error("Database INSERT failed: ", err, "\n")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Infrom about important events
	lg.Inform("New user created successfully \n")

	// Simulate serving a response
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Request processed successfully"))

	// Trace user information
	users := []User{
		{ID: 1, Name: "Ramil"},
		{ID: 2, Name: "Sasha"},
	}

	for _, user := range users {
		lg.Trace("User with ID ", user.ID, " is ", user.Name, "\n")
	}
}

// Simulate a function that inserts data into a database
func InsertIntoDatabase() error {
	// Simulating a database operation
	if time.Now().Unix()%2 == 0 { // Simulate an occasional failure
		return errors.New("simulated database error") // Should not be capitalized
	}
	return nil
}

func main() {
	logger := log.NewLogger(log.WithOutput(os.Stdout))

	// logger.Begin()
	// logger.Log("Test\n")

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	ServeHTTP(w, r, logger)
	// })

	// fmt.Println("Starting server on :8080")
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	logger.Error("Failed to start server: ", err, "\n")
	// }
	
	logger.Log("test")
}
