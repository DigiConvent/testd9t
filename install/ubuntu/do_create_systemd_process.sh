#!/bin/bash

sudo cp /tmp/testd9t/testd9t.service /etc/systemd/system/testd9t.service
sudo cp /tmp/testd9t/env /home/testd9t/env

binary_path=$(pwd)/main

if [ ! -f $binary_path ]; then
    echo "Expected this binary to be at $binary_path"
    exit 1
fi

sudo cp $binary_path /home/testd9t/backend/main

sudo systemctl enable testd9t
sudo systemctl start testd9t

sudo chown -R testd9t:testd9t /home/testd9t/