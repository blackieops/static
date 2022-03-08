# `static`

This is an extremely lightweight Go program (and Docker container) build
exclusively to serve static websites as quickly and efficiently as possible.

## Why Not {insert any other web server here}?

Simply put: they're too complicated. To limit attack surface area,
configuration complexity, and container size, it's advantageous to use a
use-case-specific program like this.

All `static` does is serve static content. We don't care about
reverse-proxying, CGI, load balancing, or an embedded Lua runtime...

## Usage

The intended deployment model is container-based.

```
$ docker run --rm -p 8080:8080 -v /path/to/www:/www blackieops/static
```

## Configuration

By default the program will try and load `config.yaml` in the same directory as
itself. Check out `config.yaml.example` to see a comprehensive example of a
config file.

You can change where the program loads the config file via a flag:

```
$ ./static -config /etc/my_other_config.yml
```

### Base Image

Generally, you'll probably not want to change settings at-runtime, and would
rather build your own container built on top of `static`.

For example, to copy in your static web files, and a custom config:

```
FROM blackieops/static:1.0.0
ADD config/static.yaml /config.yaml
ADD public/ /www
```
