package paymentxendit

import (
	"fmt"
	"gocampaign/entity"
	"gocampaign/user"
	"log"
	"strconv"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/client"
	"github.com/xendit/xendit-go/invoice"
)

type service struct {
}

type Service interface {
	GetPayment(transaction entity.Transaction, user user.User) (string, error)
}

func NewService() *service {
	return &service{}
}

func (ser *service) GetPayment(transaction entity.Transaction, user user.User) (string, error) {

	// Basic setup
	xenCli := client.New("xnd_development_hq7JLiie9KzU5HMsgrUN14JatgYDn0fLRhjgPA5X79ZTQoq6qwt3jrnDWDwigFN")

	// customerAddress := xendit.CustomerAddress{
	// 	Country:     "Indonesia",
	// 	StreetLine1: "Jalan Makan",
	// 	StreetLine2: "Kecamatan Kebayoran Baru",
	// 	City:        "Jakarta Selatan",
	// 	State:       "Daerah Khusus Ibukota Jakarta",
	// 	PostalCode:  "12345",
	// }

	customer := xendit.InvoiceCustomer{
		GivenNames:   "John",
		Email:        "johndoe@example.com",
		MobileNumber: "+6287774441111",
		Address:      "Temporary Address",
	}

	fee := xendit.InvoiceFee{
		Type:  "ADMIN WAHYUT",
		Value: 50000,
	}

	fees := []xendit.InvoiceFee{fee}

	data := invoice.CreateParams{
		ExternalID:  strconv.Itoa(transaction.ID),
		Amount:      float64(transaction.Amount),
		PayerEmail:  user.Email,
		Description: "invoice #1",
		Fees:        fees,
		Customer:    customer,
	}

	resp, err := xenCli.Invoice.Create(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("invoice url: %+v\n", resp.InvoiceURL)

	return resp.InvoiceURL, nil
}
