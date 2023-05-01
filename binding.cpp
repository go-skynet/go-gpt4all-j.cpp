#include "ggml.h"
#include "gpt4all-j/llmodel/llmodel_c.h"
#include "gpt4all-j/llmodel/llmodel.h"
#include "gpt4all-j/llmodel/llama.cpp/llama.h"

#include "gpt4all-j/llmodel/llamamodel.h"
#include "gpt4all-j/llmodel/gptj.h"

#include "binding.h"
#include "llama.h"

#include <cassert>
#include <cmath>
#include <cstdio>
#include <cstring>
#include <fstream>
#include <map>
#include <string>
#include <vector>
#include <iostream>
#include <unistd.h>

void* load_gptj_model(const char *fname, int n_threads) {
    // load the model
    auto gptj = llmodel_gptj_create();

    llmodel_setThreadCount(gptj,  n_threads);
    if (!llmodel_loadModel(gptj, fname)) {
        return nullptr;
    }

    return gptj;
}

std::string res ="";
void * mm;

void binding_model_prompt( const char *prompt, void *m,
                    void * c, char* result)
{
    llmodel_model* model = (llmodel_model*) m;
    llmodel_prompt_context* ctx = (llmodel_prompt_context*)c;

   // std::string res = "";
 
    auto lambda_prompt = [](int token_id, const char *promptchars)  {
	        return true;
    };

    mm=model;
    res="";

    auto lambda_response = [](int token_id, const char *responsechars) {
        res.append((char*)responsechars);
        return !!tokenCallback(mm, (char*)responsechars);
	};
	
	auto lambda_recalculate = [](bool is_recalculating) {
	        // You can handle recalculation requests here if needed
	    return is_recalculating;
	};

    llmodel_prompt(model, prompt,
                        lambda_prompt,
                        lambda_response,
                    lambda_recalculate,
                    ctx);

    strcpy(result, res.c_str()); 
}

void llama_free_model(void *state_ptr) {
    llmodel_model* ctx = (llmodel_model*) state_ptr;
    llmodel_llama_destroy(ctx);
}

void gptj_free_ctx(void *state_ptr) {
    llmodel_prompt_context* ctx = (llmodel_prompt_context*) state_ptr;
    llmodel_llama_destroy(ctx);
}


void* gptj_new_context( int repeat_last_n, int repeat_penalty, int n_ctx, int tokens, int top_k,
                            float top_p, float temp, int n_batch,float ctx_erase) {
    llmodel_prompt_context* params = new llmodel_prompt_context;
    params->n_predict = tokens;
    params->repeat_last_n = repeat_last_n;
    params->repeat_penalty = repeat_penalty;
    params->n_ctx = n_ctx;
    params->top_k = top_k;
    params->context_erase = ctx_erase;
    params->top_p = top_p;
    params->temp = temp;
    params->n_batch = n_batch;    
    return params;
}

