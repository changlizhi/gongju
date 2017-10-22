package peizhi

import (
	"gongju/sjkmh"
)

type Pz struct {
	Cshpz *Chushihua
	Mkmpz *Mokuaiming
	Jmpz  *Jsonmoji
	Jlpz  *Jsonlie
}

type Chushihua struct{}

var Cshs = &Chushihua{}

type Mokuaiming struct {
	Mhsydatamk sjkmh.Mhsydata
}

var Mkms = &Mokuaiming{}

type Jsonmoji struct{}

var Jms = &Jsonmoji{}

type Jsonlie struct {
	Sz  Shezhi
	Gjh Guojihua
}

var Jls = &Jsonlie{}
