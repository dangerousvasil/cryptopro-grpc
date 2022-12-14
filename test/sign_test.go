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

func TestSign(t *testing.T) {
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

	res, err := c.Sign(ctx, &grpc_service.SignRequest{
		Content: append([]byte("HelloWorld"), make([]byte, 40)...),
		Key:     "762f84827e9199f72043f01d548ad86503b7fa98",
	})

	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	log.Println(`Done`)
	log.Println(string(res.Content))
}

func BenchmarkTemplateParallel(b *testing.B) {
	log.SetFlags(log.Lshortfile)
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	ctx := context.Background()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {

			//opts = append(opts, grpc.WithBlock())
			//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			//go func() {
			//	time.Sleep(time.Second * 2)
			//	log.Println(`cancel`)
			//cancel()
			//}()
			conn, err := grpc.DialContext(ctx, `localhost:8080`, opts...)
			if err != nil {
				log.Println(`Dial`, err)
				b.FailNow()
			}
			c := grpc_service.NewServiceClient(conn)
			res, err := c.Sign(ctx, &grpc_service.SignRequest{
				Content: append([]byte("HelloWorld"), make([]byte, 40)...),
				Key:     "762f84827e9199f72043f01d548ad86503b7fa98",
			})

			if err != nil {
				log.Println(err)
				b.FailNow()
			}
			if res.Description != `` {
				b.FailNow()
			}
			_ = conn.Close()

			//log.Println(`Done`)
			//log.Println(string(res.Content))
		}
	})
}
