package tests

import (
	"gongju"
	"gongju/peizhi"
	"log"
	"testing"
)

func TestFanshejiexi(t *testing.T) {
	shezhi := gongju.Shezhi{}
	fansheg := gongju.Fanshejiexi(
		&shezhi,
		peizhi.Pz{},
	)
	shezhiret := fansheg.(*gongju.Shezhi)
	log.Println("fansheg--------", shezhiret)
}
