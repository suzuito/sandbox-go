package main

import (
	"encoding/json"
	"fmt"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
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
		{"org": "Hoge campany", "viewableStorageBuckets": []string{"/foo", "/goo"}},
		{"org": "Fuga campany", "viewableStorageBuckets": []string{"/foo", "/goo"}},
		{"org": "Hoge campany", "viewableStorageBuckets": []string{"/bar", "/goo"}},
		{"org": "Hoge campany", "viewableStorageBuckets": []string{}},
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
		cel.Declarations(
			decls.NewVar("org", decls.String),
			decls.NewVar("viewableStorageBuckets", decls.NewListType(decls.String)),
		),
	)
	if err != nil {
		return nil, err
	}
	ast, iss := env.Parse(`org == 'Hoge campany' && ('/foo' in viewableStorageBuckets)`)
	if iss.Err() != nil {
		return nil, xerrors.Errorf(": %w", iss.Err())
	}
	// printAst(ast)
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
