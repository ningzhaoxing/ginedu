package subscribes

import (
	"fmt"
	"gindeu/pkg/errors"
)

func DemoSubscribe(err *errors.DemoError) {
	fmt.Println(err.Error(), "---->")
}
