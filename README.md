# [![Go Reference](https://pkg.go.dev/badge/github.com/go-skynet/go-gpt4all-j.cpp.svg)](https://pkg.go.dev/github.com/go-skynet/go-gpt4all-j.cpp) go-gpt4all-j.cpp

This is a go binding/implementation compatible with [GPT4ALL-J](https://github.com/nomic-ai/gpt4all).

This port is high level, avoiding token-by-token inferencing in order to be faster. Check out [this](https://about.sourcegraph.com/blog/go/gophercon-2018-adventures-in-cgo-performance) and [this](https://www.cockroachlabs.com/blog/the-cost-and-complexity-of-cgo/) write-ups which summarize the impact of a low-level interface which calls C functions from Go.

If you are looking for an high-level OpenAI compatible API, check out [here](https://github.com/go-skynet/llama-cli).

## Usage

Note: This repository uses git submodules to keep track of [GPT4ALL-J](https://github.com/nomic-ai/gpt4all).

Clone the repository locally:

```bash
git clone --recurse-submodules https://github.com/go-skynet/go-gpt4all-j.cpp
```

To build the bindings locally, run:

```
cd go-gpt4all-j.cpp
make libgptj.a
```

Now you can run the example with:

```
LIBRARY_PATH=$PWD C_INCLUDE_PATH=$PWD go run ./examples -m "/model/path/here" -t 14
```

Enjoy!

The documentation is available [here](https://pkg.go.dev/github.com/go-skynet/go-gpt4all-j.cpp) and the full example code is [here](https://github.com/go-skynet/go-llama.cpp/blob/master/examples/main.go).

## License

MIT
