package tankLib

import . "TanksWar/tankLib/implement"

type Tank struct {
	hp int //生命值

	engine Engine
	armor  Armor
	weapon Weapon
	scope  Scope
	wheel  Wheel
	bullet Bullet
}

func TankFactory(engine int, armor int, weapon int, scope int, wheel int, bullet int) Tank {
	tank := Tank{}
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
	execute(command string, args ...int)
	trigger(command string, args ...int)
}
