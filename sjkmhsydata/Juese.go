package sjkmhsydata

import "changliang/zf"

type Juese struct{}

func (js *Juese) Miaoshu() string {
	return zf.Zfs.Miaoshu(false)
}
func (js *Juese) Mingcheng() string {
	return zf.Zfs.Mingcheng(false)
}
