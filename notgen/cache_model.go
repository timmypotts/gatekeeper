package main

import(
	"database/sql"
)

type ofacCache struct {
	GetcoinsID		uint		`json:"getcoinsID"`
	Name			string		`json:"name"`
	Score			int			`json:"score"`
	Status			string		`json:"status"`
	UpdatedAt		time.Time	`json:"updated_at"` // figure out datatype
	CreatedAt		time.Time	`json:"created_at"` // figure out datatype
}

func (u *user) getUserByGetcoinsID(db *sql.DB) error {
	return db.QueryRow("SELECT firstname, lastname, getcoinsID, score, status FROM users WHERE getcoinsID=$1", u.GetcoinsID)
	.Scan(&u.FirstName, &u.LastName, &u.GetcoinsID, &u.Score, &u.Status)
}

func (u *user) cacheBlockscoreResponse(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO users(firstname, lastname, ssn_hash, score, status, blockscoreID) VALUES($1, $2, $3, $4, $5) RETURNING  ",
		u.FirstName, u.LastName, u.SSN_HASH, u.Score, u.Status, u.BlockscoreID).Scan(&p.GetcoinsID)

		if err != nil{
			return err
		}
	)
}

func (u *user) checkIfUserIsCached(db *sql.DB) error {
	return db.QueryRow("SELECT * FROM user WHERE getcoinsID=$1", GetcoinsID)
}

