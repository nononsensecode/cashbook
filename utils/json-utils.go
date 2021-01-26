package utils

import (
	"cashbook/cashbookerror"
	"encoding/json"
	"log"
)

// RespondWithError creates an APIError object and converts it to string
func RespondWithError(code int64, message string) string {
	apiError := cashbookerror.NewAPIError(code, message)
	jsonObject, err := json.Marshal(apiError)
	if err != nil {
		log.Println(err)
		return "Unknown error"
	}

	return string(jsonObject)
}

// RespondWithJSON converts any payload to JSON string
func RespondWithJSON(payload interface{}) string {
	jsonObject, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		return RespondWithError(500, "Cannot convert object to JSON string")
	}

	return string(jsonObject)
}