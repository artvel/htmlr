package htmlr

import (
	"bytes"
	"fmt"
	"testing"
)

func TestResolveLinear(t *testing.T) {
	src := resolve("./res/main1.html")
	exp, er := fileToBytes("./res/main1_expected.html")
	if er != nil {
		t.Error(er)
	}
	if !bytes.Equal(exp, src) {
		fmt.Println(string(src))
		t.Error("template not as expected")
	}
}

func TestResolveRecursive(t *testing.T) {
	src := resolve("./res/main.html")
	exp, er := fileToBytes("./res/main_expected.html")
	if er != nil {
		t.Error(er)
	}
	if !bytes.Equal(exp, src) {
		fmt.Println(string(src))
		t.Error("template not as expected")
	}
}
