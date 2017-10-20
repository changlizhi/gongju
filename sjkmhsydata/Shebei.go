package sjkmhsydata

import "changliang/zf"

type Shebei struct{}

func (sb *Shebei) Mingcheng() string {
	return zf.Zfs.Mingcheng(false)
}
func (sb *Shebei) Shebeihao() string {
	return zf.Zfs.Shebeihao(false)
}
func (sb *Shebei) Miaoshu() string {
	return zf.Zfs.Miaoshu(false)
}
func (sb *Shebei) Guize() string {
	return zf.Zfs.Guize(false)
}
