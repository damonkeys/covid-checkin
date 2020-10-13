# Checkins - Server for serving all business-checkins
Note: This server uses a different database from all other services of chckr. This is for data guard and safety reasons.
We want to be able to move the checkin data independently from the rest of the system. This said you always should have in mind
this special environment situation of the checkin server.

_A good starting point might be to read the script createDBAndUser.sh_

## Environment variables
The server uses environment variables. If they are not set the server won't start. It expects the following environment variables:
   * SERVER_PORT       - the server is listening on this portnumber
   * DB_HOST           - database host for connecting the auth database
   * DB_NAME           - database name for connecting the auth database
   * DB_USER           - database user for connecting the auth database
   * DB_PASSWORD       - database user-password for connecting the auth database
   * SESSION_SECRET    - The secret that we use for secure sessions

_Use .env file to configure this vars
