#!/bin/sh
set -e
cd ..
echo "Deploying application ..."

# Update codebase
echo "Update Codebase"
git fetch origin development
git reset --hard origin/development

echo "Golang Set Up"
cd
source .profile

# source /home/fanzru/.profile
# echo "Restart pm2 service 🔥"
# pm2 restart deploy.json
echo "to Directory File"
cd /home/fanzru/backend/BE-SISFO-KLINIK/

echo "Installing dependencies 🛠"
go mod tidy

echo "Restart pm2 service 🔥"
pm2 restart deploy.json

echo "Deploying Application Successfully