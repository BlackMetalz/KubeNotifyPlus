package utils

import (
    "fmt"
    "bytes"
	"context"
	"strings"
	"io"
    "log"

    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/rest"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    v1 "k8s.io/api/core/v1"
)

// SetupK8sClient sets up the Kubernetes client using a service account token
func SetupK8sClient(apiServer string) (*kubernetes.Clientset, error) {
    // Read the service account token from the environment variable
    token := AppConfig.K8sToken
    if token == "" {
        return nil, fmt.Errorf("K8sToken is not set")
    }

    // Create a REST config with TLSClientConfig.InsecureSkipVerify set to true
    config := &rest.Config{
        Host:        apiServer,
        BearerToken: token,
        TLSClientConfig: rest.TLSClientConfig{
            Insecure: true, // Skip certificate verification
        },
    }

    // Create a Kubernetes client
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        return nil, err
    }

    return clientset, nil
}

// GetPodLogs fetches the logs of the specified container in the pod
func GetPodLogs(clientset *kubernetes.Clientset, namespace, podName, containerName string) (string, error) {
	// Get the logs of the specified container in the pod
	podLogOpts := v1.PodLogOptions{
		Container: containerName,
		Previous:  true,
	}

	// Initialize the request
	req := clientset.CoreV1().Pods(namespace).GetLogs(podName, &podLogOpts)
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		return "NULL WTF", err
	}
	// Close the stream when the function returns
	defer podLogs.Close()

	// Read the logs into a buffer
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		return "", err
	}

	// Return the logs as a string
	return buf.String(), nil
}

// GetPodEvents fetches the events of the specified pod with type "Warning" only
func GetPodEvents(clientset *kubernetes.Clientset, namespace, podName string) (string, error) {
	eventList, err := clientset.CoreV1().Events(namespace).List(context.TODO(), metav1.ListOptions{
		FieldSelector: fmt.Sprintf("involvedObject.name=%s,involvedObject.kind=Pod,type=Warning", podName),
	})
	if err != nil {
		return "", err
	}

	var events []string
	// Avoid limit of telegram
	totalLength := 0
    for _, event := range eventList.Items {
		// Show only the type, reason, age and message of the event
        eventStr := fmt.Sprintf(
            "\nType: %s\nReason: %s\nAge: %s\nMessage: %s\n",
            event.Type, event.Reason, event.LastTimestamp.Time.String(), event.Message,
        )
		// Check if the total length of the events is greater than 3500 characters
        eventLength := len(eventStr)
        if totalLength+eventLength > 3500 {
            events = append(events, "\n... [events truncated]")
            break
        }
		// Append the event to the list
        events = append(events, eventStr)
        totalLength += eventLength
    }

	// If no events are found, return a message
	if len(events) == 0 {
		return "No warning events found", nil
	}

	return strings.Join(events, "\n"), nil
}

// PodReason returns the reason why the pod is not running. EG: CrashLoopBackOff, ImagePullBackOff, etc.
func PodReason(podName string, namespace string) string {
	// Initialize the Kubernetes client
	apiServer := AppConfig.ApiServer
	clientset, err := SetupK8sClient(apiServer)
	if err != nil {
		log.Fatalf("Failed to set up Kubernetes client: %v", err)
	}

	// Get the pod information
	pod, err := clientset.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		log.Printf("Failed to get pod information: %v", err)
	}

	// Query the pod status
	message := fmt.Sprintf(
		"💥 *%s: Pod Reason* 💥\nPod *%s* in namespace *%s* is not running!",
		AppConfig.AppName, 
		podName, 
		namespace,
	)
	message = checkContainerStatuses(clientset, namespace, podName, pod.Status.ContainerStatuses, message)

	return message
}