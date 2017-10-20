package sjkmhsydata

import "changliang/zf"

type Liuyan struct{}

func (ly *Liuyan) Liuyanyonghu() string {
	return zf.Zfs.Liuyanyonghu(false)
}
func (ly *Liuyan) Huifuyonghu() string {
	return zf.Zfs.Huifuyonghu(false)
}
func (ly *Liuyan) Biaoti() string {
	return zf.Zfs.Biaoti(false)
}
func (ly *Liuyan) Neirong() string {
	return zf.Zfs.Neirong(false)
}
func (ly *Liuyan) Shijian() string {
	return zf.Zfs.Shijian(false)
}
