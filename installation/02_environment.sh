apt-get update > /dev/null

source ./00_utils.sh

if ! is_installed "ca-certificates"; then
    echo "Step 2.0 ca-certificates is not installed. Installing..."
    apt-get install -y ca-certificates > /dev/null
else
    echo "Skipping step 2.0 (install ca-certificates)"
fi

if ! is_installed "jq" &> /dev/null; then
    echo "Step 2.1 jq is not installed. Installing..."
    apt-get install -y jq > /dev/null
else
    echo "Skipping step 2.1 (install jq)"
fi

if ! is_installed "unzip" &> /dev/null; then
    echo "Step 2.2 unzip is not installed. Installing..."
    apt-get install -y unzip > /dev/null
else
    echo "Skipping step 2.2 (install unzip)"
fi

if ! is_installed "wget" &> /dev/null; then
    echo "Step 2.3 wget is not installed. Installing..."
    apt-get install -y wget > /dev/null
else
    echo "Skipping step 2.3 (install wget)"
fi

if ! is_installed "git" &> /dev/null; then
    echo "Step 2.4 git is not installed. Installing..."
    apt-get install -y git > /dev/null
else
    echo "Skipping step 2.4 (install git)"
fi

if ! is_installed "gpg" &> /dev/null; then
    echo "Step 2.5 gnupg is not installed. Installing..."
    apt-get install -y gnupg > /dev/null
else
    echo "Skipping step 2.5 (install gnupg)"
fi

if ! is_installed "curl" &> /dev/null; then
    echo "Step 2.6 curl is not installed. Installing..."
    apt-get install -y curl > /dev/null
else
    echo "Skipping step 2.6 (install curl)"
fi

if ! is_installed "postgresql" &> /dev/null; then
    echo "Step 2.7 postgres database is not installed. Installing..."
    export DEBIAN_FRONTEND=noninteractive
    apt-get install -y postgresql postgresql-contrib > /dev/null
else
    echo "Skipping step 2.7 (install postgres database)"
fi

if ! is_installed "certbot" &> /dev/null; then
    echo "Step 2.8 certbot is not installed. Installing..."
    apt-get install -y certbot > /dev/null
else
    echo "Skipping step 2.8 (install certbot)"
fi