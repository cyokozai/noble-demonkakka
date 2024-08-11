#!/bin/bash

set -e
set -o pipefail

go build -o demonkakka main.go

sudo mv demonkakka /usr/local/bin/
sudo chmod +x /usr/local/bin/demonkakka

echo "Please run \"demonkakka\"."
