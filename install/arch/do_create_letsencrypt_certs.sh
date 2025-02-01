#!/bin/bash

sudo pacman -S certbot --noconfirm

domain="$1"

echo "Creating letsencrypt certificates for $domain"

sudo certbot certonly --standalone -d "$domain" --register-unsafely-without-email --non-interactive --agree-tos

sudo mkdir -p /home/testd9t/certs

rm /home/testd9t/certs/*

sudo cp /etc/letsencrypt/live/$domain/fullchain.pem /home/testd9t/certs/
sudo cp /etc/letsencrypt/live/$domain/privkey.pem /home/testd9t/certs/
sudo chown -R testd9t:testd9t /home/testd9t/certs