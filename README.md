# [![Go Reference](https://pkg.go.dev/badge/github.com/go-skynet/go-gpt4all-j.cpp.svg)](https://pkg.go.dev/github.com/go-skynet/go-gpt4all-j.cpp) go-gpt4all-j.cpp

This is a go binding for [GPT4ALL-J](https://github.com/nomic-ai/gpt4all).

This was originally developed by [mudler](https://github.com/mudler) for the [LocalAI](https://github.com/go-skynet/LocalAI) project.

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
