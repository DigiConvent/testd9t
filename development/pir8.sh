#!/bin/bash

# expect 2 parameters: github username and repository name
new_github_username=$1
new_repository_name=$2

old_repository_name="testd9t"
old_github_username="digiconvent"

function replace_in_files() {
  local dir="$(pwd)"

  grep -rni "$old_repository_name" "$dir" 2>/dev/null
#   find "$dir" -type f | while IFS= read -r file; do
#     line="$(grep -Hn "$old_repository_name" "$file" 2>/dev/null)"
#     if [ -z "$line" ]; then
#       becomes=$(echo "$line" | sed "s/$old_repository_name/$new_repository_name/g")
#       if [ "$line" != "$becomes" ]; then
#         echo "$file"
#         echo "$line"
#         echo "$becomes"
#       fi
#     fi
#   done
#   find . -type f -exec sed -i "s/DigiConvent/$github_username/g" {} +
#   find . -type f -exec sed -i "s/testd9t/$repository_name/g" {} +
}

function rename_files() {
  local dir="$(pwd)"

  echo $(pwd);

  find "$dir" -depth -name "*$old_repository_name*" | while IFS= read -r item; do
    # Get the directory and base name
    parent_dir=$(dirname "$item")
    base_name=$(basename "$item")
    
    # Replace the old character with the new character
    new_name="${base_name//"$old_repository_name"/"$new_repository_name"}"
    
    # Rename the file or directory
    # mv "$item" "$parent_dir/$new_name"
    if [ "$base_name" != "$new_name" ]; then
    #   mv "$item" "$parent_dir/$new_name"
        echo "$parent_dir/$base_name -> $parent_dir/$new_name"
    fi
  done
}

rename_files
replace_in_files