package sjkmhsydata

import "changliang/zf"

type Zhanghaojuese struct{}

func (zhjs *Zhanghaojuese) Weiyibianma() string {
	return zf.Zfs.Weiyibianma(false)
}
func (zhjs *Zhanghaojuese) Juesebianma() string {
	return zf.Zfs.Juesebianma(false)
}
