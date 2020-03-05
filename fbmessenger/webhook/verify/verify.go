package verify

import (
	"fmt"
	"net/http"

	"github.com/fnproject/fn/api/server"
	"github.com/fnproject/fn/fnext"
)

func init() {
	server.RegisterExtension(&verifyExt{})
}

type verifyExt struct {
}

func (e *verifyExt) Name() string {
	return "alivraison.com/chatbot/fbwebhook/verify"
}

func (e *verifyExt) Setup(s fnext.ExtServer) error {
	s.AddEndpoint("GET", "/webhook", &SimpleEndpoint{})
	return nil
}

// SimpleEndpoint is used for logging in. Returns a JWT token if successful.
type SimpleEndpoint struct {
}

func (v *SimpleEndpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SIMPLEENDPOINT SERVEHTTP")

}

//next.ServeHTTP(w, r.WithContext(ctx))

//Handle(next http.Handler) http.Handler
