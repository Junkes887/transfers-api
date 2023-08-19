package middleware

import (
	"net/http"
	"strings"

	"github.com/Junkes887/transfers-api/pkg/httperr"
	"github.com/Junkes887/transfers-api/pkg/jwtToken"
)

func ValidateRoute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isTransfers := strings.Contains(r.URL.Path, "transfers")

		if isTransfers {
			token := r.Header.Get("Authorization")
			err := jwtToken.ValidateToken(token)

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			cpf, err := jwtToken.GetDocument(token)
			if err != nil {
				httperr.ErrorHttpStatusInternalServerError(err, w)
				return
			}

			r.Header.Add("cpf", cpf)
		}

		next.ServeHTTP(w, r)
	})
}
