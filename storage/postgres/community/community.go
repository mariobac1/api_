package community

const (
	MigrateCommunity = ` CREATE TABLE IF NOT EXIST communities(
		id Serial NOT NULL
		name VARCHAR(50) NOT NULL,
		created_at TIMESTAMP NOT NULL DAFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT communities_id_pk PRIMARY KEY (id),
		)`
	CreateCommunity = `INSERT INTO communities (name, created_at)
		VALUES($1, $2) RETURNING id`
	GetAllCommunity = `SELECT id, name, created_at, updated_at
		FROM communities`
	GetByIDCommunity = GetAllCommunity + `WHERE id = $1`
	UpdateCommunity  = `UPDATE communities SET name = $1, updated_at = $2, WHERE id = $3`
)
