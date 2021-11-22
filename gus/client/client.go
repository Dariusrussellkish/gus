package client

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Dariusrussellkish/gus/gus/message"
	"github.com/apache/thrift/lib/go/thrift"
)

var defaultCtx = context.Background()

const iters = 100_000

func handleClient(client *message.MessengerClient) (err error) {

	durations := make([]time.Duration, iters)

	for i := 0; i < iters; i++ {
		msg := message.Message{
			ID: int32(i),
			To:      "Server",
			Frm:     "Client",
			Content: "Ping!",
		}
		start := time.Now()
		_, err := client.SendMessage(defaultCtx, &msg)
		for err != nil {
			return fmt.Errorf("client encountered error on msg %d, %e", i, err)
		}
		end := time.Now()
		durations[i] = end.Sub(start)
	}

	f, err := os.Create("thrift_timings.dat")
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)
	if err != nil {
		return err
	}
	w := bufio.NewWriter(f)
	for _, d := range durations {
		_, err := w.WriteString(fmt.Sprintf("%d\n", d.Microseconds()))
		if err != nil {
			panic(err)
		}
	}
	err = w.Flush()
	if err != nil {
		return err
	}
	return nil
}

func RunClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool, cfg *thrift.TConfiguration) error {
	var transport thrift.TTransport
	if secure {
		transport = thrift.NewTSSLSocketConf(addr, cfg)
	} else {
		transport = thrift.NewTSocketConf(addr, cfg)
	}
	transport, err := transportFactory.GetTransport(transport)
	if err != nil {
		return err
	}
	defer func(transport thrift.TTransport) {
		err := transport.Close()
		if err != nil {
			panic(err)
		}
	}(transport)
	if err := transport.Open(); err != nil {
		return err
	}
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	return handleClient(message.NewMessengerClient(thrift.NewTStandardClient(iprot, oprot)))
}