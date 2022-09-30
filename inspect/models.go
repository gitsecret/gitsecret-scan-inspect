package inspect

type Finding struct {
	Hash      string `json:"hash,omitempty" gorm:"primary_key"`
	FromLine  int    `json:"from_line,omitempty"`
	ToLine    int    `json:"to_line,omitempty"`
	Content   string `json:"content,omitempty"`
	Type      string `json:"type,omitempty"`
	Validated bool   `json:"validated,omitempty"`
	Meta      string `json:"meta,omitempty"`
}
