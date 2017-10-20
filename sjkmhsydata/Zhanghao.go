package sjkmhsydata

import "changliang/zf"

type Zhanghao struct {
}

func (zh *Zhanghao) Leixing() string {
	return zf.Zfs.Leixing(false)
}
func (zh *Zhanghao) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (zh *Zhanghao) Weiyibianma() string {
	return zf.Zfs.Weiyibianma(false)
}
