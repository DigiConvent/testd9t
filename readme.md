## Installation

```bash
https://github.com/DigiConvent/testd9t/releases/download/1.1.1/release.zip
```

## Local development

### Clone the repository and install dependencies

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
```

## Paths

| Description | Value |
|-------------|-------|
| Binary path | `/usr/local/bin/digiconvent/` |
| Service file | `/etc/systemd/system/digiconvent.service` |
| Certificates | `/etc/digiconvent/certs/` |
| Environment variables | `/etc/digiconvent/env` |