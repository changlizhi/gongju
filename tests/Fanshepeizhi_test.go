package tests

import (
	"log"
	"mhsydata/chushihuas"
	"mhsydata/zdguojihua"
	"testing"
	"gongju"
)

func TestFanshejiexi(t *testing.T) {
	//shezhi := chushihuas.Shezhi{}
	//fanshes := gongju.Fanshejiexi(
	//	&shezhi,
	//	zdpeizhi.Pz{},
	//)
	//shezhifh := fanshes.(*chushihuas.Shezhi)
	//log.Println("shezhifh=========", shezhifh)

	guojihua := chushihuas.Guojihua{}
	fansheg := gongju.Fanshejiexi(
		&guojihua,
		zdguojihua.Gjh{},
	)
	log.Println("fansheg--------", fansheg)
}
