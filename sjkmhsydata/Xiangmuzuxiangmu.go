package sjkmhsydata

import "changliang/zf"

type Xiangmuzuxiangmu struct{}

func (xmzxm *Xiangmuzuxiangmu) Xiangmuzubianma() string {
	return zf.Zfs.Xiangmuzubianma(false)
}
func (xmzxm *Xiangmuzuxiangmu) Xiangmubianma() string {
	return zf.Zfs.Miaoshu(false)
}
