package user

import (
	"GarageSaleAPI/domain/address"
	"time"
)

type Seller struct {
	id             string
	userId         string
	savedAddresses []SavedAddress
	createdAt      time.Time
}

type SavedAddress struct {
	id        string
	label     string
	address   address.Address
	isDefault bool
}
