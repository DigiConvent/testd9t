certbot certonly --standalone -d $1 --register-unsafely-without-email --non-interactive --agree-tos

mkdir /etc/digiconvent/certs
sudo cp /etc/letsencrypt/live/digiconvent.de/fullchain.pem /etc/digiconvent/certs/
sudo cp /etc/letsencrypt/live/digiconvent.de/privkey.pem /etc/digiconvent/certs/
chmod 777 /etc/digiconvent/certs/*