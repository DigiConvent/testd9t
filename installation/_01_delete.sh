rm -rf $APP_BIN $STATIC_FILES $HOME_DIR $LOG_DIR $CONFIG_DIR
rm -rf /tmp/*

systemctl stop $APP_NAME


su - postgres <<EOF
psql -c "DROP DATABASE IF EXISTS $APP_NAME;"
psql -c "DROP ROLE IF EXISTS d9t;"
EOF


rm -f /etc/systemd/system/$APP_NAME.service
