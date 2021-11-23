package main

import (
	"io/ioutil"
	"os"

	"google.golang.org/protobuf/proto"
	plugin "google.golang.org/protobuf/types/pluginpb"
)

const defaultRequestFileName = "request.bin"

func main() {
	req := &plugin.CodeGeneratorRequest{}
	resp := &plugin.CodeGeneratorResponse{}

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	if err := proto.Unmarshal(data, req); err != nil {
		panic(err)
	}

	reqBytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}

	content := string(reqBytes)

	outFileName := req.GetParameter()
	if outFileName == "" {
		outFileName = defaultRequestFileName
	}

	file := &plugin.CodeGeneratorResponse_File{
		Name:    &outFileName,
		Content: &content,
	}
	resp.File = []*plugin.CodeGeneratorResponse_File{file}

	marshalled, err := proto.Marshal(resp)
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(marshalled)
}
