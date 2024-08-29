echo "Step 3.0 Setting up database"

systemctl start postgresql
systemctl enable postgresql

su - postgres -c "
psql -c \"CREATE ROLE d9t WITH LOGIN PASSWORD '$1';\"
psql -c \"ALTER ROLE d9t CREATEDB;\"
createdb -E UTF8 -O d9t digiconvent;
psql -c \"ALTER ROLE d9t WITH SUPERUSER;\"
" > /dev/null