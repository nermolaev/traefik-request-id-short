# X-Request-ID plugin for Traefik

This plugin will add the X-Request-ID header with a generated UUIDv4 value (without dashes) to HTTP requests and responses, allowing downstream services to identify requests.

Based upon:
- github.com/pipe01/plugin-requestid
- github.com/gamblingpro/plugin-requestid
- github.com/mdklapwijk/traefik-plugin-request-id
