# covid-checkin - a location based checkin app ready to run

## about

### What it is

This repository contains a covid/corona related checkin app that allows users to checkin at a given location.
This may help in tracking the location of potentially infected persons at the given location.

It is a feature complete app and its frontend is optimized for mobile.
### History

Coming from a financial transactions background in spring 2020 we reckognized that there hasn't been any good checkin app in Germany and that we have had a lot of the ingredients neccessary to build one at hand. So we decided to create chckr.de (the original name of the 'product') in our spare time. It turned out that we had been a little late to the party and that other took the "covid checkin biz" more seriously. So we called it "an app" and stopped our efforts right after we finished this MVP.

## features
* Mobile (responsive) frontend with pwa features
* Backend with admin panel
* QR-code generation
* Backend for companies using this app for checkins (kind of yellow pages)
* API-gateway
* Short-Code generation (5-letter code to use if no QR-code reader is at hand)
* SSH termination through Caddy
* JWT based auth (Apple, Facebook, Google)
* Static landing page server (with automatic file embedding)
* Microservice like architecture
* Database-Migrations included
* Persistence via mariadb
* Full i18n integrated (english and german)
* Jaeger tracing integrated
* Everything containerized
* Reay for use with Docker stacks ('lightweight' docker container orchestrator)
* Ready to get deployed via github actions
* Hetzner Cloud prepared

## Screenshots

### Checkin app main screen - optional working social media logins
[<img alt="Checkin app main screen - optional working social media logins" width="250px" src="/doc/images/working-social-logins.png" />](/doc/images/working-social-logins.png)
### Checkin app main screen (mobile) - with 5 letter code checking method
[<img alt="Checkin app main screen - 5 letter code" src="/doc/images/checkin-via-5letter-code.png" width="250px" />](/doc/images/checkin-via-5letter-code.png)
### Checkin process: Filling the form (mobile)
[<img alt="Checkin app - filling the form during checking" src="/doc/images/checkin-form-mobile.png" width="250px" />](/doc/images/checkin-form-mobile.png)
### Checkin process: Read business infos before checkin in (mobile)
[<img alt="Checkin app - Read business infos before checkin in" src="/doc/images/business-infos-before-checkin.png" width="250px" />](/doc/images/business-infos-before-checkin.png)
### Checkin process: Successful checkin (mobile)
[<img alt="Checkin app - Successful checkin" src="/doc/images/successful-checkin-message.png" width="250px" />](/doc/images/successful-checkin-message.png)
### Administrative backend
[<img alt="Main Page of backoffice admin panel for managing businesses" src="/doc/images/backoffice-business-administration-main.png" width="250px" />](/doc/images/backoffice-business-administration-main.png)

[<img alt="Detail Page of backoffice admin panel for managing businesses" src="/doc/images/backoffice-business-adminstration-detail.png" width="250px" />](/doc/images/backoffice-business-adminstration-detail.png)
### Landing page - i18n included
[<img alt="Full Screenshot of landing page" src="/doc/images/landingpage-full-screen.png" width="250px" />](/doc/images/landingpage-full-screen.png)
