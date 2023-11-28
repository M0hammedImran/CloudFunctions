package EnterprizeRedirect

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/m0hammedimran/CloudFunctions/EnterpriseRedirection/types"
	"github.com/m0hammedimran/CloudFunctions/EnterpriseRedirection/utils"
)

func init() {
	functions.HTTP("EnterprizeRedirect", EnterprizeRedirect)
}

// EnterprizeRedirect is an HTTP Cloud Function with a request parameter.
func EnterprizeRedirect(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Header)
	auth := r.Header.Get("Authorization")
	if auth == "" {
		log.Println("No Authorization header found")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorized")
		return
	}
	auth = strings.Replace(auth, "Bearer ", "", 1)
	log.Println(auth)

	if auth == "" {
		log.Println("No Authorization header found")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorized")
		return
	}

	jwt, err := utils.DecodeAccessToken(auth, types.StandardClaims{})
	if err != nil {
		log.Println("Could not decode the JWT")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorized")
		return
	}

	log.Println(jwt)
	prettyJWT, err := json.MarshalIndent(jwt, "", " ")
	if err != nil {
		log.Println("Could not Marshal this type")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorized")
		return
	}

	addNewLine := append(prettyJWT, '\n')
	fmt.Fprintf(w, "Hello, %s!", string(addNewLine))
}
