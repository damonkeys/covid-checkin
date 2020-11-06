#!/bin/bash         

source .env
export SERVER_PORT=${SERVER_PORT}

export DB_HOST=${DB_HOST}
export DB_NAME=${DB_NAME}
export DB_USER=${DB_USER}
export DB_PASSWORD=${DB_PASSWORD}
export SESSION_SECRET=${SESSION_SECRET}
export BASE_URL=${BASE_URL}

export P_FACEBOOK_KEY=${P_FACEBOOK_KEY}
export P_FACEBOOK_SECRET=${P_FACEBOOK_SECRET}
export P_GPLUS_KEY=${P_GPLUS_KEY}
export P_GPLUS_SECRET=${P_GPLUS_SECRET}
export P_APPLE_KEY=${P_APPLE_KEY}
export P_APPLE_SECRET=${P_APPLE_SECRET}

# static part plus dynamic replacement for the activation token
export ACTIVATION_URL=${ACTIVATION_URL}
export ACTIVATION_STATE_URL=${ACTIVATION_STATE_URL}

./authx
