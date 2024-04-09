package render

import (
	"fmt"
	"net/http"
)

func HandleError(err error, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Error executing template:", err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal server error"))
}
