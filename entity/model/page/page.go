package page

type Page struct {
	Header PageHeader
	Cells  []PageCell
}

func (ph *PageHeader) ToByteArray() []byte {

}
func (p *Page) ToByteArray() []byte {

}
