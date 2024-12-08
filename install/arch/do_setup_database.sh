#!/bin/bash

password="$(cat /tmp/testd9t/password.txt | xargs)"

psql -U postgres -c "create role digiconvent with login password '$password';"
psql -U postgres -c "alter role digiconvent;"
psql -U postgres -c "create database digiconvent owner digiconvent;"