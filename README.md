# [Caddy](https://github.com/caddyserver/caddy) with [DNSPod](https://github.com/nayukidayo/caddy/pkg/dnspod) plugin

### Caddyfile example

```Groovy
{
	email aaa@bbb.ccc
	acme_dns dnspod {
		secret_id {env.SECRET_ID}
		secret_key {env.SECRET_KEY}
	}
}

*.aaa.bbb.ccc {
	@foo host foo.aaa.bbb.ccc
	handle @foo {
		reverse_proxy 10.0.16.4:3451
	}

	@bar host bar.aaa.bbb.ccc
	handle @bar {
		handle_path /asd/* {
			reverse_proxy 10.0.16.4:3452
		}
		handle_path /qwe/* {
			reverse_proxy 10.0.16.4:3453
		}
	}

	abort
}
```
