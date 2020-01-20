#ifndef __WORKER_H__
#define __WORKER_H__

#ifdef __cplusplus
extern "C" {
#endif

// worker type
#define SOURCE      0
#define PROCESSOR   1
#define TARGET      2

typedef struct worker worker_t;

typedef struct worker {
    void *user;

    int worker_type;

    // return value: on success, the number of bytes read is returned,
    // zero indicates end of in,
    // on error, -1 is returned.
    int(*worker_in)(worker_t *this, char *data, int len);

    // return value: on success, the number of bytes out is returned,
    // zero indicates end of in,
    // on error, -1 is returned.
    int(*worker_process)(worker_t *this, const char *in, int len_in, char *out, int len_out);

    // return value: on success, the number of bytes out is returned,
    // zero indicates end of in,
    // on error, -1 is returned.
    int(*worker_out)(worker_t *this, const char *data, int len);
}worker_t;

extern void get_config(char *buf, int len);
extern int run(worker_t *worker);

#ifdef __cplusplus
}
#endif

#endif