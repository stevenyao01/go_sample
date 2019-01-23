#include "worker.h"

int call_in(int(*worker_in)(worker_t*, char*, int), worker_t *this, char *data, int len)
{
	return worker_in(this, data, len);
}
int call_process(int(*worker_process)(worker_t*, const char*, int, char*, int), worker_t *this, const char *in, int len_in, char *out, int len_out)
{
	return worker_process(this, in, len_in, out, len_out);
}
int call_out(int(*worker_out)(worker_t*, const char*, int), worker_t *this, const char *data, int len)
{
	return worker_out(this, data, len);
}

extern GetLocalConfig(char *, int);
extern int Run(worker_t *);

void get_config(char *buf, int len)
{
    GetLocalConfig(buf, len);
}
int run(worker_t *worker)
{
    return Run(worker);
}