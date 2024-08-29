# allow d9t to host its own dns records alongside systemd-resolved
sudo ip addr add 192.0.2.2/24 dev eth0
systemctl daemon-reload
systemctl enable $APP_NAME
systemctl start $APP_NAME
