package main 

import (
	"fmt"
	"net/http"
	"go-jwt/handler"
	"go-jwt/driver"
)

func main() {
	driver.ConnectMongoDB()
	
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/register", handler.Register)
	http.HandleFunc("/user", handler.GetUser)

	fmt.Println("Server running [:8000]")
	http.ListenAndServe(":8000", nil)
}