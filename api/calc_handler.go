package api

import (
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

type Input struct {
	NumberA     	float64 `json:"number_a"`
	NumberB     	float64 `json:"number_b"`
	Operation     	string `json:"operation"`

}

type Output struct{
	Result string `json:"result"`
}

/*func GetResult(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var intput Input
	_ = json.NewDecoder(r.Body).Decode(&intput)

	output := calculator(intput)

	json.NewEncoder(w).Encode(output)

}*/

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error){
	 if request.HTTPMethod == "POST"{

		var intput Input

		json.Unmarshal([]byte(request.Body) , &intput)

		if validErrs := intput.validateOperation(); validErrs!="" {
			err := errors.New(validErrs)
			body, _ := json.Marshal(Output{})
			ApiResponse := events.APIGatewayProxyResponse{Body: string(body), StatusCode: http.StatusBadRequest}
			return ApiResponse , err

		}else {

			output := calculator(intput)
			body, _ := json.Marshal(output)

			ApiResponse := events.APIGatewayProxyResponse{Body: string(body), StatusCode: http.StatusOK}
			return ApiResponse, nil
		}


	} else {
		err := errors.New("Method Not Allowed!")
		ApiResponse := events.APIGatewayProxyResponse{Body: "", StatusCode: http.StatusMethodNotAllowed}
		return ApiResponse, err
	}
}

func (requestInput *Input) validateOperation() string {
	var err string

	if !(requestInput.Operation == "ADD" || requestInput.Operation == "SUBTRACT" || requestInput.Operation == "MULTIPLY" || requestInput.Operation == "DIVIDE") {
		err = "Accepted operation are ADD, SUBTRACT, MULTIPLY and DIVIDE"
	}

	if requestInput.Operation == "DIVIDE" && requestInput.NumberB == 0 {
		err = "Number B can not be zero"
	}

	return err
}
