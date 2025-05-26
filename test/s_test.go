package test

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"
)

type ts struct {
}

type Query struct {
}

func (q *Query) push() {
	//把连接放到连接池里
}

func (q *Query) Find(p interface{}) {
	q.exec(p, q.push)
}

func (q *Query) exec(p interface{}, r func()) {
	defer r()
	// p------
}

var ErrUserNotFound = errors.New("not User")

var ErrHandler = make(map[error]func(err error))

func (receiver *ts) Aname() *ts {

	fmt.Println(1)

	return receiver
}

func (receiver *ts) Bname() *ts {

	fmt.Println(2)

	return receiver

}

func (receiver *ts) Cname() {

	fmt.Println(9)
}

func TestRun(t *testing.T) {
	regiter(ErrUserNotFound, func(err error) {
		fmt.Println(err)
	})

	regiter(os.ErrNotExist, func(err error) {
		fmt.Println(err)
	})
	fmt.Println(reflect.TypeOf(os.ErrNotExist).Kind())
	fmt.Println(reflect.ValueOf(time.Time{}).Interface().(time.Time))
	fmt.Println(reflect.TypeOf(ErrUserNotFound).Kind())
	ErrHandler[os.ErrNotExist](os.ErrNotExist)
}

func regiter(err error, fn func(err error)) {
	ErrHandler[err] = fn
}
