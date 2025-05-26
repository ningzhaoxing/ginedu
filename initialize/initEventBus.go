package initialize

import (
	"fmt"
	"gindeu/pkg/globals"
	"gindeu/pkg/subscribes"
	"github.com/asaskevich/EventBus"
)

func InitEventBus() EventBus.Bus {
	bus := EventBus.New()

	err := bus.Subscribe(globals.TestTopic, subscribes.TestTopicSubscribe)
	err = bus.Subscribe(globals.UserIntegralTopic, subscribes.UserCoupon)
	err = bus.Subscribe(globals.UserIntegralTopic, subscribes.UserMedal)
	//err = bus.Subscribe(globals.DemoErrTopic, subscribes.DemoSubscribe)

	if err != nil {
		panic(fmt.Errorf("initEventBus Error--> %v", err))
	}

	return bus
}
