/* cointoss3000 - coin toss simulator. BSD 3-clause.
 * See licence for full text.
 */

#define _POSIX_C_SOURCE 200809L

#include <errno.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include <unistd.h>

int invalidArgs(int argc, char **argv) {
	const char *progname = (argc > 0 && argv != NULL) ? argv[0] : "ct3";
	fprintf(stderr, "Usage: %s [-v] [-n <tosses>]\n", progname);
	return 1;
}

int main(int argc, char **argv) {
	bool verbose = false;
	unsigned long long tosses = 3000;

	int opt;
	while ((opt = getopt(argc, argv, "vn:")) != -1) {
		switch (opt) {
		case 'v':
			verbose = true;
			break;
		case 'n': {
			errno = 0;
			char *endptr;
			tosses = strtoull(optarg, &endptr, 10);
			if (errno != 0 || *endptr != '\0') {
				return invalidArgs(argc, argv);
			}
			break;
		}
		default:
			return invalidArgs(argc, argv);
		}
	}

	if (optind < argc) {
		return invalidArgs(argc, argv);
	}

	srand(time(NULL));

	unsigned long long results[2] = {0, 0};
	for (unsigned long long i = 0; i < tosses; i++) {
		results[rand() % 2]++;
	}

	const char *winner;
	if (results[0] > results[1]) {
		winner = "heads";
	} else if (results[1] > results[0]) {
		winner = "tails";
	} else {
		winner = "draw";
	}

	if (verbose) {
		printf("heads: %llu\ntails: %llu\nwinner: ", results[0], results[1]);
	}
	printf("%s\n", winner);

	return 0;
}
