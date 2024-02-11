package gaia

// 1st of rows
type Globs struct {
	Name string `json:"name"`
	En   string `json:"en"`
}

type Post struct {
	ID       int    `json:"id"`
	UID      int    `json:"uid"`
	Uri      string `json:"uri"`
	Img      string `json:"img"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Excerpt  string `json:"excerpt"`
	Content  string `json:"content"`
}

type User struct {
	ID        int    `json:"id"`
	UID       int    `json:"uid"`
	Url       string `json:"url"`
	Img       string `json:"img"`
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
