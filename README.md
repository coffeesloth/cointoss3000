# cointoss3000

## Description

You need to toss a coin. But why settle for the best of 3?
**cointoss3000** gives you the **best of 3000**.

## Dependencies

* make
* go

## Building

```shell
make
```

## Usage

### Typical usage (best of 3000)

```shell
./ct3
```

Sample output:

```
heads
```

### Customise number of tosses

#### Simple customisation

Use `-t <number of tosses>`. For example, if you want the best of 3,

```shell
./ct3 -t 3
```

#### Extreme customisation

Of course, as the proud owner of cointoss3000, you won't want to limit
yourself to a mere _three_ coin tosses.

cointoss3000 is capable of tossing a coin up to 2^64 - 1 times.
**This will take a long time.**

```shell
./ct3 -t 18446744073709551615
```

#### Overflows not allowed

Overflowing the `uint64` is not allowed. Sorry, Matt Parker.

```shell
# not allowed
./ct3 -t 18446744073709551616
```

```shell
# not allowed
./ct3 -t -1
```

### Verbose

If you want to know how many heads and tails you got, pass the `-v` flag.

```shell
./ct3 -v
```

Sample output:

```
heads: 1546
tails: 1454
winner: heads
```

## Running as an HTTP server

**cointoss3000** can be run as an HTTP server. It listens for `GET /`
requests on a customisable port. To run it as a server, use the
`-s <port>` flag.  For example, to run it on port 8080,

```shell
./ct3 -s 8080
```

You can now `curl` it:

```shell
$ curl localhost:8080
tails
```

> Note that other flags are not allowed when using the `-s` flag.
> Use the query parameters instead.

### Query parameters

#### Verbose

You can set verbose output with the query parameter `v`. Example:

```shell
$ curl 'localhost:8080?v'
heads: 1459
tails: 1541
winner: tails
```

#### Number of tosses

You can set the number of tosses with the `t=` parameter. Example:

```shell
$ curl 'localhost:8080?v&t=3'
heads: 0
tails: 3
winner: tails
```

> Note:
> 1. If the verbose flag is not set, only 1 toss is performed to save
>    production CPU cycles.
> 2. The maximum number of tosses is 10,000 to prevent **cointoss3000**
>    from getting slammed in production.
> 3. Rate limiting is left as an exercise for the reader.

## Licence

BSD 3-Clause. See `LICENCE` for full text.
