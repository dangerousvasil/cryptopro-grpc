package test

import (
	"context"
	"cryptopro-jsonrpc/src/lib/grpc_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testing"
	"time"
)

func TestType(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithBlock())
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	go func() {
		time.Sleep(time.Second * 2)
		log.Println(`cancel`)
		cancel()
	}()

	conn, err := grpc.DialContext(ctx, `localhost:8080`, opts...)
	if err != nil {
		log.Println(`Dial`, err)
		return
	}
	c := grpc_service.NewServiceClient(conn)

	res, err := c.Type(ctx, &grpc_service.ContentRequest{
		Content: append([]byte("HelloWorld"), make([]byte, 40)...),
		//Key:     "762f84827e9199f72043f01d548ad86503b7fa98",
	})

	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	log.Println(`Done`)
	log.Println((res))
}
