package gaia
import "database/sql"

// 1st of rows
type Globs struct {
	Name string `json:"name"`
	En   string `json:"en"`
}

type Post struct {
	ID       int    `json:"id"`
	UID      int    `json:"uid"`
	Uri      sql.NullString `json:"uri"`
	Img      sql.NullString `json:"img"`
	Title    string `json:"title"`
	Subtitle sql.NullString `json:"subtitle"`
	Excerpt  sql.NullString `json:"excerpt"`
	Content  sql.NullString `json:"content"`
}

type User struct {
	ID        int    `json:"id"`
	Url       sql.NullString `json:"url"`
	Img       sql.NullString `json:"img"`
	Name      string `json:"name"`
	Pass      string `json:"pass"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// album represents data about a record album.
type Album struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
