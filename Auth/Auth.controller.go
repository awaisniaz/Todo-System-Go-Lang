package Auth

import (
	"fmt"
	"net/http"

	"github.com/task-manager/Utilities"
)

func Login(w http.ResponseWriter, r *http.Request) {
	token := Utilities.GenerateToken("awaisniaz")
	fmt.Fprintf(w, token)
}
func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I am Register Route")
}
