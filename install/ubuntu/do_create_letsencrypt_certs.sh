#!/bin/bash

sudo apt-get install -y certbot python3-certbot-nginx

certbot certonly --standalone -d $1 --register-unsafely-without-email --non-interactive --agree-tos

sudo cp /etc/letsencrypt/live/digiconvent.de/fullchain.pem /home/digiconvent/certs/
sudo cp /etc/letsencrypt/live/digiconvent.de/privkey.pem /home/digiconvent/certs/
chmod 777 /etc/digiconvent/certs/*