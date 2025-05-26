package subscribes

import "fmt"

func UserMedal(p UserIntegralTopic) {
	fmt.Println("赠送了 勋章", "用户是：", p.User)
}
