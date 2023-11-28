package EnterprizeRedirect

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/m0hammedimran/CloudFunctions/EnterpriseRedirection/types"
	"github.com/m0hammedimran/CloudFunctions/EnterpriseRedirection/utils"
)

func init() {
	functions.HTTP("EnterprizeRedirect", EnterprizeRedirect)
}

// EnterprizeRedirect is an HTTP Cloud Function with a request parameter.
func EnterprizeRedirect(w http.ResponseWriter, r *http.Request) {
	auth, err := utils.GetTokenFromHeader(r.Header)
	if err != nil {
		log.Println("Could not get the token from the header")
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

	fmt.Fprintf(w, "Valid Enterprize: %d", jwt.Claims.Id)
}
