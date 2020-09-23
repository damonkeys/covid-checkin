#!/bin/bash         

export SERVER_PORT=2000

export DB_HOST=""
export DB_NAME="ch3ck1n"
export DB_USER="ch3ck1n_user"
export DB_PASSWORD="==>"
export SESSION_SECRET="ThatsASessionStoreSecret"
export BASE_URL="https://dev.checkin.chckr.de"

export P_FACEBOOK_KEY=""
export P_FACEBOOK_SECRET=""
export P_GPLUS_KEY=""
export P_GPLUS_SECRET=""
export P_APPLE_KEY=""
export P_APPLE_SECRET=""

# static part plus dynamic replacement for the activation token
export ACTIVATION_URL="https://dev.checkin.chckr.de/auth/activation/%s"
export ACTIVATION_STATE_URL="https://dev.checkin.chckr.de/activation/%s"

./authx
