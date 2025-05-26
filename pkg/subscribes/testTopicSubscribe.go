package subscribes

import (
	"fmt"
)

type TestTopic struct {
	A int
	B int
}

func TestTopicSubscribe(topic TestTopic) {
	fmt.Printf("a-->%d b-->%d", topic.A, topic.B)
}
