package dnspod

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/libdns/tencentcloud"
)

// Provider lets Caddy read and manipulate DNS records hosted by this DNS provider.
type Provider struct{ *tencentcloud.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.dnspod",
		New: func() caddy.Module { return &Provider{new(tencentcloud.Provider)} },
	}
}

// Provision sets up the module. Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	repl := caddy.NewReplacer()
	p.Provider.SecretId = repl.ReplaceAll(p.Provider.SecretId, "")
	p.Provider.SecretKey = repl.ReplaceAll(p.Provider.SecretKey, "")
	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
//	dnspod {
//	  secret_id <secret_id>
//	  secret_key <secret_key>
//	}
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "secret_id":
				if d.NextArg() {
					p.Provider.SecretId = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "secret_key":
				if d.NextArg() {
					p.Provider.SecretKey = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Provider.SecretId == "" || p.Provider.SecretKey == "" {
		return d.Err("missing SecretId or SecretKey")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
