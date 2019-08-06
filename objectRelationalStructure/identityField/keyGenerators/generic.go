package keygenerators

import "database/sql"

// KeyPlaceholder is the default key for newly created but not inserted
const KeyPlaceholder = -1

// KeyGenerator manages ids for domain objects or row data gateways
type KeyGenerator struct {
	db         *sql.DB
	keyName    string
	startValue int
}

// NewKeyGenerator creates new KeyGenerator
func NewKeyGenerator(db *sql.DB, keyName string, startValue int) *KeyGenerator {
	return &KeyGenerator{db, keyName, startValue}
}

// NextID returns the next id and increment it by 1
func (generator *KeyGenerator) NextID() (int, error) {
	db := generator.db

	// start transaction
	tx, err := db.Begin()
	if err != nil {
		return KeyPlaceholder, err
	}
	// defer rollback
	defer func() {
		_ = tx.Rollback()
	}()

	// get next id
	row := tx.QueryRow("select nextid from keys where name = $1", generator.keyName)
	var id int
	if err := row.Scan(&id); err != nil {
		return KeyPlaceholder, err
	}

	// update next id in database
	_, err = tx.Exec("update keys set nextid = $1 where name = $2", id+1, generator.keyName)
	if err != nil {
		return KeyPlaceholder, err
	}

	// commit and return
	if err := tx.Commit(); err != nil {
		return KeyPlaceholder, err
	}
	return id, nil
}

// Reset resets id in database
func (generator *KeyGenerator) Reset() error {
	db := generator.db

	// reset id in database
	_, err := db.Exec("update keys set nextid = $1 where name = $2", generator.startValue, generator.keyName)
	if err != nil {
		return err
	}
	return nil
}
