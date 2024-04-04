package database

import "database/sql"

type Attributes struct {
	db            *sql.DB
	ID            int32
	Name          string
	AttributeType string
	Description   string
	Value         int32
}

func NewAttributes(db *sql.DB) *Attributes {
	return &Attributes{
		db: db,
	}
}

func (a *Attributes) CreateAttibutes(name string, attributeType string, description string, value int32) (Attributes, error) {

	_, err := a.db.Exec("INSERT INTO attributes (name, attributeType, description, value) VALUES (?,?,?,?)", name, attributeType, description, value)
	if err != nil {
		return Attributes{}, err
	}

	return Attributes{
		Name:          name,
		AttributeType: attributeType,
		Description:   description,
		Value:         value,
	}, nil

}

func (a *Attributes) GetAllAtributes() ([]Attributes, error) {

	rows, err := a.db.Query("SELECT id, name, attributeType, description, value FROM attributes")
	if err != nil {
		return nil, err
	}

	attributes := []Attributes{}

	for rows.Next() {
		var id, value int32
		var name, attributeType, description string

		if err := rows.Scan(&id, &name, &attributeType, &description, &value); err != nil {
			return nil, err
		}

		attributes = append(attributes, Attributes{
			ID:            id,
			Name:          name,
			AttributeType: attributeType,
			Description:   description,
			Value:         value,
		})
	}

	return attributes, nil
}

func (a *Attributes) GetByAttributeType(attributeType string) ([]Attributes, error) {

	rows, err := a.db.Query("SELECT id, name, attributeType, description, value FROM attributes WHERE attributeType = ?", attributeType)
	if err != nil {
		return nil, err
	}

	attributes := []Attributes{}

	for rows.Next() {
		var id, value int32
		var name, attributeType, description string

		if err := rows.Scan(&id, &name, &attributeType, &description, &value); err != nil {
			return nil, err
		}

		attributes = append(attributes, Attributes{
			ID:            id,
			Name:          name,
			AttributeType: attributeType,
			Description:   description,
			Value:         value,
		})
	}

	return attributes, nil
}
