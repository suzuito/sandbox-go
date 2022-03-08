package main

import (
	"encoding/json"
	"fmt"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
	"github.com/suzuito/sandbox-go/cmd/003/pb"
	"golang.org/x/xerrors"
	"google.golang.org/protobuf/proto"
)

func main() {
	// CELを用いた評価器を生成する
	filter, err := newCELEvaluator()
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
	// 評価器を用いてデータ評価する
	for _, data := range []map[string]interface{}{
		// 評価されるテストデータ
		{
			"data": &pb.Data{
				Org: "Hoge campany",
				Permissions: map[string]*pb.Permission{
					"storage.viewer": {
						Buckets: []string{"/foo", "/goo"},
					},
				},
			},
		},
		{
			"data": &pb.Data{
				Org: "Fuga campany",
				Permissions: map[string]*pb.Permission{
					"storage.viewer": {
						Buckets: []string{"/foo", "/goo"},
					},
				},
			},
		},
		{
			"data": &pb.Data{
				Org: "Hoge campany",
				Permissions: map[string]*pb.Permission{
					"storage.viewer": {
						Buckets: []string{"/bar", "/goo"},
					},
				},
			},
		},
		{},
	} {
		// 評価する
		result, detail, err := filter.Eval(data)
		dataJSON, _ := json.Marshal(data)
		if err != nil {
			fmt.Printf("input=%s detail=%+v err=%+v\n", dataJSON, detail, err)
			continue
		}
		fmt.Printf("input=%s result=%+v\n", dataJSON, result)
	}
}

func newCELEvaluator() (cel.Program, error) {
	env, err := cel.NewEnv(
		cel.Types(&pb.Data{}),
		cel.Declarations(
			decls.NewVar("data", decls.NewObjectType("cel_sandbox.Data")),
		),
	)
	if err != nil {
		return nil, err
	}
	ast, iss := env.Parse(`data.org == 'Hoge campany' && ('/foo' in data.permissions['storage.viewer'].buckets)`)
	if iss.Err() != nil {
		return nil, xerrors.Errorf(": %w", iss.Err())
	}
	checked, iss := env.Check(ast)
	if iss.Err() != nil {
		return nil, xerrors.Errorf(": %w", iss.Err())
	}
	if !proto.Equal(checked.ResultType(), decls.Bool) {
		return nil, xerrors.Errorf(": %w", err)
	}
	program, err := env.Program(checked)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	return program, nil
}

func printAst(ast *cel.Ast) {
	fmt.Printf("%+v\n", ast)
	s, _ := json.MarshalIndent(ast.Expr(), "", " ")
	fmt.Printf("%+v\n", string(s))
}
