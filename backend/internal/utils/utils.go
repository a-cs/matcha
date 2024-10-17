package utils

import (
	"backend/internal/adpter/dto"
	"encoding/json"
	"net/http"
)

func ServiceResponse(responseWriter http.ResponseWriter, resp *dto.ResponseDTO, status int) {
	responseWriter.WriteHeader(status)
	if resp != nil {
		json.NewEncoder(responseWriter).Encode(resp)
	}
}
