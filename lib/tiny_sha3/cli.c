#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "sha3.h"

void print_help() {
    printf("usage sha3 input\n");
}

int main(int argc, char **argv) {
    if (argc != 2) {
        print_help();
        return 0;
    }

    char* input = argv[1];
    int hash_length = 64; // for sha3-512
    unsigned char* hash = malloc(hash_length);

    if (hash == NULL) {
        fprintf(stderr, "Memory allocation failed\n");
        return 1;
    }

    sha3(input, strlen(input), hash, hash_length);

    for (int i = 0; i < hash_length; i++)
        printf("%02x", hash[i]);
    printf("\n");

    free(hash);
    return 0;
}
