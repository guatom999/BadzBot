package main

import (
	"github.com/guatom999/BadzBot/configs"
	"github.com/guatom999/BadzBot/modules/server"
)

func main() {
	cfg := configs.NewConfig("./.env")

	server.NewDiscordServer(cfg).Start()

	// resp, err := http.Get("https://catfact.ninja/fact")

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// //We Read the response body on the line below.
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// _ = body
	// //Convert the body to type string
	// sb := string(body)
	// fmt.Println(sb)
}

// func testRequest() {

// 	resp, err := http.Get("https://api.publicapis.org/entries")

// 	if err != nil {
// 		fmt.Println("Get error")
// 	}

// 	fmt.Println(resp)
// }

// func initDB() {
// 	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
// 		viper.GetString("db.username"),
// 		viper.GetString("db.password"),
// 		viper.GetString("db.host"),
// 		viper.GetString("db.port"),
// 		viper.GetString("db.table"),
// 	)

// 	db, err := sqlx.Open(viper.GetString("db.driver"), dsn)
// 	if err != nil {
// 		panic(err)
// 	}

// 	_ = db

// 	resp, err := http.Get("https://api.publicapis.org/entries")

// 	if err != nil {
// 		fmt.Println("Get error")
// 	}

// 	fmt.Println(resp)

// }
