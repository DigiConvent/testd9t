echo "________________________________________________________________________"
echo "________________________________________________________________________"
echo "________________________________________________________________________"
echo "________________________________________________________________________"
echo "________________________________________________________________________"
echo "________________________________________________________________________"
echo "________________________________________________________________________"
echo "________________________________________________________________________"
echo "________/\\\\\\\\\\\\\\\\\\\\\\\\__________/\\\\\\\\\\\\\\\\\\______/\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\_________"
echo "________\\/\\\\\\////////\\\\\\______/\\\\\\///////\\\\\\___\\///////\\\\\\/////_________"
echo "_________\\/\\\\\\______\\//\\\\\\____/\\\\\\______\\//\\\\\\________\\/\\\\\\_____________"
echo "__________\\/\\\\\\_______\\/\\\\\\___\\//\\\\\\_____/\\\\\\\\\\________\\/\\\\\\____________"
echo "___________\\/\\\\\\_______\\/\\\\\\____\\///\\\\\\\\\\\\\\\\/\\\\\\________\\/\\\\\\___________"
echo "____________\\/\\\\\\_______\\/\\\\\\______\\////////\\/\\\\\\________\\/\\\\\\__________"
echo "_____________\\/\\\\\\_______/\\\\\\_____/\\\\________/\\\\\\_________\\/\\\\\\_________"
echo "______________\\/\\\\\\\\\\\\\\\\\\\\\\\\/_____\\//\\\\\\\\\\\\\\\\\\\\\\/__________\\/\\\\\\________"
echo "_______________\\////////////________\\///////////____________\\///________"
echo "________________________________________________________________________"
echo "________________________________________________________________________"
echo "________________________________________________________________________"
echo "________________________________________________________________________"
echo "________________________________________________________________________"
echo "________________________________________________________________________"
echo "________________________________________________________________________"
echo "________________________________________________________________________"


is_installed() {
    local package_name="$1"
    if dpkg-query -W -f='${Status}' "$package_name" 2>/dev/null | grep -q "installed"; then
        return 0
    else
        return 1
    fi
}

REPO_URL="https://github.com/DigiConvent/d9ttest.git"
APP_USER="digiconvent"
APP_GROUP="digiconvent_group"
APP_NAME="digiconvent"
APP_BIN="/usr/local/bin/$APP_NAME"
STATIC_FILES="/var/www/$APP_NAME"
HOME_DIR="/home/digiconvent/"
LOG_DIR="/var/log/$APP_NAME"
CONFIG_DIR="/etc/$APP_NAME"

password=$(tr -dc 'A-Za-z0-9' < /dev/urandom | head -c 16)

if [ ! $# -eq 2 ]
  then
    echo "Call this script with the Telegram bot token and domain"
    echo "./install.sh <version> <token> <domain>"
    echo "    e.g., ./install.sh 0.1.0 1224455679:AAD1448aBCCeeFFggJKKlMMMRRssVVVwyZ_ digiconvent.de"
    exit 1
fi

tg_bot_token=$1
domain=$2

SECONDS=0

if ! id "$APP_USER" &>/dev/null; then
    echo "Step 1.0 User $APP_USER does not exist, creating..."
    useradd -m "$APP_USER"
    groupadd "$APP_GROUP"
    usermod -aG $APP_GROUP $APP_USER
    usermod -aG sudo $APP_USER
else
    echo "Skipping step 1.0 (create user $APP_USER)"
fi

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


echo "Step 3.0 Setting up database"

systemctl start postgresql
systemctl enable postgresql

su - postgres -c "
psql -c \"CREATE ROLE d9t WITH LOGIN PASSWORD '$password';\"
psql -c \"ALTER ROLE d9t CREATEDB;\"
createdb -E UTF8 -O d9t digiconvent;
psql -c \"ALTER ROLE d9t WITH SUPERUSER;\"
" > /dev/null



if [ "$1" == "latest" ]; then
    TAG=$(git ls-remote --tags --sort="v:refname" $REPO_URL | tail -n1 | sed 's/.*\///')
else
    TAG=$1
fi

if ! git ls-remote --tags $REPO_URL | grep -q "refs/tags/$TAG"; then
    echo "Tag $TAG does not exist"
    exit 1
fi

RELEASE_DATA=$(curl -s "https://api.github.com/repos/DigiConvent/d9t/releases/tags/$TAG")
ASSET_URL=$(echo "$RELEASE_DATA" | jq -r '.assets[0].browser_download_url')
echo $ASSET_URL
if [ -z "$ASSET_URL" ]; then
    echo "No assets found for release $TAG"
    exit 1
fi

cd /tmp/
curl -L -o "release.zip" "$ASSET_URL"
rm -rf release/
unzip release.zip

cp -r /tmp/release/ /tmp/



# at this point, the following folder/file structure should be guaranteed:

# /tmp/server
# /tmp/migrations/
# /tmp/migrations/x.y.z/#_description.sql
# /tmp/frontend/index.html
# /tmp/frontend/favicon.ico
# /tmp/frontend/assets
# /tmp/frontend/assets/{component}.js
# /tmp/frontend/assets/{component}.css

mkdir -p $STATIC_FILES
mkdir -p $STATIC_FILES/frontend
mkdir -p $CONFIG_DIR
mkdir -p $LOG_DIR
mkdir -p $HOME_DIR

touch $CONFIG_DIR/env
echo "APP_NAME=$APP_NAME" >> $CONFIG_DIR/env
echo "APP_BIN=$APP_BIN" >> $CONFIG_DIR/env
echo "APP_USER=$APP_USER" >> $CONFIG_DIR/env
echo "APP_GROUP=$APP_GROUP" >> $CONFIG_DIR/env
echo "STATIC_FILES=$STATIC_FILES" >> $CONFIG_DIR/env
echo "CONFIG_DIR=$CONFIG_DIR/" >> $CONFIG_DIR/env
echo "LOG_DIR=$LOG_DIR" >> $CONFIG_DIR/env
echo "UPLOADS_DIR=${STATIC_FILES}/uploads" >> $CONFIG_DIR/env
echo "DB_HOST=localhost" >> $CONFIG_DIR/env
echo "DB_PORT=5432" >> $CONFIG_DIR/env
echo "DB_USER=d9t" >> $CONFIG_DIR/env
echo "DB_PASSWORD=$1" >> $CONFIG_DIR/env
echo "DB_NAME=digiconvent" >> $CONFIG_DIR/env
echo "PORT=443" >> $CONFIG_DIR/env
echo "TELEGRAM_BOT_TOKEN=$2" >> $CONFIG_DIR/env
echo "DOMAIN=digiconvent.de" >> $CONFIG_DIR/env

chown -R $APP_USER:$APP_GROUP $STATIC_FILES $CONFIG_DIR $LOG_DIR $HOME_DIR
chmod -R 750 $STATIC_FILES $CONFIG_DIR $LOG_DIR $HOME_DIR

cp /tmp/server $APP_BIN
chown root:root $APP_BIN
chmod 755 $APP_BIN

cp -r /tmp/frontend/* $STATIC_FILES/frontend/
cp -r /tmp/migrations $CONFIG_DIR/

SERVICE_FILE="/etc/systemd/system/$APP_NAME.service"
cat <<EOL > $SERVICE_FILE
[Unit]
Description=$APP_NAME Service
After=network.target

[Service]
ExecStart=$APP_BIN
WorkingDirectory=$HOME_DIR
Restart=always
User=$APP_USER
Group=$APP_GROUP
EnvironmentFile=$CONFIG_DIR/env
StandardOutput=journal
StandardError=journal
AmbientCapabilities=CAP_NET_BIND_SERVICE

[Install]
WantedBy=multi-user.target
EOL
sudo setcap 'cap_net_bind_service=+ep' $APP_BIN


echo "$APP_NAME installation completed."


certbot certonly --standalone -d $domain --register-unsafely-without-email --non-interactive --agree-tos

mkdir /etc/digiconvent/certs
sudo cp /etc/letsencrypt/live/digiconvent.de/fullchain.pem /etc/digiconvent/certs/
sudo cp /etc/letsencrypt/live/digiconvent.de/privkey.pem /etc/digiconvent/certs/
chmod 777 /etc/digiconvent/certs/*





# allow d9t to host its own dns records alongside systemd-resolved
sudo ip addr add 192.0.2.2/24 dev eth0
systemctl daemon-reload
systemctl enable $APP_NAME
systemctl start $APP_NAME





duration=$SECONDS
echo "$((duration / 60)) minutes and $((duration % 60)) seconds elapsed."

journalctl -u digiconvent -f