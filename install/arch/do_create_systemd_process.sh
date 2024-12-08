#!/bin/bash

sudo cp /tmp/testd9t/testd9t.service /etc/systemd/system/testd9t.service

binary_path=$(pwd)/main

sudo cp $binary_path /home/digiconvent/backend/main

sudo systemctl enable testd9t
sudo systemctl start testd9t

sudo chown -R digiconvent:digiconvent /home/digiconvent/