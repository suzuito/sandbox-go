package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
	"github.com/suzuito/sandbox-go/cmd/002/pb"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

type server struct {
	pb.UnimplementedViewStorageServiceServer
	program cel.Program
}

func (s *server) ReadFile(ctx context.Context, request *pb.RequestReadFile) (*pb.File, error) {
	result, _, err := s.program.Eval(map[string]interface{}{"request": request})
	if err != nil {
		return nil, xerrors.Errorf("eval error : %+v", err)
	}
	fmt.Println(result.Value().(bool))
	return &pb.File{
		Path: "foo/bar/hoge.txt",
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	program, err := newProgram(`request.org == 'Hoge campany' && ('/foo' in request.permissions['storage.viewer'].buckets)`)
	if err != nil {
		log.Fatalf("failed: %+v", err)
	}
	s := grpc.NewServer()
	pb.RegisterViewStorageServiceServer(s, &server{
		program: program,
	})
	fmt.Println("running")
	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}

func newProgram(txt string) (cel.Program, error) {
	env, err := cel.NewEnv(
		cel.Types(&pb.RequestReadFile{}),
		cel.Declarations(
			decls.NewVar(
				"request",
				decls.NewObjectType("cel_sandbox.RequestReadFile"),
			),
		),
	)
	if err != nil {
		return nil, err
	}
	// ast, iss := env.Parse(`request.org == 'Hoge campany' && ('/foo' in request.permissions['storage.viewer'].buckets)`)
	ast, iss := env.Parse(txt)
	if iss.Err() != nil {
		return nil, iss.Err()
	}
	checked, iss := env.Check(ast)
	if iss.Err() != nil {
		return nil, iss.Err()
	}
	if !proto.Equal(checked.ResultType(), decls.Bool) {
		return nil, fmt.Errorf("Got %v, wanted %v result type", checked.ResultType(), decls.String)
	}
	program, err := env.Program(checked)
	if err != nil {
		return nil, fmt.Errorf("program error: %v", err)
	}
	return program, nil
}
