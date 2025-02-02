#!/bin/bash

sudo cp /tmp/testd9t/testd9t.service /etc/systemd/system/testd9t.service

SOURCE_ENV="/home/testd9t/env"
TARGET_ENV="/tmp/testd9t/env"

if [[ ! -f "$SOURCE_ENV" || ! -f "$TARGET_ENV" ]]; then
  echo "Both source and target env files must exist."
  exit 1
fi

declare -A env_vars
while IFS='=' read -r key value; do
  if [[ -n "$key" && "$key" != "#"* ]]; then
    env_vars["$key"]="$value"
  fi
done < "$SOURCE_ENV"

while IFS='=' read -r key value; do
  if [[ -n "$key" && "$key" != "#"* ]]; then
    if [[ -n "${env_vars[$key]}" ]]; then
      echo "$key=${env_vars[$key]}"
    else
      echo "$key=$value"
    fi
  else
    echo "$key"
  fi
done < "$TARGET_ENV" > "${TARGET_ENV}.merged"

mv "${TARGET_ENV}.merged" "$SOURCE_ENV"

binary_path=$(pwd)/main

if [ ! -f $binary_path ]; then
    echo "Expected this binary to be at $binary_path"
    exit 1
fi

sudo cp $binary_path /home/testd9t/backend/main

sudo systemctl enable testd9t
sudo systemctl start testd9t

sudo chown -R testd9t:testd9t /home/testd9t/