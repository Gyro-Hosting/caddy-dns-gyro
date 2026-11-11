# caddy-dns-gyro

[Caddy](https://caddyserver.com) module **`dns.providers.gyro`**: get Let's Encrypt certificates — wildcards included —
with the ACME DNS-01 challenge on **Gyro DNS**. Built on [libdns-gyro](https://github.com/Gyro-Hosting/libdns-gyro).

*Türkçe açıklama aşağıda.*

## Build
```
xcaddy build --with github.com/Gyro-Hosting/caddy-dns-gyro
```

## ACME key
Gyro DNS service page → **ACME / Let's Encrypt** → *Create key*. The key can only add and remove `_acme-challenge`
TXT records of the zones you select — nothing else. Keep it in an environment variable.

## Caddyfile
```
*.example.com, example.com {
	tls {
		dns gyro {env.GYRO_TOKEN}
	}
	respond "hello"
}
```
Global (all sites):
```
{
	acme_dns gyro {env.GYRO_TOKEN}
}
```
Options:
```
dns gyro {
	api_token {env.GYRO_TOKEN}
	api_base  https://www.gyrohosting.com/api/v1/client/   # optional
	skip_wait                                            # optional: do not wait for Gyro DNS name servers
}
```
JSON: `{"module": "gyro", "api_token": "{env.GYRO_TOKEN}"}`.

## Türkçe
Caddy için Gyro DNS modülü: joker sertifikalar dahil Let's Encrypt sertifikalarını DNS doğrulamasıyla otomatik alır.
`xcaddy build --with github.com/Gyro-Hosting/caddy-dns-gyro` ile derleyin; Gyro DNS hizmet sayfasındaki
**ACME / Let's Encrypt** bölümünden oluşturduğunuz ACME anahtarını `GYRO_TOKEN` ortam değişkenine yazın ve Caddyfile'da
`tls { dns gyro {env.GYRO_TOKEN} }` kullanın. ACME anahtarı yalnız `_acme-challenge` TXT kaydı ekleyip silebilir.

## License
MIT
