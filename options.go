package gptj

type PredictOptions struct {
	RepeatPenalty, ContextSize, RepeatLastN, Tokens, TopK, Batch int
	TopP, Temperature, ContextErase                              float64
}

type PredictOption func(p *PredictOptions)

var DefaultOptions PredictOptions = PredictOptions{
	Tokens:       200,
	TopK:         40,
	TopP:         0.90,
	Temperature:  0.96,
	Batch:        9,
	ContextErase: 0.75,
	ContextSize:  1024,
}

// SetTokens sets the number of tokens to generate.
func SetTokens(tokens int) PredictOption {
	return func(p *PredictOptions) {
		p.Tokens = tokens
	}
}

// SetTopK sets the value for top-K sampling.
func SetTopK(topk int) PredictOption {
	return func(p *PredictOptions) {
		p.TopK = topk
	}
}

// SetTopP sets the value for nucleus sampling.
func SetTopP(topp float64) PredictOption {
	return func(p *PredictOptions) {
		p.TopP = topp
	}
}

// SetContextErase sets the context erase %.
func SetContextErase(ce float64) PredictOption {
	return func(p *PredictOptions) {
		p.ContextErase = ce
	}
}

// SetTemperature sets the temperature value for text generation.
func SetTemperature(temp float64) PredictOption {
	return func(p *PredictOptions) {
		p.Temperature = temp
	}
}

// SetBatch sets the batch size.
func SetBatch(size int) PredictOption {
	return func(p *PredictOptions) {
		p.Batch = size
	}
}

// Create a new PredictOptions object with the given options.
func NewPredictOptions(opts ...PredictOption) PredictOptions {
	p := DefaultOptions
	for _, opt := range opts {
		opt(&p)
	}
	return p
}

var DefaultModelOptions ModelOptions = ModelOptions{
	Threads: 4,
}

type ModelOptions struct {
	Threads int
}
type ModelOption func(p *ModelOptions)

// SetThreads sets the number of threads to use for text generation.
func SetThreads(c int) ModelOption {
	return func(p *ModelOptions) {
		p.Threads = c
	}
}

// Create a new PredictOptions object with the given options.
func NewModelOptions(opts ...ModelOption) ModelOptions {
	p := DefaultModelOptions
	for _, opt := range opts {
		opt(&p)
	}
	return p
}
