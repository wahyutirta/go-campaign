package payment

import (
	"context"
	"fmt"
	"gocampaign/user"
	"reflect"
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type service struct {
}

type Service interface {
	GetPayment(transaction Transaction, user user.User) (string, error)
}

func NewService() *service {
	return &service{}
}

var s snap.Client

// func setupGlobalMidtransConfig() {
// 	midtrans.ServerKey = "SB-Mid-server-0men3mQGRfApcGptDcNo573B"
// 	midtrans.Environment = midtrans.Sandbox

// 	// Optional : here is how if you want to set append payment notification globally
// 	// midtrans.SetPaymentAppendNotification("https://example.com/append")
// 	// Optional : here is how if you want to set override payment notification globally
// 	// midtrans.SetPaymentOverrideNotification("https://example.com/override")

// 	//// remove the comment bellow, in cases you need to change the default for Log Level
// 	// midtrans.DefaultLoggerLevel = &midtrans.LoggerImplementation{
// 	//	 LogLevel: midtrans.LogInfo,
// 	// }
// }

// func createTransactionWithGlobalConfig() {
// 	res, err := snap.CreateTransactionWithMap(example.SnapParamWithMap())
// 	if err != nil {
// 		fmt.Println("Snap Request Error", err.GetMessage())
// 	}
// 	fmt.Println("Snap response", res)
// }

func initializeSnapClient() {
	s.New("SB-Mid-server-0men3mQGRfApcGptDcNo573B", midtrans.Sandbox)
}

func GenerateSnapReq(transaction Transaction, user user.User) *snap.Request {

	// Initiate Customer address
	custAddress := &midtrans.CustomerAddress{
		FName:       user.Name,
		LName:       "Doe",
		Phone:       "081234567890",
		Address:     "Temporary Address",
		City:        "Jakarta",
		Postcode:    "16000",
		CountryCode: "IDN",
	}

	// Initiate Snap Request
	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: transaction.Amount,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName:    user.Name,
			LName:    "Doe",
			Email:    user.Email,
			Phone:    "081234567890",
			BillAddr: custAddress,
			ShipAddr: custAddress,
		},
		EnabledPayments: snap.AllSnapPaymentType,
		Items: &[]midtrans.ItemDetails{
			{
				ID:    strconv.Itoa(transaction.ID),
				Price: transaction.Amount,
				Qty:   1,
				Name:  "Crows Funding",
			},
		},
	}
	return snapReq
}

func createTransaction(snapReq *snap.Request) {
	// Optional : here is how if you want to set append payment notification for this request
	// s.Options.SetPaymentAppendNotification("https://example.com/append")

	// Optional : here is how if you want to set override payment notification for this request
	// s.Options.SetPaymentOverrideNotification("https://example.com/override")
	// Send request to Midtrans Snap API

	resp, err := s.CreateTransaction(snapReq)
	if err != nil {
		fmt.Println("Error :", err.GetMessage())
	}
	fmt.Println("Response : ", resp)
	fmt.Println(reflect.TypeOf(resp))

}

func createTokenTransactionWithGateway(snapReq *snap.Request) string {
	// s.Options.SetPaymentOverrideNotification("https://example.com/url2")

	resp, err := s.CreateTransactionToken(snapReq)
	if err != nil {
		fmt.Println("Error :", err.GetMessage())
	}
	// fmt.Println("Response : ", resp)
	// fmt.Println(reflect.TypeOf(resp))
	return resp
}

func createUrlTransactionWithGateway(snapReq *snap.Request) string {
	s.Options.SetContext(context.Background())

	resp, err := s.CreateTransactionUrl(snapReq)
	if err != nil {
		fmt.Println("Error :", err.GetMessage())
	}
	// fmt.Println("Response : ", resp)
	// fmt.Println(reflect.TypeOf(resp))
	return resp
}

func (ser *service) GetPayment(transaction Transaction, user user.User) (string, error) {
	snapReq := GenerateSnapReq(transaction, user)
	initializeSnapClient()

	// fmt.Println("================== create transaction")
	// createTransaction(snapReq)

	// fmt.Println("================== create token transaction with gateway")
	// paymentToken := createTokenTransactionWithGateway(snapReq)
	// fmt.Println("Token : ", paymentToken)
	// fmt.Println(reflect.TypeOf(paymentToken))

	fmt.Println("================== create url transaction")
	paymentURL := createUrlTransactionWithGateway(snapReq)
	fmt.Println("RedirectURL : ", paymentURL)
	// fmt.Println(reflect.TypeOf(paymentURL))

	return paymentURL, nil
}
