package domain

type Square struct {
	X        int    `json:"x"`
	Y        int    `json:"y"`
	IsEmmpty bool   `json:"is_empty"`
	Piece    *Piece `json:"piece"`
}
