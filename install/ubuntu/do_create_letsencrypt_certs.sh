#!/bin/bash

sudo apt-get update -y
sudo apt-get install -y certbot

domain="$1"

echo "Creating letsencrypt certificates for $domain"

sudo mkdir -p /home/testd9t/certs

sudo certbot certonly --webroot -w /home/testd9t/certs -d "$domain"


rm /home/testd9t/certs/*

sudo cp /etc/letsencrypt/live/$domain/fullchain.pem /home/testd9t/certs/
sudo cp /etc/letsencrypt/live/$domain/privkey.pem /home/testd9t/certs/

sudo chown -R testd9t:testd9t /home/testd9t/certs