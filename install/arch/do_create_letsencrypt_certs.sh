#!/bin/bash

sudo pacman -S certbot --noconfirm

domain="$(cat /tmp/d9t/domain.txt | xargs)"

echo "Creating letsencrypt certificates for $domain"

sudo certbot certonly --standalone -d "$domain" --register-unsafely-without-email --non-interactive --agree-tos

sudo mkdir -p /home/digiconvent/certs

sudo cp /etc/letsencrypt/live/$domain/fullchain.pem /home/digiconvent/certs/
sudo cp /etc/letsencrypt/live/$domain/privkey.pem /home/digiconvent/certs/
sudo chown -R digiconvent:digiconvent /home/digiconvent/certs