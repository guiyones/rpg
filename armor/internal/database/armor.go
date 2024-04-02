package database

import "database/sql"

type Armor struct {
	db       *sql.DB
	ID       int32
	Name     string
	Category string
	Defense  int32
}

func NewArmor(db *sql.DB) *Armor {
	return &Armor{
		db: db,
	}
}
func (a *Armor) CreateArmor(name string, category string, defense int32) (Armor, error) {

	_, err := a.db.Exec("INSERT INTO armors (name, category, defense) VALUES (?, ?, ?)", name, category, defense)

	if err != nil {
		return Armor{}, err
	}

	return Armor{
		Name:     name,
		Category: category,
		Defense:  defense,
	}, nil
}

func (a *Armor) GetAllArmors() ([]Armor, error) {

	rows, err := a.db.Query("SELECT id, name, category, defense FROM armors")
	if err != nil {
		return nil, err
	}

	armors := []Armor{}

	for rows.Next() {

		var id, defense int32
		var name, category string

		if err := rows.Scan(&id, &name, &category, &defense); err != nil {
			return nil, err
		}

		armors = append(armors, Armor{
			ID:       id,
			Name:     name,
			Category: category,
			Defense:  defense,
		})
	}

	return armors, nil
}

func (a *Armor) GetByCategory(category string) ([]Armor, error) {

	rows, err := a.db.Query("SELECT id, name, category, defense FROM armors WHERE category = ?", category)
	if err != nil {
		return nil, err
	}

	armors := []Armor{}

	for rows.Next() {

		var id, defense int32
		var name, category string

		if err := rows.Scan(&id, &name, &category, &defense); err != nil {
			return nil, err
		}

		armors = append(armors, Armor{
			ID:       id,
			Name:     name,
			Category: category,
			Defense:  defense,
		})
	}

	return armors, nil
}
