# Multi Proxy Test

## Objective

Create a dynamically scaling multi proxy setup that I can test before trying to implement similar system for nimbusmc.

- [ ] Kubernetes deployment for testing
- [ ] Proxy controller in Go the can create and delete proxies through player tracking (Go)
- [ ] DNS controller to update cloudflare SRV records with low TTL
- [ ] A way to transfer players to new proxy when they change servers 
- [ ] Manage proxy state (RUNNING, DRAINING, ??)
- [ ] Tests (it would be a good idea to test thoroughly here as the actual deployment wont be able to be tested until prod deployment)


