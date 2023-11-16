package tripgrp

import (
	"time"

	"github.com/TSMC-Uber/server/business/core/trip"
	"github.com/TSMC-Uber/server/business/sys/validate"
	"github.com/google/uuid"
)

// AppTrip represents information about an individual trip.
type AppTrip struct {
	ID             string `json:"id"`
	DriverID       string `json:"driver_id"`
	PassengerLimit int    `json:"passenger_limit"`
	SourceID       string `json:"source_id"`
	DestinationID  string `json:"destination_id"`
	Status         string `json:"status"`
	StartTime      string `json:"start_time"`
	CreatedAt      string `json:"createdAt"`
}

func toAppTrip(trip trip.Trip) AppTrip {

	return AppTrip{
		ID:             trip.ID.String(),
		DriverID:       trip.DriverID.String(),
		PassengerLimit: trip.PassengerLimit,
		SourceID:       trip.SourceID.String(),
		DestinationID:  trip.DestinationID.String(),
		Status:         trip.Status,
		StartTime:      trip.StartTime.Format(time.RFC3339),
		CreatedAt:      trip.CreatedAt.Format(time.RFC3339),
	}
}

// =============================================================================

// AppNewUser contains information needed to create a new user.
type AppNewTrip struct {
	DriverID       string `json:"driver_id" binding:"required"`
	PassengerLimit int    `json:"passenger_limit" binding:"required"`
	SourceID       string `json:"source_id" binding:"required"`
	DestinationID  string `json:"destination_id" binding:"required"`
	StartTime      string `json:"start_time" binding:"required"`
}

func toCoreNewTrip(app AppNewTrip) (trip.NewTrip, error) {
	uuDriverID, err := uuid.Parse(app.DriverID)
	if err != nil {
		return trip.NewTrip{}, err
	}
	uuSourceID, err := uuid.Parse(app.SourceID)
	if err != nil {
		return trip.NewTrip{}, err
	}
	uuDestinationID, err := uuid.Parse(app.DestinationID)
	if err != nil {
		return trip.NewTrip{}, err
	}

	// turn string to time
	startTime, err := time.Parse(time.RFC3339, app.StartTime)
	if err != nil {
		return trip.NewTrip{}, err
	}
	trip := trip.NewTrip{
		DriverID:       uuDriverID,
		PassengerLimit: app.PassengerLimit,
		SourceID:       uuSourceID,
		DestinationID:  uuDestinationID,
		StartTime:      startTime,
	}

	return trip, nil
}

// Validate checks the data in the model is considered clean.
func (app AppNewTrip) Validate() error {
	if err := validate.Check(app); err != nil {
		return err
	}
	return nil
}

type AppTripPassenger struct {
	TripID        string `json:"trip_id"`
	PassengerID   string `json:"passenger_id"`
	SourceID      string `json:"source_id"`
	DestinationID string `json:"destination_id"`
	Status        string `json:"status"`
	CreatedAt     string `json:"createdAt"`
}

func toAppTripPassenger(tripPassenger trip.TripPassenger) AppTripPassenger {

	return AppTripPassenger{
		TripID:        tripPassenger.TripID.String(),
		PassengerID:   tripPassenger.PassengerID.String(),
		SourceID:      tripPassenger.SourceID.String(),
		DestinationID: tripPassenger.DestinationID.String(),
		Status:        tripPassenger.Status,
		CreatedAt:     tripPassenger.CreatedAt.Format(time.RFC3339),
	}
}

// =============================================================================
type AppNewTripPassenger struct {
	TripID        string `json:"trip_id" binding:"required"`
	SourceID      string `json:"source_id" binding:"required"`
	DestinationID string `json:"destination_id" binding:"required"`
}

func toCoreNewTripPassenger(app AppNewTripPassenger) (trip.NewTripPassenger, error) {
	uuTripID, err := uuid.Parse(app.TripID)
	if err != nil {
		return trip.NewTripPassenger{}, err
	}
	uuSourceID, err := uuid.Parse(app.SourceID)
	if err != nil {
		return trip.NewTripPassenger{}, err
	}
	uuDestinationID, err := uuid.Parse(app.DestinationID)
	if err != nil {
		return trip.NewTripPassenger{}, err
	}

	tripPassenger := trip.NewTripPassenger{
		TripID:        uuTripID,
		SourceID:      uuSourceID,
		DestinationID: uuDestinationID,
	}

	return tripPassenger, nil
}
