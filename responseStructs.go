package main

type Currency struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	MinSize      string `json:"min_size"`
	MaxPrecision string `json:"max_precision"`
	Status       string `json:"status"`
	Details      struct {
		Type               string   `json:"type"`
		Symbol             string   `json:"symbol"`
		SortOrder          int      `json:"sort_order"`
		PushPaymentMethods []string `json:"push_payment_methods"`
		DisplayName        string   `json:"display_name"`
		GroupTypes         []string `json:"group_types"`
	} `json:"details"`
}
type SignedPrices struct {
	Timestamp  string   `json:"timestamp"`
	Messages   []string `json:"messages"`
	Signatures []string `json:"signatures"`
	Prices     struct {
		AdditionalProp string `json:"additionalProp"`
	} `json:"prices"`
}

type ErrorCoinbasePro struct {
	Message string `json:"message"`
}
