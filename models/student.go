package models

type Student struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Age          int64   `json:"age"`
	AverageScore float64 `json:"average_score"`
}

type SearchOpt struct {
	Key   string
	Value interface{}
}
