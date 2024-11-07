package schemas

type ChargeCustomerInput struct {
	Amount   float64
	Customer Customer `validate:"required" json:"customer"`

	MetaData map[string]any `validate:"json" json:"meta_data"`
}

type Customer struct {
	Phone   string `validate:"required|ghPhone" json:"phone"`
	Network string `validate:"required|in:mtn,vodafone,airteltigo" json:"network"`
}
