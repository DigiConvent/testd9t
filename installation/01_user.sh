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

# Check if the user exists
if ! id "$APP_USER" &>/dev/null; then
    echo "Step 1.0 User $APP_USER does not exist, creating..."
    useradd -m "$APP_USER"
    groupadd "$APP_GROUP"
    usermod -aG $APP_GROUP $APP_USER
    usermod -aG sudo $APP_USER
else
    echo "Skipping step 1.0 (create user $APP_USER)"
fi