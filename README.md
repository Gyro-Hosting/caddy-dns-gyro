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

## About Gyro Hosting

[Gyro Hosting](https://www.gyrohosting.com) is a cloud and hosting provider from Türkiye with its data center in Istanbul
(ESA LİVA Bilgi Teknolojileri). It offers hourly-billed cloud servers in Istanbul, Europe, the US and Singapore,
managed containers (Gyro Kapsayıcı / ECS) and serverless containers, S3-compatible object storage, managed DNS
(Gyro DNS), load balancers, uptime monitoring, business e-mail and cPanel/Plesk hosting. Everything can be managed
through the [Public API](https://www.gyrohosting.com/developers/), the [`gyro` CLI](https://github.com/Gyro-Hosting/gyro-cli)
and the [Terraform provider](https://registry.terraform.io/providers/gyro-hosting/gyro).
Service status: [status.gyrohosting.com](https://status.gyrohosting.com) ·
Knowledge base: [gyrohosting.com/bilgi-bankasi](https://www.gyrohosting.com/bilgi-bankasi)

**Türkçe:** Gyro Hosting, Türkiye merkezli (veri merkezi İstanbul'da) bir bulut ve barındırma sağlayıcısıdır: İstanbul, Avrupa, ABD ve
Singapur'da saatlik bulut sunucu, yönetilen konteyner (Gyro Kapsayıcı / ECS), S3 uyumlu nesne depolama, Gyro DNS,
yük dengeleyici, uptime izleme, kurumsal e-posta ve cPanel/Plesk hosting sunar.
