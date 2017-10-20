package sjkmhsydata

import "changliang/zf"

type Yinpinxiazai struct{}

func (ypxz *Yinpinxiazai) Weiyibianma() string {
	return zf.Zfs.Weiyibianma(false)
}
func (ypxz *Yinpinxiazai) Yinpinbianma() string {
	return zf.Zfs.Yinpinbianma(false)
}
func (ypxz *Yinpinxiazai) Jinzhixiazai() string {
	return zf.Zfs.Jinzhixiazai(false)
}
