package tankLib

import (
	. "TanksWar/tankLib/model/parts"
)

type Tank struct {
	tankID string

	hp int //生命值
	//direction int //存入redis

	engine Engine
	armor  Armor
	weapon Weapon
	scope  Scope
	wheel  Wheel
	bullet Bullet
}

func TankFactory(id string, engine int, armor int, weapon int, scope int, wheel int, bullet int) Tank {
	tank := Tank{}
	tank.tankID = id
	tank.hp = 100
	tank.engine = EngineFactory(engine)
	tank.armor = ArmorFactory(armor)
	tank.weapon = WeaponFactory(weapon)
	tank.scope = ScopeFactory(scope)
	tank.wheel = WheelFactory(wheel)
	tank.bullet = BulletFactory(bullet)

	return tank
}

type TankFunc interface {
	Run()
	Execute(command string, args ...int)
	Trigger(command string, args ...int)
}

/*
todo 坦克中断操作
*/
func (tank Tank) Trigger() {

}

func (tank Tank) Run() {
	engineCooling := tank.engine.EngineCooling() //引擎冷却计时
loop:
	for tank.hp > 0 {
		if true { //todo 坦克的中断触发
			break loop
		} else if engineCooling() {
			tank.engine.CycleInstruction()
		}
	}
}
