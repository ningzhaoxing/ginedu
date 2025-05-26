package subscribes

import "fmt"

func UserPrivilege(p UserIntegralTopic) {
	fmt.Println("赠送了 特权", "用户是：", p.User)
}
