package commons

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/josergdev/go-comments/models"
)

// DisplayMessage returns a message to the client
func DisplayMessage(w http.ResponseWriter, m models.Message) {
	j, err := json.Marshal(m)
	if err != nil {
		log.Fatal("Error parsing message to json", err)
	}

	w.WriteHeader(m.Code)
	w.Write(j)
}
