package lamoda

type Storage struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Accessibility string `json:"accessibility"`
}

type RawStorage struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Accessibility string `json:"accessibility"`
}
