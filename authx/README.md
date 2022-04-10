# Bongo - Auth-Server

Auth-Server for chckr.de.

## Providers for oauth-login
Actuall providers
- Facebook
- Google+

## Environment variables
The server uses environment variables. If they are not set the server won't start. It expects the following environment variables:
   * SERVER_PORT       - the server is listening on this portnumber
   * DB_HOST           - database host for connecting the auth database
   * DB_NAME           - database name for connecting the auth database
   * DB_USER           - database user for connecting the auth database
   * DB_PASSWORD       - database user-password for connecting the auth database
   * P_FACEBOOK_KEY    - Facebook-key for login-provider (goth)
   * P_FACEBOOK_SECRET - Facebook-secret for login-provider (goth)
   * P_GPLUS_KEY    	 - Google+-key for login-provider (goth)
   * P_GPLUS_SECRET    - Google+-secret for login-provider (goth)

## Database-connection for Maria-DB
It stores now in MySQL Database. Use the following command to create the database:
```
mysql -u root -p < database/createDatabase.sql
```
There is a file named 'database/createDatabaseTest.sql' for a testing database. It is needed for some unit-test.

---

*Fun fact: Bongo is the bouncer of the Ink and Paint Club in the movie "Who Framed Roger Rabbit".*
