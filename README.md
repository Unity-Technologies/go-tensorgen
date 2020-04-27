[![Go Report Card](https://goreportcard.com/badge/github.com/Unity-Technologies/go-tensorgen)](https://goreportcard.com/report/github.com/Unity-Technologies/go-tensorgen)

# Tensorgen

Lightweight tool that generates Go structures and [Tensorflow](https://www.tensorflow.org/) inference execution code from [SavedModel](https://github.com/tensorflow/tensorflow/blob/master/tensorflow/python/saved_model/README.md).
The tool is designed to save the developer from the headache of tensors construction when working with Tensorflow from Go.

# Usage

There are two steps that you need to perform to use this tool:
* Generate Go code from `SavedModel`.
* Perform model execution using generated code.

## Code Generating

The simplest way is to use Docker image that has `libtensorflow` dependencies inside. But if you have `libtensorflow` locally — sky is the limit.

### With Docker

TBD

### Without Docker

To run Go code `libtensorflow` must be available as dynamic library. It means that `.so` and `.h` files, compatible with used version of [tensorflow/go](https://pkg.go.dev/github.com/tensorflow/tensorflow/tensorflow/go) package, must be provided. More info: [Install TensorFlow for C](https://www.tensorflow.org/install/lang_c)

Install tool with `go get`
```
go get -u github.com/Unity-Technologies/go-tensorgen
```

Exec `go-tensorgen` with defined SavedModel files directory and output dir, where generated code will be placed.
```
go-tensorgen -model-dir=./savedmodel -output-dir=./mymodel
```

## Inference execution

For each available [signature](https://www.tensorflow.org/tfx/serving/signature_defs) separate "Runner" will be generated. You can find example, how to use it, below.

```go
import (
	"github.com/user/project/mymodel"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

dir := "./savedmodel"
metaGraph := "serve"
model, err := tf.LoadSavedModel(dir, []string{metaGraph}, &tf.SessionOptions{})
if err != nil { ... }

runner := mymodel.NewSignatureRunner(model.Graph, model.Session)

err = runner.LoadOperations()
if err != nil { ... }

req := mymodel.SignatureRequest{
	Feature: []float32{0.5},
}

resp, err := runner.Run(ctx, req)
if err != nil { ... }

// resp - is ready to use go struct
```

# TODO:

* [ ] Prepare Docker image that includes `libtensorflow`.
* [ ] Explain how to use binary distributions of `libtensorflow` for local development

# Contributing

Pull requests are very much welcomed. Create your pull request, make sure a test or example is included that covers your change and
your commits represent coherent changes that include a reason for the change.

Use [golangci-lint](https://github.com/golangci/golangci-lint) to check code with linters:
```
golangci-lint run ./...
```
