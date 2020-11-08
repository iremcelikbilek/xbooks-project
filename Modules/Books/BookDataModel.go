package Books

type AddBookDataModel struct {
	BookName      string `json:"bookName"`
	Author        string `json:"author"`
	Explanation   string `json:"explanation"`
	BookImageURL  string `json:"bookImageURL"`
	BookDetailURL string `json:"bookDetailURL"`
	BookURL       string `json:"bookURL"`
	BookCost      int    `json:"bookCost"`
}
