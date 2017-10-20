package sjkmhsydata

import "changliang/zf"

type Wenzhang struct{}

func (wz *Wenzhang) Mingcheng() string {
	return zf.Zfs.Mingcheng(false)
}
func (wz *Wenzhang) Neirong() string {
	return zf.Zfs.Neirong(false)
}
func (wz *Wenzhang) Shifoufabu() string {
	return zf.Zfs.Shifoufabu(false)
}
