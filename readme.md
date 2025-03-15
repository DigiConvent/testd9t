## Installation 🧙🏻‍♂️

The following will install DigiConvent version 0.0.0 on a linux machine with ubuntu

```bash
wget https://github.com/DigiConvent/testd9t/releases/download/0.0.0/main && chmod +x main && ./main --install ubuntu
```

or

```bash
rm -f main && wget https://github.com/DigiConvent/testd9t/releases/download/0.0.0/main && chmod +x main && ./main --install ubuntu
```

A complete list of all versions can be found by running `./main --versions`.

## Local development 👷🏻

Clone the repository, install the dependencies of the frontend and run the frontend in development mode

```bash
git clone https://github.com/Digiconvent/testd9t
cd testd9t/frontend
npm --version
# 10.7.0
node --version
# v18.20.4
npm audit
# found 0 vulnerabilities
npm install

npm run dev
```

Open a new terminal, navigate to the backend folder, install and migrate the backend and run the backend in development mode

```bash
cd ../backend
# flavour being one of the supported linux flavours mentioned under install/<flavour>/
go run main.go --install <flavour>
# migrates using local migration files under /data/migrations/ to version .env -> VERSION
# use --force to force migrating. This will reset data in the database
go run main.go --migrate-db
go run main.go --run # runs the backend, open in the browser under localhost under port .env -> PORT
```

## Deployed development 🧪

In the case of a fresh server

```bash
apt update
apt install vsftpd golang-go -y
sudo systemctl start vsftpd.service
```

## Piracy 🏴‍☠️

```bash
git clone https://github.com/DigiConvent/testd9t
cd testd9t
./development/pir8.sh <your github username> <your github repository name>
```

### Folder/File structure

```
/home/digiconvent/
 - backend/
   - main # this is the binary that runs the backend
 - frontend/ # this is where the frontend files are stored
 - certs/ # certificates for ssl
 - data/
   - uploads # uploaded files and folders
```
