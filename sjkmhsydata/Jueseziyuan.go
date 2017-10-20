package sjkmhsydata

import "changliang/zf"

type Jueseziyuan struct{}

func (jszy *Jueseziyuan) Juesebianma() string {
	return zf.Zfs.Juesebianma(false)
}
func (jszy *Jueseziyuan) Ziyuanbianma() string {
	return zf.Zfs.Ziyuanbianma(false)
}
