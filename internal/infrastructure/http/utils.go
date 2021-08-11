package http

import (
	"encoding/json"
	"github.com/cipma/smsbin/infrastructure/dto"
	"log"
	"net/http"
)

func writeJsonResponse(w http.ResponseWriter, data interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)

	_ = json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, err error, code int) {
	log.Println(err)

	writeJsonResponse(
		w,
		dto.NewError(err),
		code,
	)
}
