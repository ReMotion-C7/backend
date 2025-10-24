package reset

import "ReMotion-C7/config"

func main() {

	config.ConnectDatabase()

	database := config.GetDatabase()

	ClearDatabase(database)
	
}
