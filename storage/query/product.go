package query

const (
	QueryCreateProduct = `
	INSERT INTO products(id,
		product_name,
		price,
		sku,
		picture,
		created_at,
		updated_at,
		outlet_id
	)
	VALUES(?,?,?,?,?,?,?,?)
	`

	QueryFindAllProduct = `
	SELECT
      	id,
      	product_name,
		price,
		sku,
		picture,
		outlet_id
    FROM
    	products
    ORDER BY created_at DESC
	`

	QueryFindProductById = `
	SELECT 
		id,
      	product_name,
		price,
		sku,
		picture,
		outlet_id 
	FROM 
		products 
	WHERE 
		id = ?
	`

	QueryUpdateProductByID = `
	UPDATE 
		products 
	SET 
		product_name = ?,
		price = ?,
		sku = ?,
		picture = ?,
		outlet_id = ?
	WHERE 
		id = ?
	`

	QueryDeleteProductById = `
	DELETE FROM 
		products 
	WHERE 
		id = ?
	`
)
