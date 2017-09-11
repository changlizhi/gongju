package gongju

import (
	"strings"
	"changliang/zf"
	"regexp"
	"io/ioutil"
	"changliang/zfzhi"
	"changliang/biaolie"
	"reflect"
	"log"
)

func Biaolies(mokuaiming string) (biaolie map[string]string, biao map[string]string, lie map[string]string) {
	biaolie = make(map[string]string)
	biao = make(map[string]string)
	lie = make(map[string]string)
	ffs := Pipeifangfa(
		Getjichupath(),
		zf.Zfs.Sjk(true) + Mokuaimings[zf.Zfs.Hfxyonghu(false)].Zhi,
		zf.Zfs.Sjk(false) + Mokuaimings[zf.Zfs.Hfxyonghu(false)].Zhi,
		zf.Zfs.Sjk(true) + Mokuaimings[zf.Zfs.Hfxyonghu(false)].Zhi,
		zf.Zfs.Biaolie(false))
	// "[A-Z][a-z]+"//正则表达式匹配驼峰命名的方法
	repstr := zfzhi.Zhi.Zkhz() +
		zf.Zfs.A(false) + zfzhi.Zhi.Jian() + zf.Zfs.Z(false) +
		zfzhi.Zhi.Zkhy() +
		zfzhi.Zhi.Zkhz() +
		zf.Zfs.A(true) + zfzhi.Zhi.Jian() + zf.Zfs.Z(true) +
		zfzhi.Zhi.Zkhy() +
		zfzhi.Zhi.Jia()
	rep, _ := regexp.CompilePOSIX(repstr)

	for _, ff := range ffs {
		biaolie[ff] = strings.ToLower(ff)
		fp := rep.FindAllString(ff, zfzhi.Zhi.Shuzifu1())
		biao[fp[zfzhi.Zhi.Shuzi0()]] = strings.ToLower(fp[zfzhi.Zhi.Shuzi0()])
		lie[fp[zfzhi.Zhi.Shuzi1()]] = strings.ToLower(fp[zfzhi.Zhi.Shuzi1()])
	}
	return
}

func Suoyoubiaojiegou(mokuaiming string) map[string][]string {
	ret := make(map[string][]string)
	// "[A-Z][a-z]+"//正则表达式匹配驼峰命名的方法
	repstr := zfzhi.Zhi.Zkhz() +
		zf.Zfs.A(false) + zfzhi.Zhi.Jian() + zf.Zfs.Z(false) +
		zfzhi.Zhi.Zkhy() +
		zfzhi.Zhi.Zkhz() +
		zf.Zfs.A(true) + zfzhi.Zhi.Jian() + zf.Zfs.Z(true) +
		zfzhi.Zhi.Zkhy() +
		zfzhi.Zhi.Jia()
	rep, _ := regexp.CompilePOSIX(repstr)

	biaolie, biao, _ := Biaolies(mokuaiming)
	for bk, _ := range biao {
		lies := []string{}
		for blk, _ := range biaolie {
			fp := rep.FindAllString(blk, zfzhi.Zhi.Shuzifu1())
			if fp[zfzhi.Zhi.Shuzi0()] == bk {
				lies = append(lies, fp[zfzhi.Zhi.Shuzi1()])
			}
		}
		ret[bk] = lies
	}
	return ret
}
func Liechangdu(lieming string) int {
	sjk := biaolie.Bl{}
	ffm := lieming + zf.Zfs.Changdu(true)
	v := reflect.ValueOf(&sjk)
	ret := v.MethodByName(ffm).Call(nil)[zfzhi.Zhi.Shuzi0()].Interface().(int)
	return ret
}

func Lieleixing(lieming string) string {
	sjk := biaolie.Bl{}
	ffm := lieming + zf.Zfs.Leixingzhi(true)
	v := reflect.ValueOf(&sjk)
	ret := v.MethodByName(ffm).Call(nil)[zfzhi.Zhi.Shuzi0()].Interface().(string)
	if ret == zfzhi.Zhi.Kzf() {
		log.Println("类型错误-----空")
		return zfzhi.Zhi.Kzf()
	}
	return ret
}

func Biao(mokuaiming string, mingcheng string) []string {
	biaojiegou := Suoyoubiaojiegou(mokuaiming)
	return biaojiegou[mingcheng]
}

func Pipeifangfa(pathfrom string, canshu string, canshuleixing string, muluming string, wenjianming string) []string {
	fu1 := zfzhi.Zhi.Shuzifu1()

	path := pathfrom + zfzhi.Zhi.Xx() + muluming +
		zfzhi.Zhi.Xx() + wenjianming + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
	b, _ := ioutil.ReadFile(path)
	//从匹配的方法名中去除前面对于的正则表达式
	// func \(zf \*Zf\)
	// func \(sjk \*Sjk\)
	// func \(zfzhi \*Zfzhi\)
	repstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Fxx() +
		zfzhi.Zhi.Xkhz() + canshu + zfzhi.Zhi.Kgf() +
		zfzhi.Zhi.Fxx() + zfzhi.Zhi.Xh() + canshuleixing +
		zfzhi.Zhi.Fxx() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Kgf()
	//拼接从文件中匹配方法名的正则表达式
	// func \(zf \*Zf\) [[:word:]]+
	// func \(sjk \*Sjk\) [[:word:]]+
	// func \(zfzhi \*Zfzhi\) [[:word:]]+
	regstr := repstr + zfzhi.Zhi.Zkhz() + zfzhi.Zhi.Zkhz() +
		zfzhi.Zhi.Mh() + zf.Zfs.Word(true) + zfzhi.Zhi.Mh() +
		zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Jia()
	//生成正则表达式
	reg, _ := regexp.CompilePOSIX(regstr)
	repreg, _ := regexp.CompilePOSIX(repstr)
	//从文件中匹配方法名
	sfind := reg.FindAllString(string(b), fu1)
	// 匹配到的方法名合在一起
	joinstr := strings.Join(sfind, zfzhi.Zhi.Kgf())
	//去除前缀func (zf *Zf)
	srep := repreg.ReplaceAllString(joinstr, zfzhi.Zhi.Kzf())
	// 再恢复为数组返回调用
	ret := strings.Split(srep, zfzhi.Zhi.Kgf())

	return ret
}
