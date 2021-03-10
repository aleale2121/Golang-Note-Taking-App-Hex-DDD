package initiator

import (
	"fmt"
	"github.com/aleale2121/Golang-TODO-Hex-DDD/internal/constant/model"
	"github.com/aleale2121/Golang-TODO-Hex-DDD/internal/glue/routing"
	"github.com/aleale2121/Golang-TODO-Hex-DDD/internal/handler/rest"
	note "github.com/aleale2121/Golang-TODO-Hex-DDD/internal/module/user"
	"github.com/aleale2121/Golang-TODO-Hex-DDD/internal/storage/postgress"
	"github.com/aleale2121/Golang-TODO-Hex-DDD/platform/routers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/lib/pq"
)

const (
	postgresURL = "postgres://%s:%s@%s/%s?sslmode=disable"

	dialect = "postgres"
)

func createTable(dbConn *gorm.DB) []error {
	dbConn.Debug().DropTableIfExists(&model.Note{})
	errs := dbConn.Debug().CreateTable(&model.Note{}).GetErrors()
	if errs != nil {
		return errs
	}
	return nil
}
func User(testInit bool) {
	//dbUser := os.Getenv("DB_USER")
	//dbPass := os.Getenv("DB_PASS")
	//dbHost := os.Getenv("DB_HOST")
	//dbName := os.Getenv("DB_NAME")
	dbURL := fmt.Sprintf(postgresURL, "postgres", "root", "localhost", "Note")

	dbConn, err := gorm.Open(dialect, dbURL)
	if dbConn != nil {
		defer dbConn.Close()
	}
	if err != nil {
		panic(err)
	}
	//createTable(dbConn)
	postgresUser := postgress.NewNoteRepository(dbConn)
	useCase := note.NewService(*postgresUser)
	handler := rest.NewNoteHandler(useCase)
	router := routing.NoteRouting(handler)

	//host := os.Getenv("HOST")
	//port := os.Getenv("HOST_PORT")
	server := routers.NewRouting("localhost", "8080", router)
	if testInit {
		fmt.Println("Initialize test mode Finished!")
		//os.Exit(0)
	}

	server.Serve()
}
