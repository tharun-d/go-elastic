package main

import (
	"go-elastic/elastic"
	"go-elastic/models"
)

func main() {
	esclient := elastic.GetESClient()

	elastic.Insert(esclient)

	elastic.SearchWithName(esclient, "Arjun")

	var searchOpts []models.SearchOpt
	var searchOpt models.SearchOpt
	searchOpt.Key = "name"
	searchOpt.Value = "Arjun"
	searchOpts = append(searchOpts, searchOpt)
	searchOpt.Key = "name"
	searchOpt.Value = "Gopher"
	searchOpts = append(searchOpts, searchOpt)
	elastic.SearchWithOpt(esclient, searchOpts...)

	student := models.Student{
		ID:   "1",
		Name: "Tharunqqq",
	}
	elastic.UpSertEntireDoc(esclient, "1", student)
	elastic.UpdateByID(esclient, "1", 100, 77.99)

	elastic.DeleteByID(esclient, "1")
	elastic.DeleteIndex(esclient, "students")
}
