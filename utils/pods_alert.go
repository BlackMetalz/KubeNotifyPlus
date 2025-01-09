package utils

import (
	"fmt"
	"log"
	"net/http"
)

// Alert represents the structure of an alert. For Pod alerts.
type PodAlert struct {
	Status string `json:"status"`
	Labels struct {
		Alertname string `json:"alertname"`
		Cluster   string `json:"cluster"`
		Namespace string `json:"namespace"`
		Pod       string `json:"pod"`
	} `json:"labels"`
	Annotations struct {
		Message string `json:"message"`
	} `json:"annotations"`
}

// PodPayload represents the structure of the webhook payload
type PodPayload struct {
	Data []PodAlert `json:"alerts"`
}

// HelloWorld is a simple handler that returns "Hello World". Required.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

// PodAlertHandler is a handler that processes alerts of pod from Prometheus
func PodAlertHandler(w http.ResponseWriter, r *http.Request) {
	var payload PodPayload
	
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
		processPodAlert(w, alert)
	}
}

func processPodAlert(w http.ResponseWriter, alert PodAlert) {
	message := fmt.Sprintf(
		"🚨 *%s: Pod Alert* 🚨\n*Alert Status:* %s\n*Alert:* %s\n*Cluster:* %s\n*Namespace:* %s\n*Pod:* %s\n*Message:* %s",
		AppConfig.AppName,
		alert.Status,
		alert.Labels.Alertname,
		alert.Labels.Cluster,
		alert.Labels.Namespace,
		alert.Labels.Pod,
		alert.Annotations.Message,
	)

	err := sendTelegramMessage(message, "")
	if err != nil {
		log.Printf("Failed to send alert message to Telegram: %v", err)
	} else {
		SendResponse(w, "Message sent to Telegram")
	}

	podData := alert.Labels.Pod
	podNamespace := alert.Labels.Namespace

	serviceName := getServiceName(podData)
	userID, err := GetServiceMappings(serviceName)
	if err != nil {
		log.Printf("Failed to get service mappings: %v", err)
	} else {
		sendMessageToUser(w, message, userID)
	}

	// Investigate the pod failure if only alert status is firing
	if AppConfig.PodFailureInvestigation && alert.Status == "firing" {
		InvestigatePodFailure(w, podData, podNamespace, userID)
	}
}

func InvestigatePodFailure(w http.ResponseWriter, podData, podNamespace, userID string) {
	podReason := PodReason(podData, podNamespace)
	log.Println(podReason)
	err := sendTelegramMessage(podReason, "")
	if err != nil {
		log.Printf("Failed to send podReason message to Telegram: %v", err)
	} else {
		SendResponse(w, "podReason sent to Telegram")
	}

	if userID != "" {
		err = sendTelegramMessage(podReason, userID)
		if err != nil {
			log.Printf("Failed to send podReason message to Telegram with userID: %v", err)
		} else {
			SendResponse(w, fmt.Sprintf("podReason sent to Telegram with userID %s", userID))
		}
	}
}
