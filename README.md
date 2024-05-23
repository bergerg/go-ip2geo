IP 2 Geo
===

Expose a simple REST API on the configured port (default: 8080), for quering ip address into a Country-City response.

Requires:
 - Docker

Running tests
---

In order to run tests, open a terminal and run:
```bash
docker compose -f tests.yaml run tests
```

Running locally on Docker
---

Open a terminal and run:

```bash
docker compose up --build -d
```
