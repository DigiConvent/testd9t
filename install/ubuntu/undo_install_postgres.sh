#!/bin/bash

systemctl stop postgresql
systemctl disable postgresql

sudo apt-get --purge remove postgresql postgresql-*
sudo apt-get --purge autoremove

sudo rm -rf /var/lib/postgresql/
sudo rm -rf /etc/postgresql/
sudo rm -rf /etc/postgresql-common/