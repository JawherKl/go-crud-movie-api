package main
 
import (
   "example.com/go-crud-api/db"
   "example.com/go-crud-api/router"
)
 
func main() {
   db.InitPostgresDB()
   router.InitRouter().Run()
}
