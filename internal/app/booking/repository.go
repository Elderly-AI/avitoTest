package booking

import bookingModel "github.com/Kostikans/avitoTest/internal/app/booking/model"

type Repository interface {
	AddBooking(booking bookingModel.BookingAdd) (bookingModel.BookingID, error)
	DeleteRoom(bookingID int64) error
}
