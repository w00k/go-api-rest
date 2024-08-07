package service

import (
	"encoding/json"
	"go-api-rest/src/model"
	"io"

	"log"
	"net/http"
)

const timezomeURL = "https://worldtimeapi.org/api/timezone"

func TimezoneService() (data []string, exception model.Exception) {

	// cal the backend
	response, error := http.Get(timezomeURL)

	// cheking response errors
	if error != nil {
		log.Printf("timezoneService: %s - %s", timezomeURL, error.Error())
		exception.Status = response.StatusCode
		exception.Code = http.StatusText(response.StatusCode)
		exception.Message = response.Status
		return
	}

	// get the body
	body, error := io.ReadAll(response.Body)

	// check the errors to read the body
	if error != nil {
		log.Printf("timezoneService: can't read the body - %s", error.Error())
		exception.Status = 402
		exception.Code = "Unprocessable Entity"
		exception.Message = "Error with API response"
		return
	}

	defer response.Body.Close()

	json.Unmarshal(body, &data)
	return
}
