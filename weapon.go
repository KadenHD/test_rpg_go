package main

import (
	"fmt"
	"math/rand"
)

type Weapon struct {
	Name     string
	Power    int
	Accuracy int
	Critical int
	Category string
}

func NewWeapon(name string, power int, accuracy int, critical int, category string) *Weapon {
	return &Weapon{
		Name:     name,
		Power:    power,
		Accuracy: accuracy,
		Critical: critical,
		Category: category,
	}
}

func (w *Weapon) AdvantageAgainst(o *Weapon) bool {
	switch w.Category {
	case "Sword":
		return o.Category == "Axe"
	case "Axe":
		return o.Category == "Spear"
	case "Spear":
		return o.Category == "Sword"
	default:
		return false
	}
}

func (w *Weapon) CriticalStrike() bool {
	random := rand.Intn(100)
	if w.Critical >= random {
		fmt.Printf("%s does a critical strike with a roll of %d/%d.\n", w.Name, random, w.Critical)
		return true
	}
	return false
}

func (w *Weapon) MissAttack() bool {
	random := rand.Intn(100)
	if w.Accuracy <= random {
		fmt.Printf("%s missed with a roll of %d/%d.\n", w.Name, random, w.Accuracy)
		return true
	}
	return false
}

func (w *Weapon) CalculateDamage(o *Weapon) float32 {
	if !w.MissAttack() {
		damage := float32(w.Power)
		if w.AdvantageAgainst(o) {
			fmt.Printf("%s had advantage against %s.\n", w.Name, o.Name)
			damage *= 1.2
		}
		if w.CriticalStrike() {
			damage *= 1.8
		}
		return damage
	}
	return 0
}
