package main

import "fmt"

type Character struct {
	Name   string
	HP     int
	Weapon *Weapon
}

func NewCharacter(name string, hp int, weapon *Weapon) *Character {
	return &Character{
		Name:   name,
		HP:     hp,
		Weapon: weapon,
	}
}

func (c *Character) Attack(target *Character) {
	fmt.Println(c.Name)
	fmt.Print("+--------------------------------------------+\n")
	fmt.Printf("%s attacks %s with %s.\n", c.Name, target.Name, c.Weapon.Name)
	damage := c.Weapon.CalculateDamage(target.Weapon)
	target.ReceiveDamage(int(damage))
	fmt.Print("+--------------------------------------------+\n\n")
}

func (c *Character) ReceiveDamage(damage int) {
	fmt.Printf("%s receives %d damage.\n", c.Name, damage)
	c.HP -= damage
	fmt.Printf("%s's remaining HP: %d.\n", c.Name, c.HP)
}
