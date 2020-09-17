# Admin - simple database frontend

Simple database-web-views for businesses and more.

## Environment variables
The server uses environment variables. If they are not set the server won't start. It expects the following environment variables:
   * QR_CODE_FILE_PATH  - The relative path to pixi where the generated QR codes should be places into
   * DEEP_LINK_TO_BUSINESS_BY_CODE - Describes the deep link for our frontend to show a business. The dynamic part is the 5-letter code which is appended dynamically during runtime
