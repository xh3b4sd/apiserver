package main

import (
	"context"
	"fmt"
	"net"

	"github.com/xh3b4sd/tracer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/xh3b4sd/gocode/pkg/pbf/post"
)

func main() {
	err := mainE(context.Background())
	if err != nil {
		tracer.Panic(err)
	}
}

func mainE(ctx context.Context) error {
	var err error

	var l net.Listener
	{
		l, err = net.Listen("tcp", fmt.Sprintf(":%d", 7777))
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var a post.APIServer
	{
		g := grpc.NewServer()
		reflection.Register(g)

		a = &API{}
		post.RegisterAPIServer(g, a)

		err := g.Serve(l)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

// -------------------------------------------------------------------------- //

type API struct {
	post.UnimplementedAPIServer
}

func (a *API) Create(ctx context.Context, cre *post.CreateI) (*post.CreateO, error) {
	fmt.Printf("%#v\n", cre)
	return &post.CreateO{Message: "create output"}, nil
}

func (a *API) Delete(ctx context.Context, del *post.DeleteI) (*post.DeleteO, error) {
	fmt.Printf("%#v\n", del)
	return &post.DeleteO{Message: "delete output"}, nil
}

func (a *API) Search(ctx context.Context, sea *post.SearchI) (*post.SearchO, error) {
	fmt.Printf("%#v\n", sea)
	return &post.SearchO{Message: "search output"}, nil
}

func (a *API) Update(ctx context.Context, upd *post.UpdateI) (*post.UpdateO, error) {
	fmt.Printf("%#v\n", upd)
	return &post.UpdateO{Message: "update output"}, nil
}
