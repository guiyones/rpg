package database

import "database/sql"

type Shield struct {
	db      *sql.DB
	ID      int32
	Name    string
	Defense int32
}

func NewShield(db *sql.DB) *Shield {
	return &Shield{
		db: db,
	}
}

func (s *Shield) CreateShield(name string, defense int32) (*Shield, error) {

	_, err := s.db.Exec("INSERT INTO shields (name, defense) VALUES (?, ?)",
		name, defense)

	if err != nil {
		return nil, err
	}

	return &Shield{
		ID:      s.ID,
		Name:    name,
		Defense: defense,
	}, nil
}

func (s *Shield) GetAllShields() ([]Shield, error) {

	rows, err := s.db.Query("SELECT id, name, defense FROM shields")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	shields := []Shield{}

	for rows.Next() {
		var id, defense int32
		var name string

		if err := rows.Scan(&id, &name, &defense); err != nil {
			return nil, err
		}

		shields = append(shields, Shield{
			ID:      id,
			Name:    name,
			Defense: defense,
		})
	}

	return shields, nil
}
