package UserData

type UserDataModel struct {
	PersonEmail string           `json:"personEmail"`
	CurrentBook CurrentBookModel `json:"currentBook"`
	ReadedBook  ReadedBookModel  `json:"readedBook"`
}

type CurrentBookModel struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	BookURL       string `json:"bookURL"`
	ImageURL      string `json:"imageURL"`
	TotalDuration string `json:"totalDuration"`
	CurrentPage   int    `json:"currentPage"`
	TotalPoint    int    `json:"totalPoint"`
}

type ReadedBookModel struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	BookURL       string `json:"bookURL"`
	ImageURL      string `json:"imageURL"`
	TotalDuration string `json:"totalDuration"`
	TotalPoint    []int  `json:"totalPoint"`
}
