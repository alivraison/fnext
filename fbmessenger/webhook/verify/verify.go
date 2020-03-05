package verify

import (
	"fmt"
	"html"
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
	return "github.com/alivraison/fnext/fbmessenger/webhook/verify"
}

func (e *verifyExt) Setup(s fnext.ExtServer) error {
	s.AddEndpoint("GET", "/webhook", &simpleEndpoint{})
	return nil
}

// SimpleEndpoint is used for logging in. Returns a JWT token if successful.
type simpleEndpoint struct {
}

func (v *simpleEndpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SIMPLEENDPOINT SERVEHTTP")
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

//next.ServeHTTP(w, r.WithContext(ctx))

//Handle(next http.Handler) http.Handler
