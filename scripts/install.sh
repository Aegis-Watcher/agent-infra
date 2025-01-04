#!/bin/bash

# Metrics Agent Installation Script

echo "Starting Metrics Agent installation..."

# Check for root permissions
if [[ $EUID -ne 0 ]]; then
  echo "This script must be run as root. Please run with sudo."
  exit 1
fi

AGENT_UUID=$1
AGENT_SECRET=$2

if [ -z "$AGENT_UUID" ]; then
  echo "Please provide the Agent UUID as the first argument."
  exit 1
fi

if [ -z "$AGENT_SECRET" ]; then
  echo "Please provide the Agent Secret as the second argument."
  exit 1
fi

# Define variables
BINARY_PATH="/usr/local/bin/aegis-watcher-metrics-agent"
ENV_FILE="/etc/aegis-watcher-metrics-agent.env"
SERVICE_FILE="/etc/systemd/system/aegis-watcher-metrics-agent.service"

# Function to generate a UUID (Instance ID)
generate_instance_id() {
  if command -v uuidgen >/dev/null 2>&1; then
    uuidgen
  else
    # Fallback if uuidgen is not available
    echo "$(date +%s)-$$" | sha256sum | base64 | head -c 32
  fi
}

# Check if Instance ID already exists to prevent overwriting
if [ -f "$ENV_FILE" ]; then
  # Extract existing Instance ID
  EXISTING_INSTANCE_ID=$(grep "^AEGIS_WATCHER_METRICS_AGENT_INSTANCE_ID=" "$ENV_FILE" | cut -d '=' -f2)
  if [ -n "$EXISTING_INSTANCE_ID" ]; then
    echo "Existing Instance ID found: $EXISTING_INSTANCE_ID"
    INSTANCE_ID="$EXISTING_INSTANCE_ID"
  else
    # Generate a new Instance ID if not present
    INSTANCE_ID=$(generate_instance_id)
    echo "AEGIS_WATCHER_METRICS_AGENT_INSTANCE_ID=$INSTANCE_ID" >>"$ENV_FILE"
    echo "Generated new Instance ID: $INSTANCE_ID"
  fi
else
  # Generate a new Instance ID
  INSTANCE_ID=$(generate_instance_id)
fi

# Download or build the binary
echo "Downloading Metrics Agent binary..."
LATEST_RELEASE_URL=$(curl -s https://api.github.com/repos/Aegis-Watcher/agent-infra/releases/latest | grep -oP '"browser_download_url": "\K(.*agent-infra)"' | tr -d '"')
if [ -z "$LATEST_RELEASE_URL" ]; then
  echo "Failed to fetch the latest release URL."
  exit 1
fi

curl -L -o "$BINARY_PATH" "$LATEST_RELEASE_URL"
if [ $? -ne 0 ]; then
  echo "Failed to download the Metrics Agent binary."
  exit 1
fi

chmod +x "$BINARY_PATH"

# Create environment file
echo "Creating environment configuration file..."
cat <<EOL >"$ENV_FILE"
AGENT_UUID=$AGENT_UUID
AGENT_SECRET=$AGENT_SECRET
AEGIS_WATCHER_METRICS_AGENT_INSTANCE_ID=$INSTANCE_ID
EOL

# Create systemd service file
echo "Setting up systemd service..."
cat <<EOL >"$SERVICE_FILE"
[Unit]
Description=Aegis Watcher Metrics Agent
After=network.target

[Service]
EnvironmentFile=$ENV_FILE
ExecStart=$BINARY_PATH
Restart=always
RestartSec=5
StartLimitInterval=0

[Install]
WantedBy=multi-user.target
EOL

# Reload systemd, enable, and start the service
echo "Reloading systemd daemon..."
systemctl daemon-reload

echo "Enabling Metrics Agent service to start on boot..."
systemctl enable aegis-watcher-metrics-agent

echo "Starting Metrics Agent service..."
systemctl start aegis-watcher-metrics-agent

# Confirm installation
echo "Metrics Agent installation completed successfully!"
echo "You can check the service status using: sudo systemctl status aegis-watcher-metrics-agent"
