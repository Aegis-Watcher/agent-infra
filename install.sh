#!/bin/bash

# Metrics Agent Installation Script

echo "Starting Metrics Agent installation..."

# Check for root permissions
if [[ $EUID -ne 0 ]]; then
  echo "This script must be run as root. Please run with sudo."
  exit 1
fi

# Define variables
BINARY_PATH="/usr/local/bin/aegis-watcher-metrics-agent"
ENV_FILE="/etc/aegis-watcher-metrics-agent.env"
SERVICE_FILE="/etc/systemd/system/aegis-watcher-metrics-agent.service"

# Download or build the binary
echo "Downloading Metrics Agent binary..."
curl -L -o "$BINARY_PATH" https://github.com/Aegis-Watcher/agent-infra/releases/latest/aegis-watcher-metrics-agent
chmod +x "$BINARY_PATH"

# Prompt for environment variables
echo "Please enter the following configuration details:"
read -p "Agent UUID: " AGENT_UUID
read -p "Agent Secret: " AGENT_SECRET

# Create environment file
echo "Creating environment configuration file..."
cat <<EOL >$ENV_FILE
AGENT_UUID=$AGENT_UUID
AGENT_SECRET=$AGENT_SECRET
EOL

# Create systemd service file
echo "Setting up systemd service..."
cat <<EOL >$SERVICE_FILE
[Unit]
Description=Aegis Watcher Metrics Agent
After=network.target

[Service]
EnvironmentFile=$ENV_FILE
ExecStart=$BINARY_PATH
Restart=always

[Install]
WantedBy=multi-user.target
EOL

# Reload systemd, enable, and start the service
echo "Starting Metrics Agent service..."
systemctl daemon-reload
systemctl enable aegis-watcher-metrics-agent
systemctl start aegis-watcher-metrics-agent

# Confirm installation
echo "Metrics Agent installation completed successfully!"
echo "You can check the service status using: sudo systemctl status aegis-watcher-metrics-agent"
