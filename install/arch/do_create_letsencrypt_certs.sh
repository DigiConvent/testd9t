#!/bin/bash

sudo pacman -S certbot --noconfirm

domain="$1"

echo "Creating letsencrypt certificates for $domain"

sudo mkdir -p /home/testd9t/certs
rm /home/testd9t/certs/*

sudo certbot certonly --webroot -w /home/testd9t/certs -d "$domain"


sudo cp /etc/letsencrypt/live/$domain/fullchain.pem /home/testd9t/certs/
sudo cp /etc/letsencrypt/live/$domain/privkey.pem /home/testd9t/certs/