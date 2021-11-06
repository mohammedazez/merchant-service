package query

const (
	GetAllUsers = `
	SELECT
      	id,
      	full_name ,
      	email
    FROM
    	users
    ORDER BY created_at DESC
	`

	RegisterUser = `
  	INSERT INTO users(id,
		full_name,
		email,
		password,
		created_at,
		updated_at
	)
	VALUES(?,?,?,?,?,?)
	`

	LoginUser = `
  	SELECT * 
	FROM 
		users 
	WHERE
		email = ?
  	`

	FindUserById = `
	SELECT * 
	FROM 
		users 
	WHERE 
		id = ?
  	`

	DeleteUserById = `
	DELETE FROM 
		users 
	WHERE 
		id = ?
	`

	UpdateUserByID = `
	UPDATE 
		users 
	SET 
		full_name= ?, 
		email= ?, 
		password= ?,
		updated_at = ?
	WHERE 
		id = ?
	`

	FindOutletUserByID = `
	SELECT * 
	FROM 
		outlets 
	WHERE 
		id = ?
	`

	CreateOutletbyUser = `
	INSERT INTO outlets(id,
		outlet_name,
		picture,
		user_id,
		created_at,
		updated_at
	)
	VALUES(?,?,?,?,?,?)
	`

	GetAllOutlets = `
	SELECT
      	id,
      	outlet_name,
      	picture,
		user_id
    FROM
    	outlets
    ORDER BY created_at DESC
	`
)
