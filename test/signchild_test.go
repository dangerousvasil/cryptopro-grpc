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

func TestSignChild(t *testing.T) {

	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithBlock())

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	go func() {
		time.Sleep(time.Second * 2)
		log.Println(`cancel`)
		cancel()
	}()

	conn, err := grpc.DialContext(ctx, `localhost:32955`, opts...)
	if err != nil {
		log.Println(`Dial`, err)
		return
	}
	c := grpc_service.NewServiceInternalClient(conn)
	//defer cancel()

	serviceStream, err := c.Sign(ctx)
	err = serviceStream.Send(&grpc_service.SignRequest{
		Storage: "MY",
		Content: []byte("HelloWorld"),
		Key:     "762f84827e9199f72043f01d548ad86503b7fa98",

		Flag: 0,
	})
	if err != nil {
		log.Println(`Dial`, err)
		return
	}
	res, err := serviceStream.Recv()
	log.Println(`Done`)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	log.Println(string(res.Content))

}
