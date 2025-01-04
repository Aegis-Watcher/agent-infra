# Agent Infra

The Agent Infra is a lightweight tool designed to collect metrics for CPU, memory, and disk usage and send them to a configured endpoint. This document provides instructions for installing and configuring the agent, including options for using Docker or running it directly as a system service.

---

## Table of Contents
- [Prerequisites](#prerequisites)
- [Installation](#installation)
  - [Using Docker](#using-docker)
  - [Running as a System Service](#running-as-a-system-service)
- [Configuration](#configuration)
  - [Environment Variables](#environment-variables)
- [Usage](#usage)
- [Uninstallation](#uninstallation)
- [Troubleshooting](#troubleshooting)

---

## Prerequisites
- Go 1.18 or later (if building from source).
- Docker installed (if using Docker).
- Access to required environment variables: `AGENT_UUID`, `AGENT_SECRET`.

---

## Installation

### Using Docker

1. **Pull the Docker image:**

   ```bash
   docker pull public.ecr.aws/aegis-watcher/agent/infra:lastest
   ```

2. **Run the agent using Docker:**

   ```bash
   docker run -d \
     -e AGENT_UUID="your-agent-uuid" \
     -e AGENT_SECRET="your-agent-secret" \
     public.ecr.aws/d4g3e8d6/agent/infra:latest
   ```

   Replace `your-agent-uuid`, `your-agent-secret` with appropriate values.

3. **Verify the agent is running:**

   ```bash
   docker ps | grep metrics-agent
   ```

---

### Running as a System Service

1. **Download and Run the Installation Script:**

   Download the setup script that automates the installation process:

   ```bash
   curl -o install.sh https://raw.githubusercontent.com/Aegis-Watcher/agent-infra/refs/heads/main/scripts/install.sh
   chmod +x install.sh
   ./install.sh
   ```

   The script performs the following actions:
   - Builds or downloads the Metrics Agent binary.
   - Places the binary in `/usr/local/bin/`.
   - Prompts the user to input `AGENT_UUID`, `AGENT_SECRET`.
   - Creates the environment configuration file (`/etc/aegis-watcher-metrics-agent.env`).
   - Sets up and starts a systemd service for the agent.

2. **Verify the Service is Running:**

   ```bash
   sudo systemctl status aegis-watcher-metrics-agent
   ```

---

## Configuration

### Environment Variables

The agent requires the following environment variables for configuration:

- **`AGENT_UUID`** (Required): Unique identifier for the agent.
- **`AGENT_SECRET`** (Required): Secret key for agent authentication.

---

## Usage

### Logs

To view logs:

- **Docker:**
  ```bash
  docker logs <container-id>
  ```

- **System Service:**
  ```bash
  journalctl -u aegis-watcher-metrics-agent -f
  ```

### Testing

To verify the agent is collecting and sending metrics correctly, check your monitoring system for incoming data or use logs for debugging.

---

## Uninstallation

### Docker

1. Stop the container:

   ```bash
   docker stop <container-id>
   ```

2. Remove the container:

   ```bash
   docker rm <container-id>
   ```

3. Optionally, remove the image:

   ```bash
   docker rmi public.ecr.aws/aegis-watcher/agent/infra
   ```

### System Service

1. Stop and disable the service:

   ```bash
   sudo systemctl stop aegis-watcher-metrics-agent
   sudo systemctl disable aegis-watcher-metrics-agent
   ```

2. Remove the binary and configuration files:

   ```bash
   sudo rm /usr/local/bin/aegis-watcher-metrics-agent
   sudo rm /etc/aegis-watcher-metrics-agent.env
   sudo rm /etc/systemd/system/aegis-watcher-metrics-agent.service
   ```

3. Reload systemd:

   ```bash
   sudo systemctl daemon-reload
   ```

---

## Troubleshooting

- **Service Fails to Start:**
  - Ensure the environment variables are correctly set in `/etc/metrics-agent.env`.
  - Check logs for detailed error messages.
    ```bash
    journalctl -u aegis-watcher-metrics-agent
    ```

- **Metrics Not Received:**
  - Verify the network connection and endpoint availability.
  - Check logs for connectivity errors.

---

For further assistance, please contact support at support@aegiswatcher.com


