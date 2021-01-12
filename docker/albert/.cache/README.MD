# Information for this folder

This folder contains the current letsencrypt server-side ssl certificate.
Do not check in the corresponding file (*acme_account+key*, *dev.checkin.chckr.de*) as it is handled automatically 
by the server itself.

The reason you need this certificate file during development time is that your host is not
reachable for the acme service with the correct dns name (dev.checkin.chckr.de u name it).

This results in a copy of the server file. Sometimes it will need an update with a news file.
Your browser will tell you when.

_If you don't know what this means to you ask a developer/ mentor_
