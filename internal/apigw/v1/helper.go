package v1

import (
	"appGo/pkg/api/apiv1"
	"appGo/pkg/httputil"
	"encoding/json"
	"net/http"

	"google.golang.org/grpc/status"
)

func handleGRPCError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	st := status.Convert(err)
	code := st.Code()
	w.WriteHeader(httputil.ConvertGRPCCodeToHTTP(code))
	if err := json.NewEncoder(w).Encode(
		apiv1.Error{
			Code:    httputil.ConvertGRPCToErrorCode(code),
			Message: nil,
		},
	); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
