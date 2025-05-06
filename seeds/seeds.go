package main

import (
	"datcomtd/backend/initializers"
	"datcomtd/backend/seeds/booking"
	"datcomtd/backend/seeds/documents"
	"datcomtd/backend/seeds/products"
	"datcomtd/backend/seeds/users"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	// migrate the schemas
	users.LoadUsers()
	documents.LoadDocuments()
	booking.LoadBookings()
	products.LoadProducts()
}