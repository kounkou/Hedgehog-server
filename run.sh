#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# Function to print usage instructions
usage() {
  echo "Usage: $0 [-d] [-h]"
  echo "Options:"
  echo "  -d    Run containers in detached mode (background)"
  echo "  -h    Show this help message"
}

# Default options
DETACHED=false

# Parse command-line options
while getopts "dh" opt; do
  case $opt in
    d)
      DETACHED=true
      ;;
    h)
      usage
      exit 0
      ;;
    *)
      usage
      exit 1
      ;;
  esac
done

# Define cleanup function to handle errors
cleanup() {
  echo "Error occurred. Cleaning up..."
  docker-compose down
  sudo lsof -i :27017
  # sudo kill <PID>
  exit 1
}
trap cleanup ERR

# Start the script
echo "Shutting down any running containers..."
docker-compose down

echo "Rebuilding Docker containers without cache..."
docker-compose build --no-cache

if $DETACHED; then
  echo "Starting Docker containers in detached mode..."
  docker-compose up -d
  echo "Containers are running in the background. Use 'docker-compose logs' to view logs."
else
  echo "Starting Docker containers in attached mode..."
  docker-compose up
fi
