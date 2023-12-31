package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"sync"

	"github.com/GRRashid/lamoda"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type ProductPostgres struct {
	db *sqlx.DB
	mx sync.Mutex
}

func NewProductPostgres(db *sqlx.DB) *ProductPostgres {
	return &ProductPostgres{
		db: db,
		mx: sync.Mutex{},
	}
}

func (p *ProductPostgres) Create(product lamoda.RawProduct) (int, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createProductQuery := fmt.Sprintf("INSERT INTO %s (size, storage_id, count, name, status) VALUES ($1, $2, $3, $4, $5) RETURNING id", productsTable)
	row := tx.QueryRow(createProductQuery, product.Size, product.StorageId, product.Count, product.Name, "available")
	err = row.Scan(&id)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return 0, fmt.Errorf("Rollback error:", rollbackErr)
		}

		return 0, fmt.Errorf("Insert error:", err)
	}

	createProductStoragesQuery := fmt.Sprintf("INSERT INTO %s (storage_id, product_id) VALUES ($1, $2)", productStoragesTable)
	_, err = tx.Exec(createProductStoragesQuery, product.StorageId, id)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return 0, fmt.Errorf("Rollback error:", rollbackErr)
		}

		return 0, fmt.Errorf("Insert error:", err)
	}

	return id, tx.Commit()
}

func (p *ProductPostgres) FindStorage(storageId int) (int, error) {
	var id int
	selectStorageQuery := fmt.Sprintf("SELECT id FROM %s WHERE id = $1", storagesTable)
	err := p.db.QueryRow(selectStorageQuery, storageId).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("The storage with this id does not exist")
		}

		return 0, fmt.Errorf("Select error:", err)
	}

	return id, nil
}

func (p *ProductPostgres) FindAvailableProducts(storageId int) ([]lamoda.Product, error) {
	var products []lamoda.Product
	query := fmt.Sprintf(
		"SELECT pt.id, pt.size, pt.storage_id, pt.count, pt.name, pt.status "+
			"FROM %s pt INNER JOIN %s pst on pt.id = pst.product_id "+
			"WHERE pst.storage_id = $1 AND pt.status = 'available'",
		productsTable, productStoragesTable)

	err := p.db.Select(&products, query, storageId)
	if err != nil {
		return nil, fmt.Errorf("Select error", err)
	}

	return products, nil
}

func (p *ProductPostgres) ReservedProduct(updatableProducts []int) error {
	p.mx.Lock()
	defer p.mx.Unlock()

	query := fmt.Sprintf("UPDATE %s SET status = 'reserved' WHERE id = ANY($1)", productsTable)
	_, err := p.db.Exec(query, pq.Array(updatableProducts))
	if err != nil {
		return fmt.Errorf("Update error:", err)
	}

	return nil
}

func (p *ProductPostgres) UnreservedProduct(ids []int) error {
	p.mx.Lock()
	defer p.mx.Unlock()

	query := fmt.Sprintf("UPDATE %s SET status = 'available' WHERE id = ANY($1)", productsTable)
	_, err := p.db.Exec(query, pq.Array(ids))
	if err != nil {
		return fmt.Errorf("Update error:", err)
	}

	return nil
}

func (p *ProductPostgres) FindUpdatableProducts(ids []int, status string) ([]int, error) {
	selectProductQuery := fmt.Sprintf("SELECT id, status FROM %s WHERE id = ANY ($1) AND status = $2", productsTable)
	rows, err := p.db.Query(selectProductQuery, pq.Array(ids), status)
	if err != nil {
		return nil, fmt.Errorf("Select error:", err)
	}

	var updatableProducts []int
	for rows.Next() {
		var product lamoda.Product
		if err = rows.Scan(&product.ID, &product.Status); err != nil {
			return nil, fmt.Errorf("Error scanning product:", err)
		}

		updatableProducts = append(updatableProducts, product.ID)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterating over product rows:", err)
	}

	return updatableProducts, nil
}
