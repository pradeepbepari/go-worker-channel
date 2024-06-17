package model

type User struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Address  *Address `json:"address"`
	Phone    string   `json:"phone"`
	Website  string   `json:"website"`
	Company  *Company `json:"company"`
}
type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     *Geo   `json:"geo"`
}
type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}
type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}
type Post struct {
	UserId int    `json:"userid"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
type Comments struct {
	PostId int    `json:"postid"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}
type Albums struct {
	UserId int    `json:"userid"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
}
type Photos struct {
	AlbumId     int    `json:"albumid"`
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Url         string `json:"url"`
	ThumbnaiUrl string `json:"thumbnailurl"`
}
type Todo struct {
	UserId    int    `json:"userid"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
type CombaineStruct struct {
	User     *User
	Post     *[]Post
	Comments *[]Comments
	Albums   *Albums
	Photos   *Photos
	Todo     *Todo
}
