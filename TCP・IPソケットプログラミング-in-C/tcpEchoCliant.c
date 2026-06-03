#include <stdio.h>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

#define RCVBUFSIZE 32

voidDieWithError(char *errorMessage);

int main(int argc, char **argv) {
    int socket;
    struct sockaddr_in echoServerAddr;
}