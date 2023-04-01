package query

const (
	AddBook = `
		INSERT INTO
			books
		(
			title,
			author,
			des
			
		)
		VALUES ($1, $2, $3)
		RETURNING *;
	`
)
