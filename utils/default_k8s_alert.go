package utils

import (
	"fmt"
	"net/http"
)

// DefaultAlert struct for the default alerts
type DefaultK8SAlert struct {
	Status string `json:"status"`
	Labels struct {
		AlertName string `json:"alertname"`
		Cluster   string `json:"cluster"`
	} `json:"labels"`
	Annotations struct {
		Message string `json:"message"`
	} `json:"annotations"`
}

// DefaultPayload struct for the default alerts payload
type DefaultPayload struct {
	Data []DefaultK8SAlert `json:"alerts"`
}

// Payload for the default alerts


// DefaultAlertHandler is a handler that processes default alerts
func DefaultAlertHandler(w http.ResponseWriter, r *http.Request) {
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
		processDefaultAlert(w, alert)
	}
}

func processDefaultAlert(w http.ResponseWriter, alert DefaultK8SAlert) {
	// Struct the message
	message := fmt.Sprintf(
		"🚨 *%s: K8S Alert* 🚨\n"+
			"*Alert Status*: %s\n"+
			"*Alert Name*: %s\n"+
			"*Cluster*: %s\n"+
			"*Message*: %s\n",
		AppConfig.AppName,
		alert.Status,
		alert.Labels.AlertName,
		alert.Labels.Cluster,
		alert.Annotations.Message,
	)

	// Send the message to the user
	sendMessageToUser(w, message, "")
}
