package main

import (
	"semay.com/manager"
)

//	@title			Swagger Blue API
//	@version		0.1
//	@description	This is User Admin server.
//	@termsOfService	http://swagger.io/terms/
//  @BasePath  /api/v1

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						X-APP-TOKEN
//	@description				Description for what is this security definition being used

//	@securityDefinitions.apikey Refresh
//	@in							header
//	@name						X-REFRESH-TOKEN
//	@description				Description for what is this security definition being used

func main() {

	manager.Execute()
}
