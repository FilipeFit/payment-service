package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/filipeFit/payment-service/client/restclient"
	"github.com/filipeFit/payment-service/config"
	"github.com/filipeFit/payment-service/domain/api"
	"io"
	"io/ioutil"
	"log"
)

func ChangeAccountBalance(request *api.PaymentRequest) (*api.PaymentResponse, error) {
	url := fmt.Sprintf("%s%s", config.Config.AccountServiceUrl, "/balances")

	response, err := restclient.Post(url, request, nil)
	if err != nil {
		log.Println(fmt.Sprintf("error in performing the payment in the account service %s", err.Error()))
		return nil, err
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(fmt.Sprintf("error in read the account service response %s", err.Error()))
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(response.Body)

	if response.StatusCode != 201 {
		log.Println("the response of the account service is an error")
		var errorResponse api.ErrorResponse
		if err := json.Unmarshal(bytes, &errorResponse); err != nil {
			return nil, errors.New(
				fmt.Sprintf("Unknown error in calling the accoutns service %s", err.Error()))
		}
		return nil, errors.New(errorResponse.Message)
	}

	var result api.PaymentResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Println(fmt.Sprintf("error in read the account service response body%s", err.Error()))
		return nil, err
	}

	return &result, nil
}
