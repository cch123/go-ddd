package app

import (
	"context"
	"fmt"
	"github.com/cch123/go-ddd/app/protocol"
	"github.com/cch123/go-ddd/domain/aggregate"
)

func Main() {
	app, err := initApp()
	if err != nil {
		// log fatal
		fmt.Println(err)
		return
	}

	err = app.run()
	if err != nil {
		// log fatal
		fmt.Println(err)
		return
	}

	fmt.Println(app)

	app.removeIllegalOrder(context.Background())
	protocol.InitHTTPServer()
}

type app struct {
	customerAgg *aggregate.CustomerAgg
	orderAgg    *aggregate.OrderAgg
}

func newApp(c *aggregate.CustomerAgg, o *aggregate.OrderAgg) *app {
	return &app{customerAgg: c, orderAgg: o}
}

func (a *app) run() error {
	return nil
}

func (a *app) removeIllegalOrder(ctx context.Context) {
	fmt.Println("removing illegal order")
	req := &aggregate.RemoveIllegalOrderReq{}
	resp, err := a.orderAgg.RemoveIllegalOrder(ctx, req)
	if err != nil {
	}

	fmt.Println("remove result", resp)
}
