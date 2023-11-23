package lamoda

type Product struct {
	ID        int
	Name      string `json:"name" db:"name"`
	Size      int    `json:"size" db:"size"`
	Count     int    `json:"count" db:"count"`
	StorageId int    `json:"storage" db:"storage_id"`
	Status    string `json:"status"  db:"status"`
}

type RawProduct struct {
	Name      string `json:"name" db:"name"`
	Size      int    `json:"size" db:"size"`
	Count     int    `json:"count" db:"count"`
	StorageId int    `json:"storage" db:"storage_id"`
	Status    string `json:"status"  db:"status"`
}

type ProductIds struct {
	IDs []int `json:"ids"`
}
