#!/usr/bin/env bash

set -euo pipefail

APP_NAME="xorbit"
INSTALL_PATH="/usr/local/bin/xorbit"

CONFIG_DIR="/etc/xorbit"
CONFIG_FILE="${CONFIG_DIR}/config.json"

SERVICE_FILE="/etc/systemd/system/${APP_NAME}.service"

GITHUB_USER="amintoorchi"
GITHUB_REPO="xorbit-agent"


echo "================================"
echo " Installing XOrbit"
echo "================================"


# Check root
if [ "$EUID" -ne 0 ]; then
    echo "❌ Please run as root"
    echo "Example:"
    echo "curl -fsSL URL/install.sh | sudo bash -s TOKEN"
    exit 1
fi


# Token
TOKEN="${1:-}"

if [ -z "$TOKEN" ]; then
    echo "❌ Agent token missing"
    exit 1
fi


# Check systemd
if ! command -v systemctl >/dev/null 2>&1; then
    echo "❌ systemd is required"
    exit 1
fi


# Install curl
if ! command -v curl >/dev/null 2>&1; then

    echo "Installing curl..."

    if command -v apt >/dev/null 2>&1; then
        apt update
        apt install -y curl

    elif command -v yum >/dev/null 2>&1; then
        yum install -y curl

    else
        echo "❌ Cannot install curl"
        exit 1
    fi
fi


# Detect architecture
ARCH=$(uname -m)

case "$ARCH" in

    x86_64)
        BINARY="xorbit-linux-amd64"
        ;;

    aarch64|arm64)
        BINARY="xorbit-linux-arm64"
        ;;

    *)
        echo "❌ Unsupported architecture: $ARCH"
        exit 1
        ;;

esac


DOWNLOAD_URL="https://github.com/${GITHUB_USER}/${GITHUB_REPO}/releases/latest/download/${BINARY}"


echo ""
echo "Architecture: $ARCH"
echo "Downloading: $BINARY"


mkdir -p "$CONFIG_DIR"


# Stop old service
if systemctl list-unit-files | grep -q "${APP_NAME}.service"; then
    systemctl stop "$APP_NAME" || true
fi


# Download binary
curl -fL "$DOWNLOAD_URL" \
    -o "$INSTALL_PATH"


chmod +x "$INSTALL_PATH"


echo "Installed binary:"
echo "$INSTALL_PATH"


# Config

cat > "$CONFIG_FILE" <<EOF
{
    "token": "$TOKEN"
}
EOF

chmod 600 "$CONFIG_FILE"



# Systemd service

cat > "$SERVICE_FILE" <<EOF
[Unit]
Description=XOrbit Monitoring Agent
After=network-online.target
Wants=network-online.target


[Service]
Type=simple

ExecStart=${INSTALL_PATH}

Restart=always
RestartSec=5

User=root

StandardOutput=journal
StandardError=journal


[Install]
WantedBy=multi-user.target
EOF



systemctl daemon-reload

systemctl enable "$APP_NAME"

systemctl restart "$APP_NAME"


echo ""
echo "================================"
echo " XOrbit installed successfully"
echo "================================"

systemctl status "$APP_NAME" --no-pager