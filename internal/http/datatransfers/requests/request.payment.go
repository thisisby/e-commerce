package requests

type CreatePaymentRequest struct {
	Amount         float64        `json:"amount"`
	Currency       string         `json:"currency"`
	Description    string         `json:"description"`
	Language       string         `json:"language"`
	ReturnUrl      string         `json:"return_url"`
	Test           bool           `json:"test"`
	BillingAddress BillingAddress `json:"billing_address"`
	CreditCard     CreditCard     `json:"credit_card"`
	Customer       Customer       `json:"customer"`
	AdditionalData AdditionalData `json:"additional_data"`
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
	Number                       string `json:"number"`
	VerificationValue            string `json:"verification_value"`
	Holder                       string `json:"holder"`
	ExpMonth                     string `json:"exp_month"`
	ExpYear                      string `json:"exp_year"`
	SkipThreeDSecureVerification bool   `json:"skip_three_d_secure_verification"`
}

type Customer struct {
	Email string `json:"email"`
}

type AdditionalData struct {
	Browser Browser `json:"browser"`
}

type Browser struct {
	AcceptHeader      string `json:"accept_header"`
	ScreenWidth       int    `json:"screen_width"`
	WindowHeight      int    `json:"window_height"`
	ScreenColorDepth  int    `json:"screen_color_depth"`
	WindowWidth       int    `json:"window_width"`
	JavaEnabled       bool   `json:"java_enabled"`
	JavascriptEnabled bool   `json:"javascript_enabled"`
	Language          string `json:"language"`
	ScreenHeight      int    `json:"screen_height"`
	TimeZone          int    `json:"time_zone"`
	UserAgent         string `json:"user_agent"`
	TimeZoneName      string `json:"time_zone_name"`
}
