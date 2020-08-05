# Service-Gateway - SSL-API-Proxy-Server

Simple SSL-Proxy-Server using echo. It uses a json (default routes.json in current dir)  config file like the following example:

```
{
    "routes": [{
        "name": "React-App",
        "path": "/*",
        "urls": [
            "http://dev.checkin.chckr.de:8080",
            "http://dev.checkin.chckr.de:8081"
        ],
        "balancer": "random",
        "rewrite": true
    }, {
        "name": "Admin-Tests",
        "path": "/admin*",
        "urls": [
            "http://dev.monkeycash.io:9000/admin/behaviour/device_infos",
            "http://dev.monkeycash.io:9000/admin/gaming/gameturns",
            "http://dev.monkeycash.io:9000/admin/auth/users"
        ],
        "balancer": "roundrobin",
        "rewrite": false
    }]
}
```

- name: name your route (echo use it!?)
- path: the path for the route
- urls: all URLs from where the websites will delivered
- balancer: Echo knows two loadbalander. "random" and "roundrobin"
- rewrite: rewrites the URL to the microservice. If it is set to true, the listening path will be cutted away by connecting the microservice.

## Environment variables
Service-Gateway uses environment variables. If they are not set the server won't start. It expects the following environment variables:
   * SERVER_PORT_SSL       - the server is listening on this portnumber and starts an HTTPS-Server
   * ROUTES_CONFIG         - path and filename where to find the routes.json config file to define all known routes

---
