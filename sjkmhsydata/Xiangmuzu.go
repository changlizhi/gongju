package sjkmhsydata

import "changliang/zf"

type Xiangmuzu struct{}

func (xmz *Xiangmuzu) Mingcheng() string {
	return zf.Zfs.Mingcheng(false)
}
func (xmz *Xiangmuzu) Miaoshu() string {
	return zf.Zfs.Miaoshu(false)
}
