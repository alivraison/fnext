package verify

import (
	"fmt"
	"net/http"

	"github.com/fnproject/fn/api/server"
	"github.com/fnproject/fn/fnext"
)

func init() {
	server.RegisterExtension(&fbwebhookExt{})
}

type fbwebhookExt struct {
}

func (e *fbwebhookExt) Name() string {
	return "github.com/alivraison/fnext/fbmessenger/webhook/verify"
}

func (e *fbwebhookExt) Setup(s fnext.ExtServer) error {
	s.AddEndpoint("GET", "/webhook", &verifyEndpoint{})
	return nil
}

type verifyEndpoint struct {
}

func (v *verifyEndpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("hub.mode") == "subscribe" && r.FormValue("hub.verify_token") == "CHANGE_TO_ENV" {
		fmt.Fprintf(w, r.FormValue("hub.challenge"))
		fmt.Println("Verified")
		return
	}
	fmt.Println("Failed Validation")
	http.Error(w, "Failed validation. Make sure the validation tokens match.", http.StatusForbidden)
	return
}
