package main

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
	"github.com/suzuito/sandbox-go/cmd/002/pb"
	"google.golang.org/protobuf/proto"
)

func main() {
	env, err := cel.NewEnv(
		cel.Types(&pb.Data{}),
		cel.Declarations(
			decls.NewVar(
				"data",
				decls.NewObjectType("cel_sandbox.Data"),
			),
		),
	)
	if err != nil {
		glog.Exit(err)
	}
	ast, iss := env.Parse(`data.org == 'Hoge campany' && ('/foo' in data.permissions['storage.viewer'].buckets)`)
	if iss.Err() != nil {
		glog.Exit(iss.Err())
	}
	checked, iss := env.Check(ast)
	if iss.Err() != nil {
		glog.Exit(iss.Err())
	}
	if !proto.Equal(checked.ResultType(), decls.Bool) {
		glog.Exitf(
			"Got %v, wanted %v result type",
			checked.ResultType(), decls.String,
		)
	}
	program, err := env.Program(checked)
	if err != nil {
		glog.Exitf("program error: %v", err)
	}

	result, _, err := program.Eval(map[string]interface{}{
		"data": &pb.Data{
			Org: "Hoge campany",
			Permissions: map[string]*pb.Permission{
				"storage.viewer": {
					Buckets: []string{"/foo", "/bar"},
				},
			},
		},
	})
	if err != nil {
		glog.Exitf("eval error: %v", err)
	}
	fmt.Println(result)
}
