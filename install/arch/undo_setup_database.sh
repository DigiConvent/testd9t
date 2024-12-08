#!/bin/bash

sudo psql -U postgres -c "drop database digiconvent;"
sudo psql -U postgres -c "drop owned by digiconvent;"
sudo psql -U postgres -c "drop role digiconvent;"