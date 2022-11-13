package models

type Response struct {
	ContactEmail  string `json:"contactEmail"`
	PhoneNumber   string `json:"phoneNumber"`
	NumPeople     string `json:"numPeople"`
	Date          string `json:"date"`
	Name          string `json:"name"`
	BookingNumber string `json:"bookingNumber"`
	Tour          string `json:"tour"`
	RawEmail      string `json:"rawemail"`
}
