package peizhi

import "changliang/zf"

func (jl *Jsonlie)Guojihua() string {
	return zf.Zfs.Guojihua(true)
}
func (jl *Jsonlie)Shezhi() string {
	return zf.Zfs.Shezhi(true)
}