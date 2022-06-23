package transaction

import (
	"errors"
	"gocampaign/campaign"
	"gocampaign/paymentmidtrans"
	"gocampaign/paymentstripe"
	"gocampaign/paymentxendit"
	"strconv"
)

type service struct {
	repository             Repository
	campaignRepository     campaign.Repository
	paymentServiceMidtrans paymentmidtrans.Service
	paymentServiceXendit   paymentxendit.Service
	paymentServiceStripe   paymentstripe.Service
}

type Service interface {
	GetTransactionByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
	GetTransactionByUserID(userID int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
	ProcessPayment(input TransactionNotificationInput) error
}

func NewService(repository Repository, campaignRepository campaign.Repository, paymentServiceMidtrans paymentmidtrans.Service, paymentServiceXendit paymentxendit.Service, paymentServiceStripe paymentstripe.Service) *service {
	return &service{repository, campaignRepository, paymentServiceMidtrans, paymentServiceXendit, paymentServiceStripe}
}

func (s *service) GetTransactionByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error) {

	campaign, err := s.campaignRepository.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("not the owner of the campaign")
	}

	transactions, err := s.repository.GetByCampaignID(input.ID)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
func (s *service) GetTransactionByUserID(userID int) ([]Transaction, error) {
	transactions, err := s.repository.GetByUserID(userID)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	transaction := Transaction{
		CampaignID: input.CampaignID,
		Amount:     input.Amount,
		UserID:     input.User.ID,
		Status:     "pending",
	}

	newTransaction, err := s.repository.CreateTransaction(transaction)
	if err != nil {
		return newTransaction, err
	}

	if input.PaymentProvider == "midtrans" {
		paymentTransaction := paymentmidtrans.Transaction{
			ID:     newTransaction.ID,
			Amount: int64(newTransaction.Amount),
		}

		paymentURL, err := s.paymentServiceMidtrans.GetPayment(paymentTransaction, input.User)

		if err != nil {
			return newTransaction, err
		}

		newTransaction.PaymentURL = paymentURL

		newTransaction, err = s.repository.UpdateTransaction(newTransaction)
		if err != nil {
			return newTransaction, err
		}

	} else if input.PaymentProvider == "xendit" {
		paymentTransaction := paymentxendit.Transaction{
			ID:     transaction.ID,
			Amount: int64(transaction.Amount),
		}

		paymentURL, err := s.paymentServiceXendit.GetPayment(paymentTransaction, input.User)

		if err != nil {
			return transaction, err
		}

		transaction.PaymentURL = paymentURL

		transaction, err = s.repository.UpdateTransaction(newTransaction)
		if err != nil {
			return newTransaction, err
		}

		newTransaction.PaymentURL = paymentURL

		newTransaction, err = s.repository.UpdateTransaction(newTransaction)
		if err != nil {
			return newTransaction, err
		}
	} else if input.PaymentProvider == "stripe" {
		paymentTransaction := paymentstripe.Transaction{
			ID:     transaction.ID,
			Amount: int64(transaction.Amount),
		}

		paymentURL, err := s.paymentServiceStripe.GetPayment(paymentTransaction, input.User)

		if err != nil {
			return transaction, err
		}

		transaction.PaymentURL = paymentURL

		transaction, err = s.repository.UpdateTransaction(newTransaction)
		if err != nil {
			return newTransaction, err
		}

		newTransaction.PaymentURL = paymentURL

		newTransaction, err = s.repository.UpdateTransaction(newTransaction)
		if err != nil {
			return newTransaction, err
		}
	}

	return newTransaction, nil
}

func (s *service) ProcessPayment(input TransactionNotificationInput) error {
	transaction_id, _ := strconv.Atoi(input.OrderID)
	transaction, err := s.repository.GetByID(transaction_id)
	if err != nil {
		return err
	}
	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}
	updatedTransaction, err := s.repository.UpdateTransaction(transaction)
	if err != nil {
		return err
	}
	campaign, err := s.campaignRepository.FindByID(updatedTransaction.CampaignID)
	if err != nil {
		return err
	}

	if updatedTransaction.Status == "paid" {
		campaign.BackerCount = campaign.BackerCount + 1
		campaign.CurrentAmount = campaign.CurrentAmount + updatedTransaction.Amount
		_, err := s.campaignRepository.UpdateCampaign(campaign)
		if err != nil {
			return err
		}
	}
	return nil
}
