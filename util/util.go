package util

import (
	"fmt"
	"math/rand"

	"github.com/BKR-dev/AtollHex/db"
)

func PopluateDB() {
	db := db.NewDBEngine()
	err := db.Connect()
	if err != nil {
		fmt.Printf("error connecting to db: %s\n", err)
	}
	defer db.Close()

	createDbTableIfNotExist()
	createDummyDataIfTableEmpty()

}

func createDbTableIfNotExist() {
	db := db.NewDBEngine()
	err := db.Connect()
	if err != nil {
		fmt.Printf("error connecting to db: %s\n", err)
	}
	defer db.Close()

	_, err = db.Query("CREATE TABLE IF NOT EXISTS articles (id SERIAL PRIMARY KEY, article TEXT, article_number TEXT, fabric_type TEXT, fabric_color TEXT)")
	if err != nil {
		panic(err)
	}
}

func createDummyDataIfTableEmpty() {
	db := db.NewDBEngine()
	err := db.Connect()
	if err != nil {
		fmt.Printf("error connecting to db: %s\n", err)
	}
	defer db.Close()

	// fix: see if empty returns 0 rows

	result, err := db.Query("SELECT * FROM articles")
	if err != nil {
		fmt.Printf("error querying db: %s\n", err)
	}
	if len(result) == 0 {
		fmt.Println("creating dummy data")
		writeDummyData()
	}
}

type Article struct {
	Article       string
	ArticleNumber string
	FabricType    string
	FabricColor   string
}

func writeDummyData() {
	db := db.NewDBEngine()
	err := db.Connect()
	if err != nil {
		fmt.Printf("error connecting to db: %s\n", err)
	}
	defer db.Close()
	for i := 1; i <= 1_000; i++ {
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

func GetArticles() []Article {
	db := db.NewDBEngine()
	err := db.Connect()
	if err != nil {
		fmt.Printf("error connecting to db: %s\n", err)
	}
	defer db.Close()

	result, err := db.Query("SELECT * FROM articles")
	if err != nil {
		fmt.Printf("error querying db: %s\n", err)
	}

	articles := make([]Article, 0)
	for _, row := range result {
		articles = append(articles, Article{
			Article:       row["article"].(string),
			ArticleNumber: row["article_number"].(string),
			FabricType:    row["fabric_type"].(string),
			FabricColor:   row["fabric_color"].(string),
		})
	}
	return articles
}

func DropArticlesTable() {
	db := db.NewDBEngine()
	err := db.Connect()
	if err != nil {
		fmt.Printf("error connecting to db: %s\n", err)
	}
	defer db.Close()

	_, err = db.Query("DROP TABLE articles")
	if err != nil {
		panic(err)
	}
}
