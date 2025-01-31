migrations_folder='data/migrations/0.0.0'
migrations=$(find "$migrations_folder" -type f -name "*.sql" | while read -r file; do
  filename=$(basename "$file")
  echo "{\"name\":\"$filename\",\"url\":\"$REPO_URL/$file\"}"
done | paste -sd, -)
migrations="[$migrations]"

echo $migrations