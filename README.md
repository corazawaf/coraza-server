# coraza-server
Whitepaper for the coraza-server implementation

## Current definitions

### One server will supply only one protocol or multiple?

### Are we supporting only one WAF instance per port ? Or we are going to identify WAF IDs using a unique waf id?

### Are we using a standard data structure for configuration or we are creating other syntax?

### Are we creating a server just for coraza? We could create a whole new different project to create multiple "processors"

### It should be possible to add protocols as plugins, maybe each plugin but REST could be a plugin

### What will happen in the future when we achieve data persistance?

### Should each protocol wrapper consider load balancing?

### Should we create a universal wrapper library (CGO) and connect it to other languages like javascript and python?

### Should we even consider wrappers? We could just expose the api endpoints

### Are we supporting the 5 phases?

Should each phase call be able to call phase 5 and logging?

- **Support all phases and helpers for 1-2-5 and 3-4-5**
- **Support all phases**
- **Support phases 1-2-5 and 3-4-5**

### Are we supporting reloading? How?

### Are we adding auto CRS loading?

### How do we handle transaction timeouts in case it takes too long or it doesn't get closed?

### How do we handle system logging? Should we wrap the Coraza logs to it?

### Should we add Kubernetes features?


## Supported protocols

### GRPC

[Link to proto](#)

### REST

[Link to OpenAPI](#) - [Link to Postman](#)

### SPOP

[Link to specification](#)
