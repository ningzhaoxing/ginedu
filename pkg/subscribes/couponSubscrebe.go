package subscribes

import (
	"fmt"
)

type UserIntegralTopic struct {
	User           string
	BeforeIntegral int
	AfterIntegral  int
}

func UserCoupon(p UserIntegralTopic) {
	fmt.Println("赠送了 10 元优惠券", "用户是：", p.User)
}
