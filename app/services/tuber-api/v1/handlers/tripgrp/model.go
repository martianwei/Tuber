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
		SourceID:       trip.Source.ID.String(),
		DestinationID:  trip.Destination.ID.String(),
		Status:         trip.Status,
		StartTime:      trip.StartTime.Format(time.RFC3339),
		CreatedAt:      trip.CreatedAt.Format(time.RFC3339),
	}
}

type AppNewTripLocation struct {
	Name    string  `json:"name" binding:"required"`
	PlaceID string  `json:"place_id" binding:"required"`
	Lat     float64 `json:"lat" binding:"required"`
	Lon     float64 `json:"lon" binding:"required"`
}

type AppNewTrip struct {
	DriverID       string               `json:"driver_id"`
	PassengerLimit int                  `json:"passenger_limit"`
	Source         AppNewTripLocation   `json:"source"`
	Destination    AppNewTripLocation   `json:"destination"`
	Mid            []AppNewTripLocation `json:"mid"`
	StartTime      string               `json:"start_time" binding:"required"`
}

func toCoreNewTrip(app AppNewTrip) (trip.NewTrip, error) {
	uuDriverID, err := uuid.Parse(app.DriverID)
	if err != nil {
		return trip.NewTrip{}, err
	}

	// turn string to time
	startTime, err := time.Parse(time.RFC3339, app.StartTime)
	if err != nil {
		return trip.NewTrip{}, err
	}

	// construct mid
	mid := []trip.TripLocation{}
	for _, appMid := range app.Mid {
		mid = append(mid, trip.TripLocation{
			Name:    appMid.Name,
			PlaceID: appMid.PlaceID,
			Lat:     appMid.Lat,
			Lon:     appMid.Lon,
		})
	}

	trip := trip.NewTrip{
		DriverID:       uuDriverID,
		PassengerLimit: app.PassengerLimit,
		Source: trip.TripLocation{
			Name:    app.Source.Name,
			PlaceID: app.Source.PlaceID,
			Lat:     app.Source.Lat,
			Lon:     app.Source.Lon,
		},
		Destination: trip.TripLocation{
			Name:    app.Destination.Name,
			PlaceID: app.Destination.PlaceID,
			Lat:     app.Destination.Lat,
			Lon:     app.Destination.Lon,
		},
		Mid:       mid,
		StartTime: startTime,
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

// ID:                   dbTripView.ID,
//
//	DriverName:           dbTripView.DriverName,
//	DriverBrand:          dbTripView.DriverBrand,
//	DriverModel:          dbTripView.DriverModel,
//	DriverColor:          dbTripView.DriverColor,
//	DriverPlate:          dbTripView.DriverPlate,
//	SourceName:           dbTripView.SourceName,
//	SourcePlaceID:        dbTripView.SourcePlaceID,
//	SourceLatitude:       dbTripView.SourceLatitude,
//	SourceLongitude:      dbTripView.SourceLongitude,
//	DestinationName:      dbTripView.DestinationName,
//	DestinationPlaceID:   dbTripView.DestinationPlaceID,
//	DestinationLatitude:  dbTripView.DestinationLatitude,
//	DestinationLongitude: dbTripView.DestinationLongitude,
//	Status:               dbTripView.Status,
//	StartTime:            dbTripView.StartTime.In(time.Local),
//	CreatedAt:            dbTripView.CreatedAt.In(time.Local),
type AppTripView struct {
	ID                   string  `json:"id"`
	DriverName           string  `json:"driver_name"`
	DriverBrand          string  `json:"driver_brand"`
	DriverModel          string  `json:"driver_model"`
	DriverColor          string  `json:"driver_color"`
	DriverPlate          string  `json:"driver_plate"`
	SourceName           string  `json:"source_name"`
	SourcePlaceID        string  `json:"source_place_id"`
	SourceLatitude       float64 `json:"source_latitude"`
	SourceLongitude      float64 `json:"source_longitude"`
	DestinationName      string  `json:"destination_name"`
	DestinationPlaceID   string  `json:"destination_place_id"`
	DestinationLatitude  float64 `json:"destination_latitude"`
	DestinationLongitude float64 `json:"destination_longitude"`
	Status               string  `json:"status"`
	StartTime            string  `json:"start_time"`
	CreatedAt            string  `json:"createdAt"`
}

func toAppTripView(tripView trip.TripView) AppTripView {

	return AppTripView{
		ID:                   tripView.ID.String(),
		DriverName:           tripView.DriverName,
		DriverBrand:          tripView.DriverBrand,
		DriverModel:          tripView.DriverModel,
		DriverColor:          tripView.DriverColor,
		DriverPlate:          tripView.DriverPlate,
		SourceName:           tripView.SourceName,
		SourcePlaceID:        tripView.SourcePlaceID,
		SourceLatitude:       tripView.SourceLatitude,
		SourceLongitude:      tripView.SourceLongitude,
		DestinationName:      tripView.DestinationName,
		DestinationPlaceID:   tripView.DestinationPlaceID,
		DestinationLatitude:  tripView.DestinationLatitude,
		DestinationLongitude: tripView.DestinationLongitude,
		Status:               tripView.Status,
		StartTime:            tripView.StartTime.Format(time.RFC3339),
		CreatedAt:            tripView.CreatedAt.Format(time.RFC3339),
	}
}
