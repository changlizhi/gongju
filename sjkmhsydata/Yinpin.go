package sjkmhsydata

import "changliang/zf"

type Yinpin struct{}

func (yp *Yinpin) Lujing() string {
	return zf.Zfs.Lujing(false)
}
func (yp *Yinpin) Jiamifangshi() string {
	return zf.Zfs.Jiamifangshi(false)
}
func (yp *Yinpin) Mingcheng() string {
	return zf.Zfs.Mingcheng(false)
}
func (yp *Yinpin) Miaoshu() string {
	return zf.Zfs.Miaoshu(false)
}
