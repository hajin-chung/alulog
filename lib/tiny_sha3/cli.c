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

    int i;
    char* input = argv[1];
    char* hash = malloc(sizeof(char) * 512);
    sha3(input, strlen(input), hash, 512);
    for (int i = 0; i < sizeof(hash); i++)
        printf("%x", hash[i]&0xff);
    printf("\n");

    free(hash);
}
