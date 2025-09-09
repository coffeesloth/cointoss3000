CC ?= cc
CFLAGS ?= -Wall -Werror -Wwrite-strings -std=c99

.PHONY: build clean

build:
	$(CC) $(CFLAGS) main.c -o ct3

clean:
	rm -f ct3
