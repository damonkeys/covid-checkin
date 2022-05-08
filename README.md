# covid-checkin - a location based checkin app ready to run

## TL;DR

## about

### What it is

This repository contains a covid/corona related checkin app that allows users to checkin at a given location.
This may help in tracking the location of potentially infected persons at the given location.

It is full featured with a mobile focussed frontend
### History

Coming from a financial transactions background in spring 2020 we reckoned that there hasn't been any good checkin app in germany and that we have had a lot of the necessary ingredients neccessary to build one at hand. So we decided to create chckr.de (the original name of the 'product') in our spare time. It turned out that we had been a little late to the party and that other took the "covid checkin biz" more seriously. So we called it "an app" and stopped our efforts right after we finished this MVP like system.

## features
* mobile (responsive) frontend with pwa features
* backend with admin panel
* qr-code generation
* backend for companies using this app for checkins (kind of yellow pages)
* api-gateway
* Short-Code generation (5-letter code to use if no qr-code reader is at hand)
* caddy ssh termination
* JWT based auth
* static landing page
* microservice like architecture
* migrations ready
* mariadb as backend
* full i18n integrated (english and germany)
* jaeger tracing integrated
* all docker based
* with docker stacks ('lightweight' docker container orchestrator)

### Ideas behind it

### Technical overview
### Parts

## screenshots

* Checkin app main screen - optional working social media logins
[![Checkin app main screen - optional working social media logins](/doc/images/working-social-logins.png)](/doc/images/working-social-logins.png)
* Checkin app main screen (mobile) - with 5 letter code checking method
[![Checkin app main screen - 5 letter code](/doc/images/checkin-via-5letter-code.png)](/doc/images/checkin-via-5letter-code.png)
* Checkin process: Filling the form (mobile)
[![Checkin app - filling the form during checking](/doc/images/checkin-form-mobile.png)](/doc/images/checkin-form-mobile.png)
* Checkin process: Read business infos before checkin in (mobile)
[![Checkin app - Read business infos before checkin in](/doc/images/business-infos-before-checkin.png)](/doc/images/business-infos-before-checkin.png)
* Checkin process: Successful checkin (mobile)
[![Checkin app - Successful checkin](/doc/images/successful-checkin-message.png)](/doc/images/successful-checkin-message.png)
* Administrative backend
[![Main Page of backoffice admin panel for managing businesses](/doc/images/backoffice-business-administration-main.png)](/doc/images/backoffice-business-administration-main.png)
[![Detail Page of backoffice admin panel for managing businesses](/doc/images/backoffice-business-adminstration-detail.png)](/doc/images/backoffice-business-adminstration-detail.png)
* Landing page - i18n included
[![Full Screenshot of landing page](/doc/images/landingpage-full-screen.png)](/doc/images/landingpage-full-screen.png)
  

## how to start

* Source organisation
* Containers
* How it they are build
* How they are started
* further need
