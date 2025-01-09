#!/bin/bash

# Configuration
HOST="localhost"
PORT="8080"
LOOP_COUNT=9999

# Endpoints to test
ENDPOINTS=("/499" "/500" "/502" "/503" "/504")

# Display usage information
usage() {
    echo "Usage: $0 [-h host] [-p port] [-c loop_count]"
    echo "  -h host        The host to send requests to (default: localhost)"
    echo "  -p port        The port to send requests to (default: 8080)"
    echo "  -c loop_count  The number of requests to send (default: 9999)"
    exit 1
}

# Parse command-line arguments
while getopts "h:p:c:" opt; do
    case ${opt} in
        h)
            HOST=${OPTARG}
            ;;
        p)
            PORT=${OPTARG}
            ;;
        c)
            LOOP_COUNT=${OPTARG}
            ;;
        *)
            usage
            ;;
    esac
done

# Function to send requests
send_requests() {
    for ((i=1; i<=LOOP_COUNT; i++)); do
        for endpoint in "${ENDPOINTS[@]}"; do
            URL="http://${HOST}:${PORT}${endpoint}"
            echo "Sending request to ${URL}"
            curl -s -o /dev/null -w "HTTP Status: %{http_code}\n" ${URL}
        done
    done
}

# Send the requests
send_requests