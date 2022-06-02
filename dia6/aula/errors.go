package main

import (
	"fmt"
	"net/http"
	"os"
)

type MyError struct {
	status int
	msg    string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%v - %v", e.status, e.msg)
}

func HandleStatusCode(status int) (int, error) {
	if status >= http.StatusBadRequest {
		return http.StatusInternalServerError, &MyError{status: status, msg: "Deu Errado!!!"}
	}
	return status, nil
}

func main() {
	status, err := HandleStatusCode(505)
	if err != nil {
		fmt.Println(status, err)
		os.Exit(1)
	}

	fmt.Println("Tudo certo, status:", status)
}
