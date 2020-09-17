#!/bin/bash    

export SERVER_PORT=19000

export DB_HOST=""
export DB_NAME="ch3ck1n"
export DB_USER="ch3ck1n_user"
export DB_PASSWORD="==>"
# The path where the qr code images should be stores. Used in library as env var not via ServerConfigStruct
export QR_CODE_FILE_PATH="../pixi/static/qr"
# This is used during the qr code generation when we create a business.
#We encode this deeplink (with dynamic business code appended) into the qr code
export DEEP_LINK_TO_BUSINESS_BY_CODE="https://dev.checkin.chckr.de/checkin/"

./admin
