package main

import (
	"context"
	echo "go.buf.build/loopholelabs/frpc/loopholelabs/echo-example"
	"os"
	"os/signal"
	"syscall"
)

type svc struct{}

func (s *svc) Echo(_ context.Context, req *echo.Request) (*echo.Response, error) {
	res := new(echo.Response)
	res.Message = req.Message
	return res, nil
}

func main() {
	s, err := echo.NewServer(new(svc), nil, nil)
	if err != nil {
		panic(err)
	}
	err = s.Start("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	err = s.Shutdown()
	if err != nil {
		panic(err)
	}
}
