package lib

import(
	"upper.io/db.v3/mysql"
	"upper.io/db.v3"
)

var settings = mysql.ConnectionURL{
	Host: "%HOST%",
	User: "%USER%",
	Password: "%PASSWORD%",
	Database: "%DATABASE%",
}

var Sess = db.Database

func init() {
	var err error

	Sess, err = mysql.Open(settings)

	if err != nil {
		log.Fatal(err.Error())
	}
}