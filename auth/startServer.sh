#!/bin/bash         

export SERVER_PORT=2000

export DB_HOST=""
export DB_NAME="monkey_auth"
export DB_USER="auth_user"
export DB_PASSWORD=""

export P_FACEBOOK_KEY="fb_key"
export P_FACEBOOK_SECRET="fb_secret"
export P_GPLUS_KEY="gplus_key"
export P_GPLUS_SECRET="gplus_secret"

# static part plus dynamic replacement for the activaation token
export ACTIVATION_URL="https://checkin.chckr.de/auth/activation/%s"
export ACTIVATION_STATE_URL="https://checkin.chckr.de/activation/%s"
./bongo-auth
