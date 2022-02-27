#!/bin/sh
set -e
cd ..
echo "Deploying application ..."

# Update codebase
echo "Update Codebase"
git fetch origin development
git reset --hard origin/development

# echo "Installing dependencies 🛠"
# go mod tidy

echo "Super User Access 🔥"
sudo su -c "cd /home/fanzru/"
sudo su -c "source .profile"
# echo "Super User Access 🔥"
# Affan080701

# echo "Update Root Golang"
# cd /home/fanzru/
# echo "Golang Set Up"
# source .profile

echo "Restart pm2 service 🔥"
pm2 restart deploy.json

echo "Application deployed!k