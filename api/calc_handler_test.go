package api

import (
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"

	//"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*func TestGetResult(t *testing.T) {

	input := Input{
		NumberA:5,
		NumberB:4,
		Operation:"MULTIPLY",
	}

	body, _ := json.Marshal(input)

	req, err := http.NewRequest("POST", "/result", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	hf := http.HandlerFunc(GetResult)
	hf.ServeHTTP(response, req)



	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}


	var actualOutput Output
	_ = json.NewDecoder(response.Body).Decode(&actualOutput)


	expectedOutput := Output{
		Result:"20.00",
	}

	if actualOutput != expectedOutput {
		t.Errorf("handler returned unexpected body: got %v want %v", actualOutput, expectedOutput)
	}

}*/

func TestHandleRequest(t *testing.T) {

	validInput := Input{
		NumberA:5,
		NumberB:4,
		Operation:"MULTIPLY",
	}
	validBody, _ := json.Marshal(validInput)
	expect:=  Output{
		Result:"20.00",
	}

	invalidInput := Input{
		NumberA:5,
		NumberB:4,
		Operation:"MULTIPL",
	}

	invalidBody, _ := json.Marshal(invalidInput)

	tests := []struct {
		request events.APIGatewayProxyRequest
		expect  Output
		err     error
	}{
		{
			request: events.APIGatewayProxyRequest{HTTPMethod: "POST",Body: string(validBody)},
			expect:  expect,
			err:     nil,
		},
		{
			request: events.APIGatewayProxyRequest{HTTPMethod: "POST",Body: string(invalidBody)},
			expect:  Output{},
			err:     errors.New("Accepted operation are ADD, SUBTRACT, MULTIPLY and DIVIDE"),
		},
	}

	for _, test := range tests {

		response, err := HandleRequest(test.request)
		var output Output
		json.Unmarshal([]byte(response.Body),&output)

		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, output)
	}
}

