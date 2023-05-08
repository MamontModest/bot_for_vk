package gen

import (
	"math/rand"
	"time"
)

var s = "1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"

func Password() string {
	rand.Seed(time.Now().UnixMilli())
	pass := ""
	for i := 0; i < 20; i++ {
		j := rand.Intn(len(s))
		pass += string(s[j])
	}
	return pass
}
