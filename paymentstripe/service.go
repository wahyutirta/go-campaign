package paymentstripe

import (
	"fmt"
	"gocampaign/entity"
	"gocampaign/user"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/bankaccount"
	"github.com/stripe/stripe-go/v72/customer"
	"github.com/stripe/stripe-go/v72/paymentintent"
	"github.com/stripe/stripe-go/v72/token"
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

	stripe.Key = "sk_test_51LAr8zJT2q3qqnqADiNTfKo1oVTkK0GKnm8pgEv8oPb5kM3r6VcOPBu14oLVB2bYHkWTiUa7B0BOxXne12girWCY00JjzAo76t"
	params := &stripe.CustomerParams{
		Description:   stripe.String("My First Test Customer (created for API docs at https://www.stripe.com/docs/api)"),
		Email:         stripe.String(user.Email),
		Name:          stripe.String(user.Name),
		PaymentMethod: stripe.String("pm_card_visa"),
		InvoiceSettings: &stripe.CustomerInvoiceSettingsParams{
			DefaultPaymentMethod: stripe.String("pm_card_visa"),
		},
		PreferredLocales: stripe.StringSlice([]string{"en", "es"}),
	}
	c, _ := customer.New(params)

	params = &stripe.CustomerParams{
		InvoiceSettings: &stripe.CustomerInvoiceSettingsParams{
			CustomFields: []*stripe.CustomerInvoiceCustomFieldParams{
				{
					Name:  stripe.String("VAT"),
					Value: stripe.String("123ABC"),
				},
			},
		},
	}

	c, _ = customer.Update(
		*stripe.String(c.ID),
		params,
	)

	fmt.Println(c)
	fmt.Println(*c.InvoiceSettings.CustomFields[0].Name)
	fmt.Println(*c.InvoiceSettings.CustomFields[0].Value)

	t, _ := token.New(&stripe.TokenParams{
		BankAccount: &stripe.BankAccountParams{
			Country:           stripe.String("US"),
			Currency:          stripe.String(string(stripe.CurrencyUSD)),
			AccountHolderName: stripe.String("Jenny Rosen"),
			AccountHolderType: stripe.String(string(stripe.BankAccountAccountHolderTypeIndividual)),
			RoutingNumber:     stripe.String("110000000"),
			AccountNumber:     stripe.String("000123456789"),
		},
	})

	params_bank := &stripe.BankAccountParams{
		Customer: stripe.String(c.ID),
		Token:    stripe.String(t.ID),
	}
	ba, _ := bankaccount.New(params_bank)

	fmt.Println(ba)
	fmt.Println(ba.ID)

	params_payment := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(transaction.Amount),
		Currency: stripe.String(string(stripe.CurrencyIDR)),
	}

	pi, _ := paymentintent.New(params_payment)

	fmt.Println(pi)
	fmt.Println(pi.ID)

	params_confirm := &stripe.PaymentIntentConfirmParams{
		PaymentMethod: stripe.String("pm_card_visa"),
	}

	pi, _ = paymentintent.Confirm(*stripe.String(pi.ID), params_confirm)

	fmt.Println(pi)
	fmt.Println(pi.ID)
	fmt.Println(pi.Charges.URL)
	baseURL := "https://api.stripe.com"
	fmt.Println(baseURL)
	fmt.Printf("%s%s", baseURL, pi.Charges.URL)

	return "success", nil
}
