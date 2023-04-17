package gptj

// #cgo CXXFLAGS: -I./gpt4all-j/ggml/include -I./gpt4all-j/ggml/include/ggml -I./gpt4all-j -I./gpt4all-j/ggml/examples -I./gpt4all-j/ggml
// #cgo LDFLAGS: -L./ -l:libbinding.a -lm -lstdc++
// #cgo darwin LDFLAGS: -framework Accelerate
// #include "binding.h"
import "C"
import (
	"fmt"
	"strings"
	"unsafe"
)

type GPTJ struct {
	state unsafe.Pointer
}

func New(model string, opts ...ModelOption) (*GPTJ, error) {
	mo := NewModelOptions(opts...)
	state := C.gptj_allocate_state()
	modelPath := C.CString(model)
	result := C.gptj_bootstrap(modelPath, state, C.int(mo.ContextSize))
	if result != 0 {
		return nil, fmt.Errorf("failed loading model")
	}

	return &GPTJ{state: state}, nil
}

func (l *GPTJ) Predict(text string, opts ...PredictOption) (string, error) {

	po := NewPredictOptions(opts...)

	input := C.CString(text)
	if po.Tokens == 0 {
		po.Tokens = 99999999
	}
	out := make([]byte, po.Tokens)

	params := C.gptj_allocate_params(input, C.int(po.Seed), C.int(po.Threads), C.int(po.Tokens), C.int(po.TopK),
		C.float(po.TopP), C.float(po.Temperature), C.int(po.Batch))
	ret := C.gptj_predict(params, l.state, (*C.char)(unsafe.Pointer(&out[0])))
	if ret != 0 {
		return "", fmt.Errorf("inference failed")
	}
	res := C.GoString((*C.char)(unsafe.Pointer(&out[0])))

	res = strings.TrimPrefix(res, " ")
	res = strings.TrimPrefix(res, text)
	res = strings.TrimPrefix(res, "\n")

	C.gptj_free_params(params)

	return res, nil
}
func (l *GPTJ) Free() {
	C.gptj_free_model(l.state)
}
