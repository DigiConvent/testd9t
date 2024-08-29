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
