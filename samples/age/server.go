package main

import (
	"time"

	tp "github.com/henrylee2cn/teleport"
)

func main() {
	svr := tp.NewPeer(tp.PeerConfig{
		CountTime:         true,
		ListenAddress:     ":9090",
		DefaultSessionAge: time.Second * 7,
		DefaultContextAge: time.Second,
	})
	svr.RoutePull(new(test))
	svr.Listen()
}

type test struct {
	tp.PullCtx
}

func (t *test) Ok(args *string) (string, *tp.Rerror) {
	return *args + " -> OK", nil
}

func (t *test) Timeout(args *string) (string, *tp.Rerror) {
	time.Sleep(time.Second)
	tCtx := t.Context()
	select {
	case <-tCtx.Done():
		return *args + " -> Not Timeout", nil
		return "", tp.NewRerror(
			tp.CodeHandleTimeout,
			tp.CodeText(tp.CodeHandleTimeout),
			tCtx.Err().Error(),
		)
	default:
	}
	return *args + " -> Not Timeout", nil
}
