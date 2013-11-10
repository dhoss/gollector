# Cirgonus

Cirgonus is a go-related pun on [Circonus](http://circonus.com) and is a
metrics collector for it (and anything else that can deal with json output). It
also comes with `cstat`, a platform-independent `iostat` alike for gathering
cirgonus metrics from many hosts.

Most of the built-in collectors are linux-only for now, and probably the future
unless pull requests happen. Many plugins very likely require a 3.0 or later
kernel release due to dependence on system structs and other deep voodoo.

Cirgonus does not need to be run as root to collect its metrics.

Unlike other collectors that use fat tools like `netstat` and `df` which can
take expensive resources on loaded systems, Cirgonus opts to use the C
interfaces directly when it can. This allows it to keep a very small footprint;
with the go runtime, it clocks in just above 5M resident and unnoticeable CPU
usage at the time of writing. The agent can sustain over 8000qps with a
benchmarking tool like `wrk`, so it will be plenty fine getting hit once per
minute, or even once per second.

## Quick Start

In the cirgonus directory on a Linux machine with kernel 3.0 or better:

```bash
$ make
$ ./cirgonus generate > cirgonus.json
$ ./cirgonus cirgonus.json &
$ ./cstat -hosts localhost -metric "load_average"
```

Should yield an array of floats that contain your current load average.

```bash
$ curl http://localhost:8000/
```

Will yield a json object of all current metrics.


## Wiki

Our [wiki](https://github.com/erikh/cirgonus/wiki) contains tons of information
on advanced configuration, usage, and even tools you can use with Cirgonus.
Check it out!

## License

* MIT (C) 2013 Erik Hollensbe
