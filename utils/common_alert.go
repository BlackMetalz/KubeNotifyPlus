package utils

import (
	"net/http"
	"fmt"
)

// Structure for the alert. Without cluster name for k8s
type CommonAlert struct {
	Status string `json:"status"`
	Labels struct {
		AlertName string `json:"alertname"`
	} `json:"labels"`
	Annotations struct {
		Message string `json:"message"`
	} `json:"annotations"`
}

// Structure for the payload
type CommonPayload struct {
	Data []CommonAlert `json:"alerts"`
}


// Process the alert
func CommonAlertHandler(w http.ResponseWriter, r *http.Request) {
	var payload DefaultPayload

    body, err := validateAndReadBody(w, r)
    if err != nil {
        return
    }

    err = debugAndUnmarshal(body, &payload, "PodAlertHandler")
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	for _, alert := range payload.Data {
		processCommonAlert(w, alert)
	}
}

func processCommonAlert(w http.ResponseWriter, alert DefaultK8SAlert) {
	// Struct the message
	message := fmt.Sprintf(
		"🚨 *%s: Alerts* 🚨\n"+
		"Alert Status: %s\n"+
		"Alert Name: %s\n"+
		"Message: %s\n",
		AppConfig.AppName,
		alert.Status,
		alert.Labels.AlertName,
		alert.Annotations.Message,
	)

	// Send the message to the user
	sendMessageToUser(w, message, "")
}