package repo

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          string  `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImageURL    string  `json:"image_url" db:"image_url"`
	CreatedAt   string  `json:"created_at" db:"created_at"`
	UpdatedAt   string  `json:"updated_at" db:"updated_at"`
}

type ProductRepo interface {
	Create(product Product) (*Product, error)
	Get(productID string) (*Product, error)
	List() ([]*Product, error)
	Delete(productID string) error
	Update(pr Product) (*Product, error)
}

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepo {

	repo := &productRepo{db: db}
	return repo
}

func (p *productRepo) Create(product Product) (*Product, error) {

	query := `INSERT INTO products (title, description, price, image_url) 
	VALUES (:title, :description, :price, :image_url)
	 RETURNING id , created_at, updated_at;`

	rows, err := p.db.NamedQuery(query, product)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		err = rows.Scan(&product.ID, &product.CreatedAt, &product.UpdatedAt)
	}

	if product.ID != "" {
		return &product, nil
	}

	return &product, nil
}

func (p *productRepo) Get(productID string) (*Product, error) {

	query := `SELECT *
	FROM products WHERE id=$1`

	var product Product
	err := p.db.Get(&product, query, productID)
	if err != nil {
		return nil, errors.New("product not found")
	}

	return &product, nil
}

func (p *productRepo) List() ([]*Product, error) {
	query := `SELECT * FROM products`

	var products []*Product
	err := p.db.Select(&products, query)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *productRepo) Delete(productID string) error {
	query := `DELETE FROM products WHERE id=$1`
	result, err := p.db.Exec(query, productID)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("product not found")
	}

	return nil
}

func (p *productRepo) Update(pr Product) (*Product, error) {
	query := `
		UPDATE products 
		SET 
			title = :title,
			description = :description,
			price = :price,
			image_url = :image_url,
			updated_at = NOW()
		WHERE id = :id
		RETURNING id, created_at, updated_at;
	`

	rows, err := p.db.NamedQuery(query, pr)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		rows.Scan(&pr.ID, &pr.CreatedAt, &pr.UpdatedAt)
		return &pr, nil
	}

	return nil, errors.New("product not found")
}

