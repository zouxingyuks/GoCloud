package email

import (
	"fmt"
	"testing"
)

func Test_defaultEmailTmpl(t *testing.T) {
	fmt.Println(defaultEmailTmpl())
}

func Test_getTemplate(t *testing.T) {
	fmt.Println(getTemplate())
}
