#!/bin/sh
set -e

cd ..
echo "Deploying application ..."

# Update codebase
git fetch origin master
echo "failed 1"
git reset --hard origin/master
echo "failed 2"

echo "Installing dependencies 🛠"
go mod tidy
            
echo "Restart pm2 service 🔥"
pm2 restart deploy.json

echo "Application deployed!"