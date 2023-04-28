#ifdef __cplusplus
extern "C" {
#endif

#include <stdbool.h>

void* load_gptj_model(const char *fname, int n_threads);
void llama_free_model(void *state_ptr);

void* gptj_allocate_params(const char *prompt, int seed, int threads, int tokens,
                            int top_k, float top_p, float temp, int n_batch);
void gptj_free_ctx(void* params_ptr);
void gptj_free_model(void* state_ptr);
extern unsigned char tokenCallback(void *, char *);

void binding_model_prompt(const char *prompt, void* model,  void* state_pr, char* result);

void* gptj_new_context( int repeat_last_n, int repeat_penalty, int n_ctx, int tokens, int top_k,
                            float top_p, float temp, int n_batch) ;
#ifdef __cplusplus
}
#endif