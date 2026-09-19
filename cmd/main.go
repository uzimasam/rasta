package main

func main(){
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		panic(err)
	}
}