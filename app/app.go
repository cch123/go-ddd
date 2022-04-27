package app

import (
	"context"
	"fmt"
	"github.com/cch123/go-ddd/domain/aggregate"
	"github.com/cch123/go-ddd/repo/data"
)

func Main() {
	app := initApp(1)
	err := app.run()
	if err != nil {
		// log fatal
	}

	fmt.Println(app)

	c, err := data.NewClient("root:@tcp(localhost:3306)/shopping?parseTime=True")
	if err != nil {
		fmt.Println("database init failed", err)
		return
	}
	customer, err := c.Customer.Get(context.TODO(), 1)
	fmt.Println("yes", customer, err)

	app.removeIllegalOrder(context.Background())
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
