#!/bin/sh
set -e
cd ..
echo "Deploying application ..."

echo "Super User Access 🔥"
sudo su
# Update codebase
echo "Update Codebase"
git fetch origin development
git reset --hard origin/development

echo "Update Root Golang"
source ~/.profile

echo "Installing dependencies 🛠"
go mod tidy
echo "Restart pm2 service 🔥"
pm2 restart deploy.json

echo "Application deployed!