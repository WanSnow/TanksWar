package model

import . "tankLib/model/parts"

type Tank struct {
	hp int //生命值

	engine Engine
	armor  Armor
	weapon Weapon
	scope  Scope
	wheel  Wheel
	bullet Bullet
}

type TankFunc interface {
	execute(command string, args ...int)
	trigger(command string, args ...int)
}
