package user

//go:generate mkdir -p ./restful
//go:generate go run github.com/go-swagger/go-swagger/cmd/swagger generate server -t restful -f swagger.yml --exclude-main
