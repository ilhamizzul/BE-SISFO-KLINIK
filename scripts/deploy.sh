#!/bin/sh
set -e
cd ..
echo "Deploying application ..."

# Update codebase
git fetch origin development
echo "failed 1"
git reset --hard origin/development
echo "failed 2"

echo "Super User Access 🔥"
sudo su

echo "Update Root Golang"
source ~/.profile

echo "Installing dependencies 🛠"
go mod tidy
echo "Restart pm2 service 🔥"
pm2 restart deploy.json

echo "Application deployed!