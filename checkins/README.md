# Checkins - Server for serving all business-checkins

## Environment variables
The server uses environment variables. If they are not set the server won't start. It expects the following environment variables:
   * SERVER_PORT       - the server is listening on this portnumber
   * DB_HOST           - database host for connecting the auth database
   * DB_NAME           - database name for connecting the auth database
   * DB_USER           - database user for connecting the auth database
   * DB_PASSWORD       - database user-password for connecting the auth database
   * SESSION_SECRET    - The secret that we use for secure sessions
