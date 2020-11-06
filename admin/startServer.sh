#!/bin/bash    

source .env
export SERVER_PORT=${SERVER_PORT}

export DB_BIZ_HOST=${DB_BIZ_HOST}
export DB_BIZ_NAME=${DB_BIZ_NAME}
export DB_BIZ_USER=${DB_BIZ_USER}
export DB_BIZ_PASSWORD=${DB_BIZ_PASSWORD}
# checkins database
export DB_CHECKINS_HOST=${DB_CHECKINS_HOST}
export DB_CHECKINS_NAME=${DB_CHECKINS_NAME}
export DB_CHECKINS_USER=${DB_CHECKINS_USER}
export DB_CHECKINS_PASSWORD=${DB_CHECKINS_PASSWORD}
# The path where the qr code images should be stores. Used in library as env var not via ServerConfigStruct
export QR_CODE_FILE_PATH=${QR_CODE_FILE_PATH}
# This is used during the qr code generation when we create a business.
#We encode this deeplink (with dynamic business code appended) into the qr code
export DEEP_LINK_TO_BUSINESS_BY_CODE=${DEEP_LINK_TO_BUSINESS_BY_CODE}

./admin
