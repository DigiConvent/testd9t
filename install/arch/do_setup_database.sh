#!/bin/bash

password="$(cat /tmp/password.txt | xargs)"

psql -U postgres -c "create role digiconvent with login password '$password';alter role digiconvent;create database digiconvent owner digiconvent;"