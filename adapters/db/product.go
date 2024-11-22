package db

import (
	"database/sql"

	"github.com/gaspecian/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	stmt, err := p.db.Prepare("select id, name, price, status from products where id=?")
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductDb) Save(productInterface application.ProductInterface) (application.ProductInterface, error) {
	var rows int
	p.db.QueryRow("select id from products where id=?", productInterface.GetId()).Scan(&rows)
	if rows == 0 {
		_, err := p.create(productInterface)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := p.update(productInterface)
		if err != nil {
			return nil, err
		}
	}
	return productInterface, nil
}

func (p *ProductDb) create(productInterface application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := p.db.Prepare(`insert into products(id, name, price, status) values (?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(
		productInterface.GetId(),
		productInterface.GetName(),
		productInterface.GetPrice(),
		productInterface.GetStatus(),
	)

	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}
	return productInterface, nil
}

func (p *ProductDb) update(productInterface application.ProductInterface) (application.ProductInterface, error) {
	_, err := p.db.Exec("update products set name = ?, price= ?, status= ? where id= ?",
		productInterface.GetName(),
		productInterface.GetPrice(),
		productInterface.GetStatus(),
		productInterface.GetId(),
	)
	if err != nil {
		return nil, err
	}
	return productInterface, nil
}
