package requests

type CreatePaymentRequest struct {
	Amount         float64        `json:"amount"`
	Currency       string         `json:"currency"`
	Description    string         `json:"description"`
	Language       string         `json:"language"`
	Test           bool           `json:"test"`
	BillingAddress BillingAddress `json:"billing_address"`
	CreditCard     CreditCard     `json:"credit_card"`
	Customer       Customer       `json:"customer"`
}

type BillingAddress struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Country   string `json:"country"`
	City      string `json:"city"`
	State     string `json:"state"`
	Zip       string `json:"zip"`
	Address   string `json:"address"`
}

type CreditCard struct {
	Number            string `json:"number"`
	VerificationValue string `json:"verification_value"`
	Holder            string `json:"holder"`
	ExpMonth          string `json:"exp_month"`
	ExpYear           string `json:"exp_year"`
}

type Customer struct {
	Email string `json:"email"`
}
