package elastic

import (
	"context"
	"encoding/json"
	"log"

	"go-elastic/models"

	elastic "github.com/olivere/elastic/v7"
)

func GetESClient() *elastic.Client {

	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetBasicAuth("elastic", "tharun"),
		elastic.SetHealthcheck(false))

	if err != nil {
		log.Fatal("error in connection elastic search : ", err)
	}

	return client

}

func Insert(esclient *elastic.Client) {
	ctx := context.Background()

	//creating student object
	student1 := models.Student{
		ID:           "1",
		Name:         "Gopher doe",
		Age:          10,
		AverageScore: 99.9,
	}

	student2 := models.Student{
		ID:           "2",
		Name:         "Arjun",
		Age:          10,
		AverageScore: 99.9,
	}
	var students []models.Student
	students = append(students, student1)
	students = append(students, student2)
	for _, student := range students {
		dataJSON, err := json.Marshal(student)
		if err != nil {
			panic(err)
		}
		js := string(dataJSON)
		_, err = esclient.Index().
			Index("students").
			Id(student.ID).
			BodyJson(js).
			Do(ctx)

		if err != nil {
			panic(err)
		}
	}

	log.Println("[Elastic][InsertStudent]Insertion Successful")
}

func SearchWithName(esclient *elastic.Client, name string) {
	ctx := context.Background()

	var students []models.Student

	query := elastic.NewBoolQuery()
	query = query.Must(elastic.NewMatchQuery("name", name))

	/* this block will basically print out the es query */
	queryStr, err1 := query.Source()
	queryJs, err2 := json.Marshal(queryStr)

	if err1 != nil || err2 != nil {
		log.Println("[esclient][GetResponse]err during query marshal=", err1, err2)
	}
	log.Println("[esclient]Final ESQuery=\n", string(queryJs))
	/* until this block */

	searchService := esclient.Search().Index("students").Query(query)

	searchResult, err := searchService.Do(ctx)
	if err != nil {
		log.Println("[ProductsES][GetPIds]Error=", err)
		return
	}

	for _, hit := range searchResult.Hits.Hits {
		var student models.Student
		err := json.Unmarshal(hit.Source, &student)
		if err != nil {
			log.Println("[Getting Students][Unmarshal] Err=", err)
		}

		students = append(students, student)
	}

	if err != nil {
		log.Println("Fetching student fail: ", err)
	} else {
		for _, s := range students {
			log.Printf("Student found Name: %s, Age: %d, Score: %f \n", s.Name, s.Age, s.AverageScore)
		}
	}
}

func SearchWithOpt(esclient *elastic.Client, serachOpt ...models.SearchOpt) {
	ctx := context.Background()

	var students []models.Student

	query := elastic.NewBoolQuery()
	for _, opt := range serachOpt {

		if value, ok := opt.Value.(string); ok {
			query = query.Must(elastic.NewMatchQuery(opt.Key, value))
		} else if value, ok := opt.Value.([]interface{}); ok {
			query.MinimumNumberShouldMatch(1)
			for _, optItem := range value {
				query = query.Should(elastic.NewMatchQuery(opt.Key, optItem))
			}
		}
	}

	/* this block will basically print out the es query */
	queryStr, err1 := query.Source()
	queryJs, err2 := json.Marshal(queryStr)

	if err1 != nil || err2 != nil {
		log.Println("[esclient][GetResponse]err during query marshal=", err1, err2)
	}
	log.Println("[esclient]Final ESQuery=\n", string(queryJs))
	/* until this block */

	searchService := esclient.Search().Index("students").Query(query)

	searchResult, err := searchService.Do(ctx)
	if err != nil {
		log.Println("[ProductsES][GetPIds]Error=", err)
		return
	}

	for _, hit := range searchResult.Hits.Hits {
		var student models.Student
		err := json.Unmarshal(hit.Source, &student)
		if err != nil {
			log.Println("[Getting Students][Unmarshal] Err=", err)
		}

		students = append(students, student)
	}

	if err != nil {
		log.Println("Fetching student fail: ", err)
	} else {
		for _, s := range students {
			log.Printf("Student found Name: %s, Age: %d, Score: %f \n", s.Name, s.Age, s.AverageScore)
		}
	}
}

func UpSertEntireDoc(esclient *elastic.Client, id string, student models.Student) {
	ctx := context.Background()

	_, err := esclient.Update().Index("students").Id(id).Upsert(student).Doc(student).Do(ctx)

	if err != nil {
		panic(err)
	}

	log.Println("[Elastic][UpdateByID with particular id]Update Successful")
}

func UpdateByID(esclient *elastic.Client, id string, age int64, averagseScore float64) {
	ctx := context.Background()

	updateScript := elastic.NewScriptInline("ctx._source.age = params.new_age;ctx._source.average_score = params.new_average_score").
		Param("new_age", age).
		Param("new_average_score", averagseScore)

	_, err := esclient.Update().Index("students").Id(id).Script(updateScript).Do(ctx)

	if err != nil {
		panic(err)
	}

	log.Println("[Elastic][UpdateByID with particular id]Update Successful")
}

func UpdateByQuery(esclient *elastic.Client, name string, age int64, averagseScore float64) {
	ctx := context.Background()

	query := elastic.NewBoolQuery()
	query = query.Must(elastic.NewMatchQuery("name", name))

	updateScript := elastic.NewScriptInline("ctx._source.age = params.new_age;ctx._source.average_score = params.new_average_score").
		Param("new_age", age).
		Param("new_average_score", averagseScore)

	_, err := esclient.UpdateByQuery().Query(query).Index("students").Script(updateScript).Do(ctx)

	if err != nil {
		panic(err)
	}

	log.Println("[Elastic][UpdateByQuery]Update Successful")
}

func DeleteByID(esclient *elastic.Client, id string) {
	ctx := context.Background()

	_, err := esclient.Delete().Index("students").Id(id).Do(ctx)

	if err != nil {
		panic(err)
	}

	log.Println("[Elastic][DeleteStudent with particular id]Deletion Successful")
}

func DeleteIndex(esclient *elastic.Client, index string) {
	ctx := context.Background()

	_, err := esclient.DeleteIndex(index).Do(ctx)

	if err != nil {
		panic(err)
	}

	log.Println("[Elastic][DeleteStudent index]Deletion Successful")
}
