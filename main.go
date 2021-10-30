package main

import (
	"go-elastic/elastic"
)

func main() {
	esclient := elastic.GetESClient()

	elastic.Insert(esclient)

	// elastic.SearchWithName(esclient, "Arjun")

	// var searchOpts []models.SearchOpt
	// var searchOpt models.SearchOpt
	// searchOpt.Key = "name"
	// searchOpt.Value = "Arjun"
	// searchOpts = append(searchOpts, searchOpt)
	// searchOpt.Key = "name"
	// searchOpt.Value = "Gopher"
	// searchOpts = append(searchOpts, searchOpt)
	// elastic.SearchWithOpt(esclient, searchOpts...)

	// student := models.Student{
	// 	ID:   "1",
	// 	Name: "Tharunqqq",
	// }
	// elastic.UpSertEntireDoc(esclient, "1", student)
	// elastic.UpdateByID(esclient, "1", 100, 77.99)

	//Please make sure refresh is happened before update more resarch required here
	//https://www.elastic.co/guide/en/elasticsearch/reference/7.0/indices-refresh.html
	elastic.UpdateByQuery(esclient, "Arjun", 1002, 77.99)

	//elastic.DeleteByID(esclient, "1")
	//elastic.DeleteIndex(esclient, "students")
}
