// Nome do pacote é o nome da pasta em que ele está inserido
package beer

// Inicial maiúscula para visibilidade ser pública
type Beer struct {
	ID    int64     `json:"id"` // Tags
	Name  string    `json:"name"`
	Type  BeerType  `json:"type"`
	Style BeerStyle `json:"style"`
}

// Irá receber o valor de algum tipo de cerveja, que foi definido posteriormente
type BeerType int

// Vamos definir os tipos de cerveja
const (
	TypeAle   = 1
	TypeLager = 2
	TypeMalt  = 3
	TypeStout = 4
)

// Função que está atrelada à um tipo
// var x TypeAle
// fmt.Println(x.String())
func (t BeerType) String() string {
	switch t {
	case TypeAle:
		return "Ale"
	case TypeLager:
		return "Lager"
	case TypeMalt:
		return "Malt"
	case TypeStout:
		return "Stout"
	}
	return "Unknown"

}

type BeerStyle int

// Ao usar o iota, o compilador vai definir os ids sequencialmente, iniciando de 1, automaticamente.
const (
	StyleAmber = iota + 1
	StyleBlonde
	StyleBrown
	StyleCream
	StyleDark
	StylePale
	StyleStrong
	StyleWheat
	StyleRed
	StyleIPA
	StyleLime
	StylePilsner
	StyleGolden
	StyleFruit
	StyleHoney
)

func (t BeerStyle) String() string {
	switch t {
	case StyleAmber:
		return "Amber"
	case StyleBlonde:
		return "Blonde"
	case StyleBrown:
		return "Brown"
	case StyleCream:
		return "Cream"
	case StyleDark:
		return "Dark"
	case StylePale:
		return "Pale"
	case StyleStrong:
		return "Strong"
	case StyleWheat:
		return "Wheat"
	case StyleRed:
		return "Red"
	case StyleIPA:
		return "IPA"
	case StyleLime:
		return "Lime"
	case StylePilsner:
		return "Pilsner"
	case StyleGolden:
		return "Golder"
	case StyleFruit:
		return "Fruit"
	case StyleHoney:
		return "Honey"
	}
	return "Unknown"

}
