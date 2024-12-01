#!/bin/bash

export DEBIAN_FRONTEND=noninteractive
apt-get install -y postgresql postgresql-contrib > /dev/null

systemctl start postgresql
systemctl enable postgresql