package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"myCalculator/api"
)


func main() {

	lambda.Start(api.HandleRequest)

	/*// init Router
	r:= mux.NewRouter()

	// Route Handlers / Endpoints
	r.HandleFunc("/result" , api.GetResult).Methods("POST")

	fmt.Println("Serving on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))*/


}






