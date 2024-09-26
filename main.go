package main

import "fmt"

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

//func f(st string, wg *sync.WaitGroup, ch *chan string) {
//	defer wg.Done()
//	*ch <- st
//}

type changeConfig interface {
	alterAdUnit() data
	alterEcpm() data
}

type data struct {
	adUnit string
	ecpm   string
}

func (c data) alterAdUnit() data {
	c.adUnit = "/2345678/new-adunit"
	return c
}
func (c data) alterEcpm() data {
	c.ecpm = "23"
	return c
}

func alteration(c changeConfig) {
	fmt.Println(c.alterAdUnit())
	fmt.Println(c.alterEcpm())
}

func main() {
	r := data{adUnit: "/54678909876/old-adunit", ecpm: "21"}
	alteration(r)
	//var ns []string
	//s := make([]string, 1)
	//s = []string{"a", "b", "c", "d"}
	//copy(ns, s)
	//ns = s
	//fmt.Println("ns", ns)
	//fmt.Println("length", len(s), "capacity", cap(s))
	//requests := make(chan int, 5)
	//for i := 1; i <= 5; i++ {
	//	requests <- i
	//}
	//close(requests)
	//fmt.Println("Starting server...", &requests)
	//var wg sync.WaitGroup
	//
	//c1 := make(chan string, 2)
	//wg.Add(1)
	//f("first chan", &wg, &c1)
	//wg.Add(1)
	//time.Sleep(1 * time.Second)
	//f("sec chan", &wg, &c1)
	//
	//wg.Wait()
	//close(c1)
	//for ch := range c1 {
	//	fmt.Println(ch)
	//}
	//c2 := make(chan string)

	//go func() {
	//	time.Sleep(time.Second)
	//	c1 <- "first chan"
	//}()
	//go func() {
	//	time.Sleep(time.Second)
	//	c2 <- "second chan"
	//}()
	//for i := 0; i < 2; i++ {
	//
	//}
	//var wg sync.WaitGroup
	//
	//for i := 0; i < 3; i++ {
	//	wg.Add(1)
	//	//time.Sleep(time.Second)
	//	go func() {
	//		//time.Sleep(time.Second)
	//		f(i, &wg)
	//	}()
	//}
	//wg.Wait()
	//
	//rp := &r
	//fmt.Println("area", rp.area())
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
