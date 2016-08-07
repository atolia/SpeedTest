#include<stdio.h>

int main()
{
    int x;
    double sum = 0;
    FILE *fin;
    fin = fopen("BIGDATA2", "r");

    while(fscanf(fin, "%d", &x) == 1)
    {
        sum += x;
    }

    printf("%f\n", sum);

    return 0;
}

