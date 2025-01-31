#!/bin/bash

sudo apt-get install -y certbot python3-certbot-nginx

domain="$1"

echo "Creating letsencrypt certificates for $domain"

sudo certbot certonly --standalone -d "$domain" --register-unsafely-without-email --non-interactive --agree-tos

sudo mkdir -p /home/digiconvent/certs

sudo cp /etc/letsencrypt/live/$domain/fullchain.pem /home/digiconvent/certs/
sudo cp /etc/letsencrypt/live/$domain/privkey.pem /home/digiconvent/certs/
chmod 777 /etc/digiconvent/certs/*