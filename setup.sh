# setup.sh

# Navigate to your Go application directory
cd /var/www/apigo

go clean -cache
# Run Go build
go build

#set and enable systemd file in systemd folder
cp ./setup/apigo.system /etc/systemd/system
sudo systemctl enable apigo.service

# Restart the service (assumed serviced restart command)
systemctl start apigo
