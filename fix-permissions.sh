#!/bin/bash

# Fix permissions for Docker volume mounts
# Run this script with: chmod +x fix-permissions.sh && sudo ./fix-permissions.sh

echo "🔧 Fixing permissions for PInGO Docker volumes..."

# Create directories if they don't exist
mkdir -p ./uploads ./backend/avatars ./backend/backgrounds ./backend/logos

# Change ownership to match Docker container (UID 1000)
sudo chown -R 1000:1000 ./uploads ./backend/avatars ./backend/backgrounds ./backend/logos

# Ensure read permissions for files
sudo find ./backend/avatars ./backend/backgrounds ./backend/logos -type f -exec chmod 644 {} \;
sudo find ./backend/avatars ./backend/backgrounds ./backend/logos -type d -exec chmod 755 {} \;

sudo chmod 755 ./uploads

echo "✅ Permissions fixed! Container should now be able to access static files."
echo "📝 Note: You may need to restart your Docker containers after this change."
