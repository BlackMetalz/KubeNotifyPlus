package utils

import (
	"fmt"
	"log"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

// checkContainerStatuses checks the status of the containers in the pod (Copilot reduce the complexity of the code)
func checkContainerStatuses(clientset *kubernetes.Clientset, namespace, podName string, containerStatuses []v1.ContainerStatus, message string) string {
	// Assume we have more than 1 container in the pod or even 1 container in the pod, it stills fine xD
	for _, containerStatus := range containerStatuses {
		if containerStatus.State.Waiting != nil {
			message += fmt.Sprintf("\nWaiting state with reason: *%s*", containerStatus.State.Waiting.Reason)
			message = checkLastTerminationState(clientset, namespace, podName, containerStatus, message)
		}
	}
	return message
}

// checkLastTerminationState checks the last termination state of the container (Copilot reduce the complexity of the code)
func checkLastTerminationState(clientset *kubernetes.Clientset, namespace, podName string, containerStatus v1.ContainerStatus, message string) string {
	// Check if the container has terminated before
	if containerStatus.LastTerminationState.Terminated != nil {
		terminatedState := containerStatus.LastTerminationState.Terminated
		message += fmt.Sprintf("\nLast terminated with reason: *%s*. Exit code: *%d*.", terminatedState.Reason, terminatedState.ExitCode)
		if terminatedState.Reason == "Error" || terminatedState.ExitCode != 0 {
			
			logs, err := GetPodLogs(clientset, namespace, podName, containerStatus.Name)
			if err != nil {
				log.Printf("Failed to get pod logs: %v", err)
			} else if logs != "" {
				message = appendLogsToMessage(logs, message)
			} else {
				message += "\n*Container logs*: No logs available"
			}
		}

		// Get the events of the pod
		events, err := GetPodEvents(clientset, namespace, podName)
		if err != nil {
			log.Printf("Failed to get pod events: %v", err)
		} else {
			message += fmt.Sprintf("\n*Pod events*:\n```%s```", events)
		}
	}

	return message
}

// appendLogsToMessage appends the logs to the message if the logs are too long (Copilot reduce the complexity of the code)
func appendLogsToMessage(logs, message string) string {
	// Handle the logs if it's too long
	if len(logs) > 2000 {
		logs = logs[:2000] + "...\n[log truncated]"
	} 
	// Append the logs to the message
	message += fmt.Sprintf("\n*Container logs*:\n```\n%s\n```", logs)
	
	return message
}
