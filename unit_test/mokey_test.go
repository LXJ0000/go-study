package main

import (
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey"
)

type Boy struct {
}

func (b *Boy) Speak() string {
	return "hhh"
}

func Test_gomonkey(t *testing.T) {
	b := Boy{}

	patch := gomonkey.ApplyMethod(reflect.TypeOf(&Boy{}), "Speak", func(b *Boy) string {
		return "555"
	})
	defer patch.Reset()

	t.Logf("Speak: %s", b.Speak())

}

func Laugh() string {
	return "hhh"
}

func Test_gomonkey2(t *testing.T) {
	patch := gomonkey.ApplyFunc(Laugh, func() string {
		return "555"
	})
	defer patch.Reset()

	t.Logf("Laugh: %s", Laugh()) // Output: Laugh: 555555
}
