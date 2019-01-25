package storage

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
	"users.git/config"
)

// InitializeDB needs to be abstracted to work with any kind of storage system
func InitializeDB(conf config.Conf) *gorm.DB {

	/* 	Depending on what is provided in the config, we should allow different initializations of different databases
		and a general "storage" interface to allow ease of changing between different storage options such as postgresql,
		inmemory, cassandra, etc.
		TODO: Get a full list of possible dbs to use
	*/
	switch db := strings.ToLower(conf.DBType); db {
	case "postgresql":
		break
	case "mysql":
		break
	case "cassandra":
		break
	case "inmemory":
		break
	default:
		//in memory
		break
	}

	port, _ := strconv.Atoi(conf.DBPort)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		conf.DBHost, port, conf.DBUser, conf.DBName, conf.DBPassword)

	fmt.Println(psqlInfo)
	db, _ := gorm.Open(conf.DBType, psqlInfo)
	//if err != nil {
	//	fmt.Printf("Failed to open gorm: %q\n", err)
	//	return nil
	//}
	fmt.Println("Successfully connected to database!")
	return db
}