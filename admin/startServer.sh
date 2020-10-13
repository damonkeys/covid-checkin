#!/bin/bash    

export SERVER_PORT=19000

export DB_BIZ_HOST=""
export DB_BIZ_NAME="ch3ck1n"
export DB_BIZ_USER="ch3ck1n_user"
export DB_BIZ_PASSWORD="==>"
# checkins database
export DB_CHECKINS_HOST=""
export DB_CHECKINS_NAME="checkins"
export DB_CHECKINS_USER="checkins_user"
export DB_CHECKINS_PASSWORD=""
# The path where the qr code images should be stores. Used in library as env var not via ServerConfigStruct
export QR_CODE_FILE_PATH="../pixi/static/qr"
# This is used during the qr code generation when we create a business.
#We encode this deeplink (with dynamic business code appended) into the qr code
export DEEP_LINK_TO_BUSINESS_BY_CODE="https://dev.checkin.chckr.de/checkin/"

./admin
