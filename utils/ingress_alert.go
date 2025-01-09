package utils

import (
	"fmt"
	"net/http"
	"log"
)

// struct for Ingress alerts
type IngressAlert struct {
	Status      string `json:"status"`
	Labels      struct {
		AlertName string `json:"alertname"`
		Cluster   string `json:"cluster"`
		Ingress   string `json:"ingress"`
		Status    string `json:"status"`
	} `json:"labels"`
	Annotations struct {
		Description string `json:"description"`
	} `json:"annotations"`
}

// struct for the webhook payload
type IngressPayload struct {
	Data []IngressAlert `json:"alerts"`
}

// IngressAlertHandler is a handler that processes Ingress alerts
func IngressAlertHandler(w http.ResponseWriter, r *http.Request) {
	var payload IngressPayload

    body, err := validateAndReadBody(w, r)
    if err != nil {
        return
    }

    err = debugAndUnmarshal(body, &payload, "PodAlertHandler")
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	// Process each alert in the payload
	for _, alert := range payload.Data {
		processIngressAlert(w, alert)
	}
}

// processIngressAlert processes an Ingress alert
func processIngressAlert(w http.ResponseWriter, alert IngressAlert) {
	// Struct the message
	message := fmt.Sprintf(
		"🚨 *%s: Ingress Alert* 🚨\n"+
			"*Alert Status*: %s\n"+
			"*Cluster*: %s\n"+
			"*Ingress Name*: %s\n"+
			"*Ingress Status*: %s\n"+
			"*Description*: %s",
		AppConfig.AppName, alert.Status, alert.Labels.Cluster, alert.Labels.Ingress, alert.Labels.Status, alert.Annotations.Description,	
	)

	// Send the alert to the Telegram chat default
	err := sendTelegramMessage(message, "")
	if err != nil {
		http.Error(w, "Failed to send Telegram message", http.StatusInternalServerError)
		return
	}

	// Define ingress to check for mapping
	ingress := alert.Labels.Ingress

	// Check if the ingress is in the mapping
	userID, err := GetServiceMappings(ingress)
	if err != nil {
		log.Printf("Failed to get service mappings: %v", err)
	} else {
		sendMessageToUser(w, message, userID)
	}

	// Send the response to the client
	SendResponse(w, "Alert received")
}