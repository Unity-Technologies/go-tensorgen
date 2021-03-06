// Copyright 2020 Unity Technologies SF
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// go-tensorgen generates Go structures and Tensorflow inference execution code from SavedModel file.
package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"github.com/golang/protobuf/proto"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	corepb "github.com/tensorflow/tensorflow/tensorflow/go/core/core_protos_go_proto"
	shapepb "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_shape_go_proto"
	typespb "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/types_go_proto"
)

const tmpl = `// Copyright 2020 Unity Technologies SF
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// DO NOT EDIT: generated by go-tensorgen
//
package {{.PackageName}}

import (
	"context"
	"errors"
	"fmt"

	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)
	
const (
	{{.SignatureName|Lower}}RequestPropsCount = {{ len .Inputs }}
	{{.SignatureName|Lower}}ResponsePropsCount = {{ len .Outputs }}
)

type {{.SignatureName|Title}}Request struct {
	{{range $desc := .Inputs -}}
	// {{$desc.Name}} {{$desc.Type}} Shape:{{$desc.Shape|ShapeToString}}
	{{$desc|PropName}} {{$desc|GoType}}
	{{end}}
}

type {{.SignatureName|Title}}Response struct {
	{{range $desc := .Outputs -}}
	// {{$desc.Name}} {{$desc.Type}} Shape:{{$desc.Shape|ShapeToString}}
	{{$desc|PropName}} {{$desc|GoType}}
	{{end}}
}

func (r *{{.SignatureName|Title}}Response) FromTensors(tensors []*tf.Tensor) error {
	{{range $i, $desc := .Outputs -}}
	{
		val, ok := tensors[{{$i}}].Value().({{$desc|GoType}})
		if !ok {
			return errors.New("failed to cast tensor for \"{{$desc|PropName}}\" to {{$desc|GoType}}")
		}

		r.{{$desc|PropName}} = val
	}
	{{end}}

	return nil
}

type {{.SignatureName|Title}}Runner struct {
	graph *tf.Graph
	session *tf.Session

	inputPlaceholders  [{{.SignatureName|Lower}}RequestPropsCount]tf.Output
	outputPlaceholders [{{.SignatureName|Lower}}ResponsePropsCount]tf.Output
}

func New{{.SignatureName|Title}}Runner(graph *tf.Graph, session *tf.Session) *{{.SignatureName|Title}}Runner {
	return &{{.SignatureName|Title}}Runner{
		graph: graph,
		session: session,
	}
}

func (r *{{.SignatureName|Title}}Runner) LoadOperations() error {
	inputDescriptions := [{{.SignatureName|Lower}}RequestPropsCount]struct {
		name  string
		index int
	}{
		{{range $desc := .Inputs -}}
		// {{$desc.Name}} {{$desc.Type}} Shape:{{$desc.Shape|ShapeToString}}
		{name: "{{$desc.OpName}}", index: {{$desc.OpIndex}}},
		{{end}}
	}

	for i, description := range inputDescriptions {
		op := r.graph.Operation(description.name)
		if op == nil {
			return fmt.Errorf("operation not found by name \"%s\"", description.name)
		}

		r.inputPlaceholders[i] = tf.Output{
			Op:    op,
			Index: description.index,
		}
	}

	outputDescriptions := [{{.SignatureName|Lower}}ResponsePropsCount]struct {
		name  string
		index int
	}{
		{{range $desc := .Outputs -}}
		{name: "{{$desc.OpName}}", index: {{$desc.OpIndex}}},
		{{end}}
	}

	for i, description := range outputDescriptions {
		op := r.graph.Operation(description.name)
		if op == nil {
			return fmt.Errorf("operation not found by name \"%s\"", description.name)
		}

		r.outputPlaceholders[i] = tf.Output{
			Op:    op,
			Index: description.index,
		}
	}

	return nil
}

func (r *{{.SignatureName|Title}}Runner) feedsFromRequest(req {{.SignatureName|Title}}Request) (feeds map[tf.Output]*tf.Tensor, err error) {
	feeds = make(map[tf.Output]*tf.Tensor, {{.SignatureName|Lower}}RequestPropsCount)

	{{range $i, $desc := .Inputs -}}
	feeds[r.inputPlaceholders[{{$i}}]], err = tf.NewTensor(req.{{$desc|PropName}})
	if err != nil {
		return nil, fmt.Errorf("failed to create tensor for \"{{$desc|PropName}}\": %w", err)
	}
	{{end}}

	return feeds, nil
}

func (r *{{.SignatureName|Title}}Runner) feedsFromTensors(tensors []*tf.Tensor) (feeds map[tf.Output]*tf.Tensor, err error) {
	if len(tensors) != {{.SignatureName|Lower}}RequestPropsCount {
		return nil, fmt.Errorf("tensors have %d elems, %d required", len(tensors), {{.SignatureName|Lower}}RequestPropsCount)
	}

	feeds = make(map[tf.Output]*tf.Tensor, {{.SignatureName|Lower}}RequestPropsCount)

	for i := 0; i < {{.SignatureName|Lower}}RequestPropsCount; i++ {
		feeds[r.inputPlaceholders[i]] = tensors[i]
	}

	return feeds, nil
}

// Run converts req to tensors, runs inference and casts results to the {{.SignatureName|Title}}Response struct.
func (r *{{.SignatureName|Title}}Runner) Run(ctx context.Context, req {{.SignatureName|Title}}Request) ({{.SignatureName|Title}}Response, error) {
	feeds, err := r.feedsFromRequest(req)
	if err != nil {
		return {{.SignatureName|Title}}Response{}, err
	}

	fetches := r.outputPlaceholders[:]

	output, err := r.session.Run(feeds, fetches, nil)
	if err != nil {
		return {{.SignatureName|Title}}Response{}, fmt.Errorf("failed to run session: %w", err)
	}

	var resp {{.SignatureName|Title}}Response

	err = resp.FromTensors(output)
	if err != nil {
		return {{.SignatureName|Title}}Response{}, err
	}

	return resp, nil
}

// RunTensors converts input tensors to feeds (tf.Output => tf.Tensors dict),
// runs inference and returns output tensors.
//
// The order of input tensors must be the same as the order of properties in the {{.SignatureName|Title}}Request struct.
func (r *{{.SignatureName|Title}}Runner) RunTensors(ctx context.Context, tensors []*tf.Tensor) ([]*tf.Tensor, error) {
	feeds, err := r.feedsFromTensors(tensors)
	if err != nil {
		return nil, err
	}

	fetches := r.outputPlaceholders[:]

	output, err := r.session.Run(feeds, fetches, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to run session: %w", err)
	}

	return output, nil
}

// RunForTensors runs inference for req and returns raw results as tensors.
//
// The order of items in the result array is the same as the order of properties in the {{.SignatureName|Title}}Response struct.
func (r *{{.SignatureName|Title}}Runner) RunForTensors(ctx context.Context, req {{.SignatureName|Title}}Request) ([]*tf.Tensor, error) {
	feeds, err := r.feedsFromRequest(req)
	if err != nil {
		return nil, err
	}

	fetches := r.outputPlaceholders[:]

	output, err := r.session.Run(feeds, fetches, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to run session: %w", err)
	}

	return output, nil
}

// RunWithTensors converts input tensors to feeds (tf.Output => tf.Tensors dict),
// runs inference and casts results to the {{.SignatureName|Title}}Response struct.
//
// The order of input tensors must be the same as the order of properties in the {{.SignatureName|Title}}Request struct.
func (r *{{.SignatureName|Title}}Runner) RunWithTensors(ctx context.Context, tensors []*tf.Tensor) ({{.SignatureName|Title}}Response, error) {
	feeds, err := r.feedsFromTensors(tensors)
	if err != nil {
		return {{.SignatureName|Title}}Response{}, err
	}

	fetches := r.outputPlaceholders[:]

	output, err := r.session.Run(feeds, fetches, nil)
	if err != nil {
		return {{.SignatureName|Title}}Response{}, fmt.Errorf("failed to run session: %w", err)
	}

	var resp {{.SignatureName|Title}}Response

	err = resp.FromTensors(output)
	if err != nil {
		return {{.SignatureName|Title}}Response{}, err
	}

	return resp, nil
}

`

type settings struct {
	Help      bool
	ModelDir  string
	OutputDir string
	Package   string
}

func loadSettings() settings {
	var cfg settings

	flag.StringVar(&cfg.ModelDir, "model-dir", "./model", "Directory on disk where model is located")
	flag.StringVar(&cfg.OutputDir, "output-dir", "./runner", "Directory on disk where result file will be created")
	flag.StringVar(&cfg.Package, "package", "", "Name of the package with generated code (if not defined, then equal output dir)")

	flag.BoolVar(&cfg.Help, "help", false, "Output flags usage")

	flag.Parse()

	return cfg
}

func main() {
	cfg := loadSettings()

	if cfg.Help {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()

		os.Exit(0)
	}

	statusCode := run(cfg)

	os.Exit(statusCode)
}

func run(cfg settings) int {
	filepath := path.Join(cfg.ModelDir, "saved_model.pb")

	modelStruct, err := loadModelStructure(filepath)
	if err != nil {
		fmt.Println(err)

		return 1
	}

	model, err := loadSavedModel(cfg.ModelDir)
	if err != nil {
		fmt.Println(err)

		return 1
	}

	t, err := prepareTemplate()
	if err != nil {
		fmt.Println(err)

		return 1
	}

	tags := []string{"serve"}

	metagraph, err := findMetaGraph(modelStruct, tags)
	if err != nil {
		fmt.Println(err)

		return 1
	}

	for sigName, sigDef := range metagraph.SignatureDef {
		if sigDef == nil {
			continue
		}

		// Signature can't be used for inference
		if sigName == "__saved_model_init_op" {
			continue
		}

		fmt.Println("start generating for " + sigName)

		err = os.MkdirAll(cfg.OutputDir, 0755)
		if err != nil {
			fmt.Println(err)

			return 1
		}

		filepath := path.Join(cfg.OutputDir, strings.ToLower(sigName)+".go")

		packageName := cfg.Package
		if packageName == "" {
			_, packageName = path.Split(cfg.OutputDir)
		}

		err = generateFile(filepath, t, sigName, packageName, sigDef, model.Graph)
		if err != nil {
			fmt.Println(err)

			return 1
		}

		fmt.Println("file for " + sigName + " successfuly generated")
	}

	return 0
}

func generateFile(filepath string, t *template.Template, sigName, packageName string, sigDef *corepb.SignatureDef, graph *tf.Graph) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	err = generate(f, t, sigName, packageName, sigDef, graph)
	if err != nil {
		return err
	}

	gofmt(filepath)

	return nil
}

func prepareTemplate() (*template.Template, error) {
	baseT := template.New("gen").Funcs(template.FuncMap{
		"Lower": strings.ToLower,
		"Title": strings.Title,
		"GoType": func(desc description) string {
			return typeStringFromShape(desc.Type, desc.Shape)
		},
		"PropName": func(desc description) string {
			return CamelCase(desc.Name)
		},
		"ShapeToString": shapeToString,
	})

	t, err := baseT.Parse(removeNon(tmpl))
	if err != nil {
		return nil, err
	}

	return t, nil
}

func generate(wr io.Writer, t *template.Template, sigName, packageName string, sigDef *corepb.SignatureDef, graph *tf.Graph) error {
	inputs, err := readDescriptions(graph, sigDef.Inputs)
	if err != nil {
		return err
	}

	outputs, err := readDescriptions(graph, sigDef.Outputs)
	if err != nil {
		return err
	}

	err = t.Execute(wr, struct {
		SignatureName string
		PackageName   string
		Inputs        []description
		Outputs       []description
	}{
		SignatureName: sigName,
		PackageName:   packageName,
		Inputs:        inputs,
		Outputs:       outputs,
	})
	if err != nil {
		return err
	}

	return nil
}

func gofmt(path string) {
	cmd := exec.Command("goimports", "-w", path)

	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("gofmt error: %s\n%s", err, b)
	}
}

// removeNon removes \nn\n from string
func removeNon(src string) string {
	return strings.Replace(src, "\\nn\n", "", -1)
}

var tftypes = map[typespb.DataType]string{
	typespb.DataType_DT_FLOAT:      "float32",
	typespb.DataType_DT_DOUBLE:     "float64",
	typespb.DataType_DT_INT32:      "int32",
	typespb.DataType_DT_UINT8:      "uint8",
	typespb.DataType_DT_INT16:      "int16",
	typespb.DataType_DT_INT8:       "int8",
	typespb.DataType_DT_STRING:     "string",
	typespb.DataType_DT_COMPLEX64:  "complex64",
	typespb.DataType_DT_INT64:      "int64",
	typespb.DataType_DT_BOOL:       "bool",
	typespb.DataType_DT_UINT16:     "uint16",
	typespb.DataType_DT_COMPLEX128: "complex128",
	typespb.DataType_DT_UINT32:     "uint32",
	typespb.DataType_DT_UINT64:     "uint64",
}

func typeStringFromShape(typ typespb.DataType, shape *shapepb.TensorShapeProto) string {
	if shape == nil {
		return tftypes[typ]
	}

	if shape.UnknownRank {
		return "interface{}"
	}

	depth := len(shape.Dim)

	return strings.Repeat("[]", depth) + tftypes[typ]
}

func shapeToString(shape *shapepb.TensorShapeProto) string {
	if shape == nil {
		return "SHAPE_NOT_DEFINED"
	}

	if shape.UnknownRank {
		return "UNKNOWN_NUMBER_OF_DIMENSIONS"
	}

	sizes := make([]string, len(shape.Dim))
	for i, dim := range shape.Dim {
		sizes[i] = strconv.Itoa(int(dim.Size))
	}

	return "(" + strings.Join(sizes, ",") + ")"
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Strings(a)
	sort.Strings(b)

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func findMetaGraph(m *corepb.SavedModel, tags []string) (*corepb.MetaGraphDef, error) {
	if len(m.MetaGraphs) == 0 {
		return nil, errors.New("model doesn't contain metagraphs")
	}

	for _, graph := range m.MetaGraphs {
		info := graph.GetMetaInfoDef()
		graphTags := info.GetTags()

		if !equal(graphTags, tags) {
			continue
		}

		return graph, nil
	}

	return nil, fmt.Errorf("metagraph not found by tags \"%v\"", tags)
}

func loadModelStructure(filepath string) (*corepb.SavedModel, error) {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return nil, fmt.Errorf("file \"%s\" doesn't exist", filepath)
	}

	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file \"%s\": %w", filepath, err)
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from file \"%s\": %w", filepath, err)
	}

	var model corepb.SavedModel

	err = proto.Unmarshal(b, &model)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal savedmodel \"%s\": %w", filepath, err)
	}

	return &model, nil
}

func loadSavedModel(dir string) (*tf.SavedModel, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, fmt.Errorf("directory \"%s\" doesn't exist", dir)
	}

	metaGraph := "serve"

	model, err := tf.LoadSavedModel(dir, []string{metaGraph}, &tf.SessionOptions{})
	if err != nil {
		return nil, err
	}

	return model, nil
}

type description struct {
	Name    string
	Type    typespb.DataType
	Shape   *shapepb.TensorShapeProto
	OpName  string
	OpIndex int
}

type ByOpIndexAndName []description

func (a ByOpIndexAndName) Len() int      { return len(a) }
func (a ByOpIndexAndName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByOpIndexAndName) Less(i, j int) bool {
	if a[i].OpIndex != a[j].OpIndex {
		return a[i].OpIndex < a[j].OpIndex
	}

	return a[i].OpName < a[j].OpName
}

func readDescriptions(graph *tf.Graph, dict map[string]*corepb.TensorInfo) ([]description, error) {
	res := make([]description, 0, len(dict))

	for inputName, tensorInfo := range dict {
		encoding, ok := tensorInfo.Encoding.(*corepb.TensorInfo_Name)
		if !ok {
			return nil, fmt.Errorf("it's not possible to convert encoding of \"%s\" tensor to *corepb.TensorInfo_Name", inputName)
		}

		parts := strings.Split(encoding.Name, ":")

		const expectedPartsCount = 2
		if len(parts) != expectedPartsCount {
			return nil, fmt.Errorf("unexpected format of tensor name \"%s\"", encoding.Name)
		}

		opname := parts[0]

		index, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("tensor operation index \"%s\" can't be converted to integer", parts[1])
		}

		operation := graph.Operation(opname)
		if operation == nil {
			return nil, fmt.Errorf("operation \"%s\" wasn't found in graph", opname)
		}

		n := description{
			Name:    inputName,
			Type:    tensorInfo.GetDtype(),
			Shape:   tensorInfo.GetTensorShape(),
			OpName:  opname,
			OpIndex: index,
		}

		res = append(res, n)
	}

	sort.Sort(ByOpIndexAndName(res))

	return res, nil
}

// @see: https://github.com/golang/lint/blob/738671d3881b9731cc63024d5d88cf28db875626/lint.go#L767
var commonInitialisms = map[string]bool{
	"ACL":   true,
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SQL":   true,
	"SSH":   true,
	"TCP":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XMPP":  true,
	"XSRF":  true,
	"XSS":   true,
}

// CamelCase converts s to CamelCase and takes common initialisms (e.g. HTTP, ID) into account.
func CamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i, part := range parts {
		if commonInitialisms[strings.ToUpper(part)] {
			parts[i] = strings.ToUpper(part)

			continue
		}

		parts[i] = strings.Title(part)
	}

	return strings.Join(parts, "")
}
