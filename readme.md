# cointoss3000

## Description

You need to toss a coin. But why settle for the best of 3? **cointoss3000** gives you the **best of 3000**.

## Dependencies

* make
* C compiler

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

Use `-n <number of tosses>`. For example, if you want the best of 3,

```shell
./ct3 -n 3
```

#### Extreme customisation

Of course, as the proud owner of cointoss3000, you won't want to limit yourself to a mere _three_ coin tosses.

cointoss3000 is capable of tossing a coin up to (at least) 2^64 - 1 times.
(Consult your C compiler's manual for the maximum number of coin tosses on your platform.)
**This will take a long time.**

```shell
./ct3 -n 18446744073709551615
```

#### Overflows not allowed

Overflowing the `unsigned long long` is not allowed. Sorry, Matt Parker.

```shell
# not allowed
./ct3 -n 18446744073709551616
```

#### Overflows allowed

You can do it backwards though.

```shell
./ct3 -n -1
```

is equivalent to

```shell
./ct3 -n 18446744073709551615
```

Matt Parker, you are welcome.

### Verbose

If you want to know how many heads and tails you got, pass the `-v` flag.

```shell
./ct3 -v
```

Sample output:

```shell
heads: 1546
tails: 1454
winner: heads
```

## Licence

BSD 3-Clause. See `licence` for full text.
