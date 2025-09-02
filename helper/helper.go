package helper

import (
	"api-client/api"
	"api-client/config"
	"api-client/dto"
	"context"
	"fmt"
)

func SubmitOrder(o dto.Order) error {
	c := api.NewAPIClient(config.Config{
		BaseUrl: "http://localhost:8085",
		Dump:    true,
	})

	res, err := c.SubmitOrder(o).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Order created")
	fmt.Printf("%v\n", res)
	return nil
}
