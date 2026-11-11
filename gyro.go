// Package gyro is a Caddy module (dns.providers.gyro) that solves ACME DNS-01
// challenges with Gyro DNS, wildcard certificates included.
package gyro

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	libdnsgyro "github.com/Gyro-Hosting/libdns-gyro"
)

// Provider wraps the libdns provider for Caddy.
type Provider struct{ *libdnsgyro.Provider }

func init() { caddy.RegisterModule(Provider{}) }

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.gyro",
		New: func() caddy.Module { return &Provider{new(libdnsgyro.Provider)} },
	}
}

// Provision expands placeholders such as {env.GYRO_TOKEN}.
func (p *Provider) Provision(ctx caddy.Context) error {
	repl := caddy.NewReplacer()
	p.Provider.APIToken = repl.ReplaceAll(p.Provider.APIToken, "")
	p.Provider.APIBase = repl.ReplaceAll(p.Provider.APIBase, "")
	return nil
}

// UnmarshalCaddyfile sets up the provider from Caddyfile tokens:
//
//	gyro [<api_token>] {
//	    api_token <api_token>
//	    api_base  <url>
//	    skip_wait
//	}
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			p.Provider.APIToken = d.Val()
		}
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "api_token":
				if !d.NextArg() {
					return d.ArgErr()
				}
				p.Provider.APIToken = d.Val()
			case "api_base":
				if !d.NextArg() {
					return d.ArgErr()
				}
				p.Provider.APIBase = d.Val()
			case "skip_wait":
				p.Provider.SkipWait = true
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
			if d.NextArg() {
				return d.ArgErr()
			}
		}
	}
	if p.Provider.APIToken == "" {
		return d.Err("missing API token")
	}
	return nil
}

var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
