#include <stdio.h>

int main() {
  int mid_mid = 020;
  int forward_up = 0100;
  int forward_mid = 0200;
  int forward_bot = 0400;
  printf("mid_mid    =%03o %3d\n", mid_mid, mid_mid);
  printf("forward_up =%03o %3d\n", forward_up, forward_up);
  printf("forward_mid=%03o %3d\n", forward_mid, forward_mid);
  printf("forward_bot=%03o %3d\n", forward_bot, forward_bot);
  return 0;
}
