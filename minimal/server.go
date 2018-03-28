package minimal

import (
	"fmt"
	"log"
	"net/http"
)

// Start function starts the server
func Start(port uint16) {
	logMessage := fmt.Sprintf("Server started at port %d", port)
	portString := fmt.Sprintf(":%d", port)

	log.Println(logMessage)
	http.ListenAndServe(portString, nil)
}
