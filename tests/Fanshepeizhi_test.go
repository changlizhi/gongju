package tests

import (
	"log"
	"testing"
	"gongju"
	"jichu/peizhi"
)

func TestFanshejiexi(t *testing.T) {
	shezhi := gongju.Shezhi{}
	fansheg := gongju.Fanshejiexi(
		&shezhi,
		peizhi.Pz{},
	)
	log.Println("fansheg--------", fansheg)
}
