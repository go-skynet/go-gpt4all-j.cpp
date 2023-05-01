#ifdef __cplusplus
extern "C" {
#endif

#include <stdbool.h>

void* binding_load_gptj_model(const char *fname, int n_threads);

void binding_gptj_free_model(void *state_ptr);

extern unsigned char bindingTokenCallback(void *, char *);

void binding_model_prompt(const char *prompt, void* model, char* result, int repeat_last_n, float repeat_penalty, int n_ctx, int tokens, int top_k,
                            float top_p, float temp, int n_batch,float ctx_erase);

#ifdef __cplusplus
}
#endif