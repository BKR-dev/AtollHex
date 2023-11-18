package logic

import (
	"fmt"
	"local/db"
	"math/rand"
	"time"
)

func GetCurrentTime() time.Time {
	return time.Now().UTC()
}

func GetArticles() dataObject {
	allArticles := getLotsOfData()
	return allArticles[0]
}

type dataObject struct {
	article       string
	articleNumber string
	fabricType    string
	fabricColor   string
}

func getLotsOfData() []dataObject {
	var obj dataObject
	var data []dataObject

	for i := 1; i <= 1_000_000; i++ {
		data = append(data, populateData(obj))
	}

	return data
}

func populateData(obj dataObject) dataObject {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	article := []string{"tshirt", "pants", "jacket"}
	articleNum := []string{"1001", "2002", "3003"}
	fabric := []string{"wool", "denim", "linnen"}
	colors := []string{"red", "blue", "green", "yellow"}

	obj.article = article[r.Intn(len(article))]
	obj.articleNumber = articleNum[r.Intn(len(article))]
	obj.fabricType = fabric[r.Intn(len(article))]
	obj.fabricColor = colors[r.Intn(len(article))]

	return obj
}

func writeTestDataToDb() {
	db := db.NewDBEngine()
	err := db.Connect()
	if err != nil {
		fmt.Printf("error connecting to db: %s\n", err)
	}
	defer db.Close()

	for i := 1; i <= 1_000_000; i++ {
		article := []string{"tshirt", "pants", "jacket"}
		articleNum := []string{"1001", "2002", "3003"}
		fabric := []string{"wool", "denim", "linnen"}
		colors := []string{"red", "blue", "green", "yellow"}

		articleType := article[rand.Intn(len(article))]
		articleNumber := articleNum[rand.Intn(len(article))]
		fabricType := fabric[rand.Intn(len(article))]
		fabricColor := colors[rand.Intn(len(article))]

		_, err := db.Query("INSERT INTO articles (article, article_number, fabric_type, fabric_color) VALUES ($1, $2, $3, $4)", articleType, articleNumber, fabricType, fabricColor)
		if err != nil {
			panic(err)
		}
	}
}
