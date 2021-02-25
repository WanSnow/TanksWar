package parts

import (
	"time"
)

type Engine interface {
	EngineCooling() func() bool //引擎冷却
	CycleInstruction()          //周期指令
}

func EngineFactory(code int) Engine {
	switch code {
	case EngineLv1:
		return Lv1Engine{efficiency: 1 * 1e8}
	case EngineLv2:
		return Lv2Engine{efficiency: 8 * 1e7}
	case EngineLv3:
		return Lv3Engine{efficiency: 5 * 1e7}
	case EngineLv4:
		return Lv4Engine{efficiency: 2 * 1e7}
	default:
		return NullEngine{efficiency: -1}
	}
}

/*
TODO 引擎执行
*/
func (e Lv1Engine) EngineCooling() func() bool {
	initExec := -1
	lastExec := time.Now()
	return func() bool {
		if time.Since(lastExec).Nanoseconds() > int64(e.efficiency) {
			lastExec = time.Now()
			return true
		} else if initExec < 0 {
			initExec++
			return true
		} else {
			return false
		}
	}
}

func (e Lv1Engine) CycleInstruction() {

}

func (e Lv2Engine) EngineCooling() func() bool {
	initExec := -1
	lastExec := time.Now()
	return func() bool {
		if time.Since(lastExec).Nanoseconds() > int64(e.efficiency) {
			lastExec = time.Now()
			return true
		} else if initExec < 0 {
			initExec++
			return true
		} else {
			return false
		}
	}
}

func (e Lv2Engine) CycleInstruction() {

}

func (e Lv3Engine) EngineCooling() func() bool {
	initExec := -1
	lastExec := time.Now()
	return func() bool {
		if time.Since(lastExec).Nanoseconds() > int64(e.efficiency) {
			lastExec = time.Now()
			return true
		} else if initExec < 0 {
			initExec++
			return true
		} else {
			return false
		}
	}
}

func (e Lv3Engine) CycleInstruction() {

}

func (e Lv4Engine) EngineCooling() func() bool {
	initExec := -1
	lastExec := time.Now()
	return func() bool {
		if time.Since(lastExec).Nanoseconds() > int64(e.efficiency) {
			lastExec = time.Now()
			return true
		} else if initExec < 0 {
			initExec++
			return true
		} else {
			return false
		}
	}
}

func (e Lv4Engine) CycleInstruction() {

}

func (null NullEngine) EngineCooling() func() bool {
	return func() bool {
		return false
	}
}

func (null NullEngine) CycleInstruction() {

}
