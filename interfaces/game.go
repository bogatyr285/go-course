package main

import "fmt"

type Character interface {
	Attack(opponent Character)
	Defend(damage int)
	GetHealth() int
	Name() string
}

type Warrior struct {
	health      int
	attackPower int
	name        string
}

func (w *Warrior) Attack(opponent Character) {
	fmt.Println("Warrior attacks!")
	opponent.Defend(w.attackPower)
}
func (w *Warrior) Defend(damage int) {
	fmt.Println("Warrior defends and takes damage!")
	w.health -= damage / 2 // Warriors take half damage due to armor
}
func (w *Warrior) GetHealth() int {
	return w.health
}
func (w *Warrior) Name() string {
	return w.name
}

type Mage struct {
	health     int
	spellPower int
	name       string
}

func (m *Mage) Attack(opponent Character) {
	fmt.Println("Mage casts a spell!")
	opponent.Defend(m.spellPower)
}
func (m *Mage) Defend(damage int) {
	fmt.Println("Mage defendsand takes damage!")
	m.health -= damage // Mages take full damage
}
func (m *Mage) GetHealth() int {
	return m.health
}
func (w *Mage) Name() string {
	return w.name
}

type Rogue struct {
	health    int
	stabPower int
	name      string
}

func (r *Rogue) Attack(opponent Character) {
	fmt.Println("Rogue strikes from the shadows!")
	opponent.Defend(r.stabPower * 2) // Rogues deal double damage on attack
}
func (r *Rogue) Defend(damage int) {
	fmt.Println("Rogue defends and takes damage!")
	r.health -= damage // Rogues take full damage
}
func (r *Rogue) GetHealth() int {
	return r.health
}
func (w *Rogue) Name() string {
	return w.name
}
func battle(c1, c2 Character) {
	fmt.Println("Battle Start!")
	for c1.GetHealth() > 0 && c2.GetHealth() > 0 {
		c1.Attack(c2)
		fmt.Printf("Character %s Health: %d\n", c1.Name(), c1.GetHealth())
		fmt.Printf("Character %s Health: %d\n", c2.Name(), c2.GetHealth())
		if c2.GetHealth() <= 0 {
			fmt.Printf("Character %s has fallen!\n", c2.Name())
			break
		}
		c2.Attack(c1)
		fmt.Printf("Character %s Health: %d\n", c1.Name(), c1.GetHealth())
		fmt.Printf("Character %s Health: %d\n", c2.Name(), c2.GetHealth())
		if c1.GetHealth() <= 0 {
			fmt.Printf("Character %s has fallen!\n", c1.Name())
			break
		}
	}
	fmt.Println("Battle End!")
}

func main() {
	warrior := &Warrior{
		health:      100,
		attackPower: 10,
		name:        "war",
	}
	mage := &Mage{
		health:     80,
		spellPower: 15,
		name:       "mage",
	}
	rogue := &Rogue{
		health:    70,
		stabPower: 20,
		name:      "rogue",
	}
	fmt.Println("Warrior vs Mage")
	battle(warrior, mage)
	warrior.health = 100 // Reset health for another battle
	mage.health = 80
	fmt.Println("Mage vs Rogue")
	battle(mage, rogue)
	rogue.health = 70
	fmt.Println("Warrior vs Rogue")
	battle(warrior, rogue)
}
