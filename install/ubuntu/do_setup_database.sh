#!/bin/bash

if [[ "$1" =~ \s ]]; then
  echo "Error: Password contains whitespace." >&2
  exit 1
fi

psql -U postgres -c "create role digiconvent with login password '$1';alter role digiconvent;create database digiconvent owner digiconvent;"