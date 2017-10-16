package peizhi

import "changliang/zf"

func (jm *Jsonmoji)Guilei() string {
	return zf.Zfs.Guilei(true)
}
func (jm *Jsonmoji)Bianma() string {
	return zf.Zfs.Bianma(true)
}
func (jm *Jsonmoji)Mingcheng() string {
	return zf.Zfs.Mingcheng(true)
}
func (jm *Jsonmoji)Zhi() string {
	return zf.Zfs.Zhi(true)
}
