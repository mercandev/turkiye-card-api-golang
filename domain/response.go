package domain

type Response struct {
	IdentityCard   []string `json:"identity_card"`
	Cards          []string `json:"cards"`
	Transactions   []string `json:"transactions"`
	Services       []string `json:"services"`
	EnabledService []string `json:"enabled_service"`
}

type Cards struct {
	CardType string          `json:"cardtype"`
	CardName string          `json:"cardname"`
	Bank     struct{ Banks } `json:"bank"`
}

type IdentityCard struct {
	IdentityNumber int32  `json:"identitynumber"`
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	DateOfBrith    string `json:"dateofbrith"`
	Gender         string `json:"gender"`
	Nationality    string `json:"nationality"`
	DrivingLicence string `json:"driving_licence"`
	IbanNumber     string `json:"iban_number"`
	PaymentMethod  string `json:"payment_method"`
}

type Transactions struct {
	Company         string  `json:"company"`
	TransactionDate string  `json:"transactions_date"`
	Balance         float64 `json:"balance"`
}

type Services struct {
	ServicesName    string  `json:"serices_name"`
	PaymentsService string  `json:"payments_service"`
	Balance         float32 `json:"balance"`
}

type EnabledService struct {
	SafeSecurePay  bool `json:"safe_secure_pay"`
	DrivingLicence bool `json:"driving_licence"`
	Passport       bool `json:"passport"`
	PassengerCard  bool `json:"passenger_card"`
}

type Banks struct {
	BankName string `json:"bank_name"`
}
