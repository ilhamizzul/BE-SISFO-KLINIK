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
sudo su


echo "Update Root Golang"
cd /home/fanzru/
echo "Golang Set Up"
source .profile

# source /home/fanzru/.profile
# echo "Restart pm2 service 🔥"
# pm2 restart deploy.json

echo "Application deployed!k