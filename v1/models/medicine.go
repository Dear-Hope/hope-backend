package models

import "gorm.io/gorm"

type Medicine struct {
	gorm.Model
	Kind           Kind   `json:"kind" gorm:"size:3"`
	Name           string `json:"name" gorm:"unique"`
	PriceRange     string `json:"price_range"`
	Description    string `json:"description"`
	Benefits       string `json:"benefits"`
	Subcategory    string `json:"subcategory"`
	Composition    string `json:"composition"`
	Dosage         string `json:"dosage"`
	Serving        string `json:"serving"`
	StoreProcedure string `json:"store_procedure"`
	Consideration  string `json:"consideration"`
	SideEffect     string `json:"side_effect"`
	Package        string `json:"package"`
	Producer       string `json:"producer"`
	MedicineGroup  string `json:"medicine_group"`
	Image          string `json:"image"`
}

type Kind string

var codeToNameKindMap = map[Kind]Kind{
	"AL":  "Alergi",
	"AS":  "Asma",
	"DE":  "Demam",
	"DI":  "Diabetes",
	"FB":  "Flu dan Batuk",
	"OL":  "Obat Luka",
	"SP":  "Saluran Pencernaan",
	"VS":  "Vitamin Suplemen",
	"HI":  "Hipertensi",
	"MT":  "Mulut dan Tenggorokan",
	"JK":  "Jantung dan Kolesterol",
	"TSO": "Tulang, Sendi, dan Otot",
	"DT":  "Diet",
	"MA":  "Mata",
	"KU":  "Kulit",
}

func (Medicine) TableName() string {
	return "medicine"
}

func (ths *Medicine) ConvertKindToName() {
	ths.Kind = codeToNameKindMap[ths.Kind]
}

type MedicineService interface {
	List(Kind) ([]*Medicine, error)
}

type MedicineRepository interface {
	GetAll(Medicine) ([]*Medicine, error)
}
