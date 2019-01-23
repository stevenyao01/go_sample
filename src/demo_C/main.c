#include <stdio.h>
#include <string.h>
#include "worker.h"

int in(worker_t *this, char *data, int len)
{
    if (this == NULL || data == NULL || len <= 0)
        return -1;

    const char *p = "this is c worker on in";
    if (strlen(p) + 1 < len) {
        len = strlen(p) + 1;
    }
    strncpy(data, p, len);

    return len;
}

int process(worker_t *this, const char *in, int len_in, char *out, int len_out)
{
    if (this == NULL || in == NULL || len_in <= 0 || out == NULL || len_out <= 0)
        return -1;

    strcpy(out, in);
    strcat(out, " - on process");

    return strlen(out);
}

int out(worker_t *this, const char *data, int len)
{
    if (this == NULL || data == NULL || len <= 0)
            return -1;

    return printf("c on out: %s\n", data);
}

int _main(int argc, char* argv[])
{
    char config[1024] = { 0 };
    get_config(config, 1024);
    printf("config file name:%s\n", config);

    worker_t w;
    memset(&w, 0, sizeof(w));
    w.user = NULL;
    w.worker_type = SOURCE;
    w.worker_in = in;
    w.worker_process = process;
    w.worker_out = out;

    return run(&w);
}