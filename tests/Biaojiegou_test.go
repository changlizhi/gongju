package tests

import (
	"testing"
	"gongju"
	"changliang/zf"
)

func TestFanshebiaojiegou(t *testing.T) {
	gongju.Fanshebiaojiegou(zf.Zfs.Mhsydata(true))
}
