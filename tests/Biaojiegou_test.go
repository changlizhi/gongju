package tests

import (
	"changliang/zf"
	"gongju"
	"testing"
)

func TestFanshebiaojiegou(t *testing.T) {
	gongju.Fanshebiaojiegou(zf.Zfs.Mhsydata(true))
	//gongju.Fanshejichulie()
}
