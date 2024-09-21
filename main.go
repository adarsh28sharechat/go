package main

import (
	"fmt"
	//_ "github.com/go-sql-driver/mysql"
)

//const correctPassword = "12345"

//func isAuthenticated() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		password := c.GetHeader("Password")
//
//		if password == correctPassword {
//			c.Next()
//		} else {
//			c.AbortWithStatus(http.StatusUnauthorized)
//		}
//	}
//}

//func createObjectTable(db *sql.DB) error {
//	query := `CREATE TABLE IF NOT EXISTS object(object_id int primary key auto_increment, object_type text,
//        side int)`
//	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancelfunc()
//	res, err := db.ExecContext(ctx, query)
//	if err != nil {
//		log.Printf("Error %s when creating product table", err)
//		return err
//	}
//	rows, err := res.RowsAffected()
//	if err != nil {
//		log.Printf("Error %s when getting rows affected", err)
//		return err
//	}
//	log.Printf("Rows affected when creating table: %d", rows)
//	return nil
//}

func oneval(i int) {
	i = 3
	fmt.Println("inside", &i)
}

func oneptr(i *int) {
	*i = 0
}

func main() {
	i := 1
	fmt.Println("initial", i)

	oneval(i)
	fmt.Println("oneval", i)

	oneptr(&i)
	fmt.Println("oneptr", i)

	fmt.Println("poinyer", &i)

	//cards := newDeck()
	//hand, remaining := deal(cards, 5)
	//hand.print()
	//remaining.print()
	//fmt.Println("hello world")
	//db, err := sql.Open("mysql", "root:Sharechat@28@tcp(127.0.0.1:3306)/Test")
	//
	//if err != nil {
	//	log.Fatal("Error connecting to the database--1:", err)
	//}
	//defer db.Close()
	//
	//// Test the connection
	//err = db.Ping()
	//
	//if err != nil {
	//	log.Fatal("Error connecting to the database--2:", err)
	//	return
	//}
	//
	//log.Println("Connected to the database successfully!")
	//
	///////////////////////////////
	//err = createObjectTable(db)
	//
	//if err != nil {
	//	log.Printf("Create product table failed with error %s", err)
	//	return
	//}
	//
	//fmt.Println("handling roures")
	//router := gin.Default()
	//router.Use(gzip.Gzip(gzip.DefaultCompression))
	//
	//router.Use(func(c *gin.Context) {
	//	c.Set("db", db)
	//	c.Next()
	//})
	//
	//v1 := router.Group("/calculator")
	//{
	//	v1.GET("/addition", isAuthenticated(), services.Addition)
	//	v1.GET("/subtract", isAuthenticated(), services.Subtraction)
	//	v1.GET("/multiply", isAuthenticated(), services.Multiplication)
	//	v1.GET("/divide", isAuthenticated(), services.Division)
	//}
	//
	////router.GET("/addition", isAuthenticated(), services.Addition)
	////router.GET("/subtract", isAuthenticated(), services.Subtraction)
	////router.GET("/multiply", isAuthenticated(), services.Multiplication)
	////router.GET("/divide", isAuthenticated(), services.Division)
	//
	//v2 := router.Group("/shape")
	//{
	//	v2.POST("/", isAuthenticated(), repository.AddObject)
	//	v2.PATCH("/:id", isAuthenticated(), repository.UpdateObject)
	//	v2.DELETE("/:id", isAuthenticated(), repository.RemoveObject)
	//	v2.GET("/:id", isAuthenticated(), repository.CalculateArea)
	//}
	////router.POST("/insert", isAuthenticated(), repository.AddObject)
	////router.PATCH("/update/:id", isAuthenticated(), repository.UpdateObject)
	////router.DELETE("/delete/:id", isAuthenticated(), repository.RemoveObject)
	////router.GET("/area/:id", isAuthenticated(), repository.CalculateArea)
	//
	//router.Run(":5000")
}
