package main

func main() {
	/* Instance Weapons */
	sword := NewWeapon("Iron Sword", 8, 80, 15, "Sword")
	axe := NewWeapon("Iron Axe", 10, 70, 5, "Axe")
	// spear := NewWeapon("Iron Spear", 9, 75, 10, "Spear")

	/* Instance Characters */
	Hero := NewCharacter("Hero", 20, sword)
	Enemy := NewCharacter("Enemy", 20, axe)

	// Trying to fight
	Hero.Attack(Enemy)
	Enemy.Attack(Hero)
}
