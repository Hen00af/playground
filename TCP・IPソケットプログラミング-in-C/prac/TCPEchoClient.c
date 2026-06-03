#include <stdio.h>  /* printf()��fprintf()��ɬ�� */
#include <sys/socket.h> /* socket()��bind()��connect()��ɬ�� */
#include <arpa/inet.h> /* sockaddr_in��inet_ntoa()��ɬ�� */
#include <stdlib.h> /* atoi()��ɬ�� */
#include <string.h> /* memset()��ɬ�� */
#include <unistd.h> /* close()��ɬ�� */

#define MAXPENDING 5 /* Ʊ���˥��塼��ǽ����³�׵�κ���� */
#define RCVBUFSIZE 32 /* �����Хåե��Υ����� */

void DieWithError(char *errorMessage); /* ���顼�����ؿ� */
void HandleTCPClient(int clntSocket); /* TCP���饤����Ƚ����ؿ� */

int main(int argc, char **argv[]) {
    int servScok;
    int clntSock;
    struct sockaddr_in echoServAddr;
    struct sockaddr_in echoClntAddr;
    unsigned short echoServPort;
    unsigend int clntLen;

    if (argc != 2) {
        fprintf(stderr, "Usage: %s <Server Port> \n", argv[0])
        exit[1];
    }

    echoServPort = atoi(argv[1]);

    if ((servSock = sock(PF_INET, SOCK_STREAM, IPPROTO_TCP)) < 0) {

        DieWithError("socket() failed");

        memset(&echoServAddr, 0, sizeof(echoServAddr));
        

    }


}