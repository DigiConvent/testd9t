#!/bin/bash

export DEBIAN_FRONTEND=noninteractive
sudo pacman -S postgresql postgresql-contrib --noconfirm > /dev/null

systemctl start postgresql
systemctl enable postgresql