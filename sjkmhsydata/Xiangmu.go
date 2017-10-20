package sjkmhsydata

import "changliang/zf"

type Xiangmu struct{}

func (xm *Xiangmu) Mingcheng() string {
	return zf.Zfs.Mingcheng(false)
}
func (xm *Xiangmu) Danwei() string {
	return zf.Zfs.Danwei(false)
}
func (xm *Xiangmu) Leixing() string {
	return zf.Zfs.Leixing(false)
}
func (xm *Xiangmu) Miaoshu() string {
	return zf.Zfs.Miaoshu(false)
}
func (xm *Xiangmu) Yixuebianma() string {
	return zf.Zfs.Yixuebianma(false)
}
