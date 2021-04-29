package main

import (
	"fmt"
	"github.com/Rhymond/go-money"
	"log"
	"time"
)

type Cart struct {
	ID string
	CreatedDate time.Time
	TotalPrice *money.Money
	Items []Item
}

func (c *Cart) ComputePrice() error {
	// TODO handle other currencies
	grandTotal := money.New(0,"EUR")
	var err error
	for _, v := range c.Items {
		totalItem := v.UnitPrice.Multiply(int64(v.Quantity))
		grandTotal, err = grandTotal.Add(totalItem)
		if err != nil {
			return err
		}
	}
	c.TotalPrice = grandTotal
	return nil
}

type Item struct {
	ID string
	Name string
	SKU string
	Quantity uint8
	UnitPrice *money.Money
}

func main() {
	items := []Item{
		{
			ID: "42",
			Name:"Tea Pot",
			SKU:"tea-pot",
			Quantity: 2,
			UnitPrice: money.New(122,"EUR"),
		},
		{
			ID: "43",
			Name:"Book",
			SKU:"Book",
			Quantity: 10,
			UnitPrice: money.New(69000,"EUR"),
		},
	}
	myCart := Cart{
		ID:          "4785",
		CreatedDate: time.Now(),
		Items: items,
	}
	for _,v := range myCart.Items {
		fmt.Println(v.Name,v.Quantity)
	}
	err := myCart.ComputePrice()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("TOTAL PRICE for cart : ", myCart.TotalPrice.Display())
}
