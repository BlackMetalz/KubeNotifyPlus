package main

import (
	// "fmt"
	"log"
	"net/http"
	"os"
	// Utils
	"KubeNotifyPlus/utils"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(os.Stdout) // Ensure logs go to standard output
		
    // Load the configuration from the config.json file
    err := utils.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load configuration: %v", err)
    }

	// Initialize the MySQL connection
	err = utils.InitMySQL()
	if err != nil {
		log.Fatalf("Failed to initialize MySQL connection: %v", err)
	}
	// // Test
	// apiServer := "https://hmmm:6443"
    // // Setup the Kubernetes client
    // clientset, err := utils.SetupK8sClient(apiServer)
    // if err != nil {
    //     log.Fatalf("Failed to set up Kubernetes client: %v", err)
	// }
	// logs, err := utils.GetPodLogs(clientset, "oom", "oom-example-56d77685f5-dqv7z", "oom-container")
	// if err != nil {
	// 	log.Fatalf("Failed to get pod logs: %v", err)
	// }
	// fmt.Println(logs)
	// log.Fatalf("Failed to create clientset: %v", err)

	
	http.HandleFunc("/", utils.HelloWorld)
	http.HandleFunc("/k8s-app", utils.PodAlertHandler) // KubePod*
	http.HandleFunc("/k8s-ingress", utils.IngressAlertHandler) // ingress for k8s
	http.HandleFunc("/k8s-default", utils.DefaultAlertHandler) // Rest for k8s
	http.HandleFunc("/common-alert", utils.CommonAlertHandler) // Common alert
	// Log that the server has started
	log.Println("Server started on port 6789")
	// Start the server on port 6789 with the default ServeMux
	log.Fatal(http.ListenAndServe(":6789", nil))
}