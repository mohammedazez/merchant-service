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
		password)
	VALUES(?,?,?,?)
	`

	LoginUser = `
  	SELECT * FROM users 
	WHERE email = ?
  	`

	FindUserById = `
	SELECT * FROM users;
	SELECT * FROM outlets WHERE user_id IN (?);
    ORDER BY created_at DESC
  	`
)
