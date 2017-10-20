package sjkmhsydata

import "changliang/zf"

type Zidian struct{}

func (zd *Zidian) Mingcheng() string {
	return zf.Zfs.Mingcheng(false)
}
func (zd *Zidian) Miaoshu() string {
	return zf.Zfs.Miaoshu(false)
}
func (zd *Zidian) Fubianma() string {
	return zf.Zfs.Fubianma(false)
}
