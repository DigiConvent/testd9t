#!/bin/bash

if [[ "$1" =~ \s ]]; then
  echo "Error: Password contains whitespace." >&2
  exit 1
fi

psql -U postgres -c "drop role digiconvent;drop database digiconvent;"