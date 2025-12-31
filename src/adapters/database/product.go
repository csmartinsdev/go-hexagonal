package database

import (
	"database/sql"
	"go-hexagonal/src/application"

	_ "github.com/mattn/go-sqlite3"
)

type ProductDatabase struct {
	database *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDatabase {
	return &ProductDatabase{database: db}
}

func (p *ProductDatabase) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	stmt, err := p.database.Prepare("select id, name, price, status from products where id=?")
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductDatabase) create(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := p.database.Prepare(`insert into products(id, name, price, status) values(?,?,?,?)`)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(
		product.GetID(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
	)
	if err != nil {
		return nil, err
	}
	err = stmt.Close()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductDatabase) update(product application.ProductInterface) (application.ProductInterface, error) {
	_, err := p.database.Exec("update products set name = ?, price=?, status=? where id = ?",
		product.GetName(), product.GetPrice(), product.GetStatus(), product.GetID())
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductDatabase) Save(product application.ProductInterface) (application.ProductInterface, error) {
	_, err := p.Get(product.GetID())
	if err != nil {
		if err == sql.ErrNoRows {
			return p.create(product)
		}
		// if Get returned any other error, propagate it
		return nil, err
	}
	return p.update(product)
}
