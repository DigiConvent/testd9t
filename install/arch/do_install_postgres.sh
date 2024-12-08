#!/bin/bash

export DEBIAN_FRONTEND=noninteractive
sudo pacman -S postgresql --noconfirm > /dev/null

sudo -u postgres initdb --locale en_US.UTF-8 -E UTF8 -D '/var/lib/postgres/data'

systemctl start postgresql
systemctl enable postgresql