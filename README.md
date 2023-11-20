# mailu-quadlet
An attempt to make mailu podman quadlet configurations more automagic

Just get your docker-compose.yml and mailu.env per usual from https://setup.mailu.io and download them

Then feed them to the program, e.g.

```
$ wget .../docker-compose.yml
$ wget .../mailu.env
$ podman run -v $(pwd):/data ghcr.io/guest42069/mailu-quadlet -help
Usage of /cli:
  -compose string
        docker-compose.yml file for mailu (default "docker-compose.yml")
  -envfile string
        mailu.env file for mailu (default "mailu.env")
  -uuid string
        optional custom uuid to use for generated
$ podman run run -v $(pwd):/data:z ghcr.io/guest42069/mailu-quadlet --uuid example
2023/11/06 15:54:35 example-snappymail.volume
2023/11/06 15:54:35 example-dav.volume
2023/11/06 15:54:35 example-certs.volume
2023/11/06 15:54:35 example-mailqueue.volume
2023/11/06 15:54:35 example-mail.volume
2023/11/06 15:54:35 example-filter.volume
2023/11/06 15:54:35 example-rspamd.volume
2023/11/06 15:54:35 example-fetchmail.volume
2023/11/06 15:54:35 example-redis.volume
2023/11/06 15:54:35 example-nginx.volume
2023/11/06 15:54:35 example-webmail.volume
2023/11/06 15:54:35 example-data.volume
2023/11/06 15:54:35 example-dkim.volume
2023/11/06 15:54:35 example-postfix.volume
2023/11/06 15:54:35 example-dovecot.volume
2023/11/06 15:54:35 example-default.network
2023/11/06 15:54:35 example-radicale.network
2023/11/06 15:54:35 example-webmail.network
2023/11/06 15:54:35 example-noinet.network
2023/11/06 15:54:35 example-redis.container
2023/11/06 15:54:35 example-imap.container
2023/11/06 15:54:35 example-front.container
2023/11/06 15:54:35 example-antispam.container
2023/11/06 15:54:35 example-antivirus.container
2023/11/06 15:54:35 example-webdav.container
2023/11/06 15:54:35 example-resolver.container
2023/11/06 15:54:35 example-oletools.container
2023/11/06 15:54:35 example-fetchmail.container
2023/11/06 15:54:35 example-webmail.container
2023/11/06 15:54:35 example-admin.container
2023/11/06 15:54:35 example-smtp.container
2023/11/06 15:54:35 example.env
```

Then move the quadlets to the appropriate directory, perform a `daemon-reload`, and then either manually start the containers, or reboot as they will start automatically at boot once in place.
