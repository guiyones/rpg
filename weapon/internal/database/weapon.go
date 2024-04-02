package database

import (
	"database/sql"
)

type Weapon struct {
	db          *sql.DB
	ID          int32
	Name        string
	Distance    string
	WeaponRange string
	Hand        string
	Category    string
	Damage      int32
	Defense     int32
}

func NewWeapon(db *sql.DB) *Weapon {
	return &Weapon{
		db: db,
	}
}

func (w *Weapon) CreateWeapon(name string, distance string, category string, weaponRange string, hand string, damage int32, defense int32) (Weapon, error) {

	_, err := w.db.Exec("INSERT INTO weapons (name, distance, category, weaponRange, hand, damage, defense) VALUES (?, ?, ?, ?, ?, ?, ?)",
		name, distance, category, weaponRange, hand, damage, defense)

	if err != nil {
		return Weapon{}, err
	}

	return Weapon{
		ID:          w.ID,
		Name:        name,
		Distance:    distance,
		WeaponRange: weaponRange,
		Hand:        hand,
		Category:    category,
		Damage:      damage,
		Defense:     defense,
	}, nil
}

func (w *Weapon) GetAllWeapon() ([]Weapon, error) {

	rows, err := w.db.Query("SELECT id, name, distance, category, weaponRange, hand, damage, defense FROM weapons")
	if err != nil {
		return nil, err
	}

	weapons := []Weapon{}
	for rows.Next() {

		var id, damage, defense int32
		var name, category, distance, weaponRange, hand string

		if err := rows.Scan(&id, &name, &distance, &category, &weaponRange, &hand, &damage, &defense); err != nil {
			return nil, err
		}

		weapons = append(weapons, Weapon{
			ID:          id,
			Name:        name,
			Distance:    distance,
			WeaponRange: weaponRange,
			Hand:        hand,
			Category:    category,
			Damage:      damage,
			Defense:     defense,
		})
	}

	return weapons, nil

}

func (w *Weapon) GetByCategory(category string) ([]Weapon, error) {

	rows, err := w.db.Query("SELECT id, name, distance, category, weaponRange, hand, damage, defense FROM weapons WHERE category = ?", category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	weapons := []Weapon{}

	for rows.Next() {
		var id, damage, defense int32
		var name, category, distance, weaponRange, hand string

		if err := rows.Scan(&id, &name, &distance, &category, &weaponRange, &hand, &damage, &defense); err != nil {
			return nil, err
		}

		weapons = append(weapons, Weapon{
			ID:          id,
			Name:        name,
			Distance:    distance,
			WeaponRange: weaponRange,
			Hand:        hand,
			Category:    category,
			Damage:      damage,
			Defense:     defense,
		})
	}

	return weapons, nil
}
