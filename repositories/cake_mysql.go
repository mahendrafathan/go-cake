package repositories

import (
	"database/sql"

	"github.com/mahendrafathan/go-cake/models"
)

type cakeRepoMysql struct {
	db *sql.DB
}

func (c *cakeRepoMysql) FindAll() (results []models.Cake, err error) {
	results = []models.Cake{}
	rows, err := c.db.Query("SELECT c.id, c.title, c.description, c.rating, c.image FROM cakes c order by c.rating, c.title")
	if err != nil {
		return
	}

	for rows.Next() {
		cake := models.Cake{}
		rows.Scan(
			&cake.Id,
			&cake.Title,
			&cake.Description,
			&cake.Rating,
			&cake.Image,
		)

		results = append(results, cake)
	}
	return results, nil
}
func (c *cakeRepoMysql) FindOne(id int) (result models.Cake, err error) {
	row := c.db.QueryRow("SELECT c.id, c.title, c.description, c.rating, c.image FROM cakes c where id = ?", id)
	if err = row.Scan(
		&result.Id,
		&result.Title,
		&result.Description,
		&result.Rating,
		&result.Image,
	); err != nil {
		return
	}
	return result, nil
}
func (c *cakeRepoMysql) Create(cake models.Cake) error {
	insertQuery := `INSERT INTO cakes (title, description, rating, image) VALUES (?, ?, ?, ?)`
	_, err := c.db.Exec(
		insertQuery,
		cake.Title,
		cake.Description,
		cake.Rating,
		cake.Image,
	)
	if err != nil {
		return err
	}

	return nil
}
func (c *cakeRepoMysql) Update(id int, cake models.Cake) error {
	updateQuery := `UPDATE cakes set title = ?, description = ?, rating = ?, image = ? where id = ?`
	_, err := c.db.Exec(
		updateQuery,
		cake.Title,
		cake.Description,
		cake.Rating,
		cake.Image,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}
func (c *cakeRepoMysql) Delete(id int) error {
	deleteQuery := `DELETE FROM cakes WHERE id = ?`
	_, err := c.db.Exec(
		deleteQuery,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}
