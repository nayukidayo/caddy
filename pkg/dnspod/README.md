# DNSPod module for Caddy

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with DNSPod accounts.

## Caddy module name

```
dns.providers.dnspod
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuers/acme/) like so:

```json
{
  "module": "acme",
  "challenges": {
    "dns": {
      "provider": {
        "name": "dnspod",
        "secret_id": "{env.SECRET_ID}",
        "secret_key": "{env.SECRET_KEY}"
      }
    }
  }
}
```

or with the Caddyfile:

```
# globally

acme_dns dnspod {
  secret_id {env.SECRET_ID}
  secret_key {env.SECRET_KEY}
}
```

```
# one site

tls {
  dns dnspod {
    secret_id {env.SECRET_ID}
    secret_key {env.SECRET_KEY}
  }
}
```

You can replace `{env.SECRET_ID}` `{env.SECRET_KEY}` with the actual auth token if you prefer to put it directly in your config instead of an environment variable.

## Authenticating

See [the associated README in the libdns package](https://github.com/libdns/tencentcloud) for important information about credentials.
