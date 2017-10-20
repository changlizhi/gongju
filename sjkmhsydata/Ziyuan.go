package sjkmhsydata

import "changliang/zf"

type Ziyuan struct{}

func (zy *Ziyuan) Mingcheng() string {
	return zf.Zfs.Mingcheng(false)
}
func (zy *Ziyuan) Fubianma() string {
	return zf.Zfs.Fubianma(false)
}
func (zy *Ziyuan) Lujing() string {
	return zf.Zfs.Lujing(false)
}
func (zy *Ziyuan) Leixing() string {
	return zf.Zfs.Leixing(false)
}
