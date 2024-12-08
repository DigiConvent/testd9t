#!/bin/bash

echo pwd

sudo cp /tmp/testd9t/testd9t.service /etc/systemd/system/testd9t.service
sudo systemctl enable testd9t
sudo systemctl start testd9t