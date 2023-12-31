package trip

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/TSMC-Uber/server/business/data/order"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	mockStorer := new(MockStorer)
	core := NewCore(mockStorer)

	newTrip := NewTrip{
		DriverID:       uuid.New(),
		PassengerLimit: 4,
		Source: TripLocation{
			Name:    "Test Location",
			PlaceID: "Place123",
			Lat:     10.0,
			Lon:     20.0,
		},
		Destination: TripLocation{
			Name:    "Test Location",
			PlaceID: "Place123",
			Lat:     10.0,
			Lon:     20.0,
		},
		Mid: []TripLocation{
			{
				Name:    "Test Location",
				PlaceID: "Place123",
				Lat:     10.0,
				Lon:     20.0,
			},
		},
		StartTime: time.Now(),
	}

	expectedTrip := TripView{
		ID:                   uuid.New(),
		DriverID:             newTrip.DriverID,
		PassengerLimit:       newTrip.PassengerLimit,
		SourceID:             uuid.New(),
		SourceName:           newTrip.Source.Name,
		SourcePlaceID:        newTrip.Source.PlaceID,
		SourceLatitude:       newTrip.Source.Lat,
		SourceLongitude:      newTrip.Source.Lon,
		DestinationID:        uuid.New(),
		DestinationName:      newTrip.Destination.Name,
		DestinationPlaceID:   newTrip.Destination.PlaceID,
		DestinationLatitude:  newTrip.Destination.Lat,
		DestinationLongitude: newTrip.Destination.Lon,
		Mid:                  newTrip.Mid,
		Status:               TripStatusNotStarted,
		StartTime:            newTrip.StartTime,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}
	mockStorer.On("Create", mock.Anything, mock.AnythingOfType("Trip")).Return(nil)

	createdTrip, err := core.Create(context.Background(), newTrip)

	assert.NoError(t, err)
	assert.Equal(t, expectedTrip.DriverID, createdTrip.DriverID)
	assert.Equal(t, expectedTrip.PassengerLimit, createdTrip.PassengerLimit)
	assert.Equal(t, expectedTrip.SourceName, createdTrip.Source.Name)
	assert.Equal(t, expectedTrip.SourcePlaceID, createdTrip.Source.PlaceID)
	assert.Equal(t, expectedTrip.SourceLatitude, createdTrip.Source.Lat)
	assert.Equal(t, expectedTrip.SourceLongitude, createdTrip.Source.Lon)
	assert.Equal(t, expectedTrip.DestinationName, createdTrip.Destination.Name)
	assert.Equal(t, expectedTrip.DestinationPlaceID, createdTrip.Destination.PlaceID)
	assert.Equal(t, expectedTrip.DestinationLatitude, createdTrip.Destination.Lat)
	assert.Equal(t, expectedTrip.DestinationLongitude, createdTrip.Destination.Lon)
	assert.Equal(t, expectedTrip.Mid, createdTrip.Mid)
	assert.Equal(t, expectedTrip.Status, createdTrip.Status)
	assert.Equal(t, expectedTrip.StartTime, createdTrip.StartTime)
	mockStorer.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	mockStorer := new(MockStorer)
	core := NewCore(mockStorer)

	tripID := uuid.New()
	trip := TripView{
		ID:                   tripID,
		DriverID:             uuid.New(),
		PassengerLimit:       4,
		SourceID:             uuid.New(),
		SourceName:           "Test Location",
		SourcePlaceID:        "Place123",
		SourceLatitude:       10.0,
		SourceLongitude:      20.0,
		DestinationID:        uuid.New(),
		DestinationName:      "Test Location",
		DestinationPlaceID:   "Place123",
		DestinationLatitude:  10.0,
		DestinationLongitude: 20.0,
		Mid: []TripLocation{
			{
				ID:      uuid.New(),
				Name:    "Test Location",
				PlaceID: "Place123",
				Lat:     10.0,
				Lon:     20.0,
			},
		},
		Status:    TripStatusNotStarted,
		StartTime: time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	updatedPassengerLimit := 4
	updatedStatus := TripStatusNotStarted

	updateTrip := UpdateTrip{
		PassengerLimit: &updatedPassengerLimit,
		Status:         &updatedStatus,
	}

	expectedTrip := TripView{
		PassengerLimit: updatedPassengerLimit,
		Status:         updatedStatus,
	}
	mockStorer.On("Update", mock.Anything, mock.AnythingOfType("Trip")).Return(nil)
	updatedTrip, err := core.Update(context.Background(), trip, updateTrip)

	assert.NoError(t, err)
	assert.Equal(t, expectedTrip.PassengerLimit, updatedTrip.PassengerLimit)
	assert.Equal(t, expectedTrip.Status, updatedTrip.Status)
	mockStorer.AssertExpectations(t)
}

func TestQuery(t *testing.T) {
	mockStorer := new(MockStorer)
	core := NewCore(mockStorer)

	filter := QueryFilter{} // Set appropriate filter
	orderBy := order.By{}   // Set appropriate order
	pageNumber := 1
	rowsPerPage := 10

	expectedTrips := []TripView{
		{
			ID:                   uuid.New(),
			DriverID:             uuid.New(),
			PassengerLimit:       4,
			SourceID:             uuid.New(),
			SourceName:           "Test Location",
			SourcePlaceID:        "Place123",
			SourceLatitude:       10.0,
			SourceLongitude:      20.0,
			DestinationID:        uuid.New(),
			DestinationName:      "Test Location",
			DestinationPlaceID:   "Place123",
			DestinationLatitude:  10.0,
			DestinationLongitude: 20.0,
			Mid: []TripLocation{
				{
					ID:      uuid.New(),
					Name:    "Test Location",
					PlaceID: "Place123",
					Lat:     10.0,
					Lon:     20.0,
				},
			},
			Status:    TripStatusNotStarted,
			StartTime: time.Now(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:                   uuid.New(),
			DriverID:             uuid.New(),
			PassengerLimit:       4,
			SourceID:             uuid.New(),
			SourceName:           "Test Location",
			SourcePlaceID:        "Place123",
			SourceLatitude:       10.0,
			SourceLongitude:      20.0,
			DestinationID:        uuid.New(),
			DestinationName:      "Test Location",
			DestinationPlaceID:   "Place123",
			DestinationLatitude:  10.0,
			DestinationLongitude: 20.0,
			Mid: []TripLocation{
				{
					ID:      uuid.New(),
					Name:    "Test Location",
					PlaceID: "Place123",
					Lat:     10.0,
					Lon:     20.0,
				},
			},
			Status:    TripStatusNotStarted,
			StartTime: time.Now(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	mockStorer.On("Query", mock.Anything, filter, orderBy, pageNumber, rowsPerPage).Return(expectedTrips, nil)

	trips, err := core.Query(context.Background(), filter, orderBy, pageNumber, rowsPerPage)

	assert.NoError(t, err)
	for i, trip := range trips {
		assert.Equal(t, expectedTrips[i].DriverID, trip.DriverID)
		assert.Equal(t, expectedTrips[i].PassengerLimit, trip.PassengerLimit)
		assert.Equal(t, expectedTrips[i].SourceName, trip.SourceName)
		assert.Equal(t, expectedTrips[i].SourcePlaceID, trip.SourcePlaceID)
		assert.Equal(t, expectedTrips[i].SourceLatitude, trip.SourceLatitude)
		assert.Equal(t, expectedTrips[i].SourceLongitude, trip.SourceLongitude)
		assert.Equal(t, expectedTrips[i].DestinationName, trip.DestinationName)
		assert.Equal(t, expectedTrips[i].DestinationPlaceID, trip.DestinationPlaceID)
		assert.Equal(t, expectedTrips[i].DestinationLatitude, trip.DestinationLatitude)
		assert.Equal(t, expectedTrips[i].DestinationLongitude, trip.DestinationLongitude)
		assert.Equal(t, expectedTrips[i].Mid, trip.Mid)
		assert.Equal(t, expectedTrips[i].Status, trip.Status)
		assert.Equal(t, expectedTrips[i].StartTime, trip.StartTime)
	}
	mockStorer.AssertExpectations(t)
}

func TestQueryByID(t *testing.T) {
	mockStorer := new(MockStorer)
	core := NewCore(mockStorer)

	tripID := uuid.New()

	expectedTrip := TripView{
		ID:                   tripID,
		DriverID:             uuid.New(),
		PassengerLimit:       4,
		SourceID:             uuid.New(),
		SourceName:           "Test Location",
		SourcePlaceID:        "Place123",
		SourceLatitude:       10.0,
		SourceLongitude:      20.0,
		DestinationID:        uuid.New(),
		DestinationName:      "Test Location",
		DestinationPlaceID:   "Place123",
		DestinationLatitude:  10.0,
		DestinationLongitude: 20.0,
		Mid: []TripLocation{
			{
				ID:      uuid.New(),
				Name:    "Test Location",
				PlaceID: "Place123",
				Lat:     10.0,
				Lon:     20.0,
			},
		},
		Status:    TripStatusNotStarted,
		StartTime: time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockStorer.On("QueryByID", mock.Anything, tripID).Return(expectedTrip, nil)

	trip, err := core.QueryByID(context.Background(), tripID)

	assert.NoError(t, err)
	assert.Equal(t, expectedTrip.DriverID, trip.DriverID)
	assert.Equal(t, expectedTrip.PassengerLimit, trip.PassengerLimit)
	assert.Equal(t, expectedTrip.SourceName, trip.SourceName)
	assert.Equal(t, expectedTrip.SourcePlaceID, trip.SourcePlaceID)
	assert.Equal(t, expectedTrip.SourceLatitude, trip.SourceLatitude)
	assert.Equal(t, expectedTrip.SourceLongitude, trip.SourceLongitude)
	assert.Equal(t, expectedTrip.DestinationName, trip.DestinationName)
	assert.Equal(t, expectedTrip.DestinationPlaceID, trip.DestinationPlaceID)
	assert.Equal(t, expectedTrip.DestinationLatitude, trip.DestinationLatitude)
	assert.Equal(t, expectedTrip.DestinationLongitude, trip.DestinationLongitude)
	assert.Equal(t, expectedTrip.Mid, trip.Mid)
	assert.Equal(t, expectedTrip.Status, trip.Status)
	assert.Equal(t, expectedTrip.StartTime, trip.StartTime)
	mockStorer.AssertExpectations(t)
}

func TestQueryByIDError(t *testing.T) {
	mockStorer := new(MockStorer)
	core := NewCore(mockStorer)

	tripID := uuid.New()

	mockStorer.On("QueryByID", mock.Anything, tripID).Return(TripView{}, errors.New("query by id error"))

	_, err := core.QueryByID(context.Background(), tripID)
	assert.Error(t, err)
	fmt.Println(err.Error())
	assert.Contains(t, err.Error(), "query by id error")
	mockStorer.AssertExpectations(t)
}

func TestCount(t *testing.T) {
	mockStorer := new(MockStorer)
	core := NewCore(mockStorer)

	filter := QueryFilter{} // Set appropriate filter

	expectedCount := 10
	mockStorer.On("Count", mock.Anything, filter).Return(expectedCount, nil)

	count, err := core.Count(context.Background(), filter)

	assert.NoError(t, err)
	assert.Equal(t, expectedCount, count)
	mockStorer.AssertExpectations(t)
}

func TestCountError(t *testing.T) {
	mockStorer := new(MockStorer)
	core := NewCore(mockStorer)

	filter := QueryFilter{} // Set appropriate filter

	mockStorer.On("Count", mock.Anything, filter).Return(0, errors.New("count error"))

	_, err := core.Count(context.Background(), filter)
	assert.Error(t, err)
	fmt.Println(err.Error())
	assert.Contains(t, err.Error(), "count error")
	mockStorer.AssertExpectations(t)
}

func TestQueryMyTrip(t *testing.T) {
	mockStorer := new(MockStorer)
	core := NewCore(mockStorer)

	userID := uuid.New()
	filter := QueryFilterByUser{} // Set appropriate filter
	orderBy := order.By{}         // Set appropriate order
	pageNumber := 1
	rowsPerPage := 10

	expectedTrips := []UserTrip{
		{
			TripID:          uuid.New(),
			PassengerID:     uuid.New(),
			MySourceID:      uuid.New(),
			MyDestinationID: uuid.New(),
			MyStatus:        TripStatusNotStarted,
			DriverID:        uuid.New(),
			PassengerLimit:  4,
			SourceID:        uuid.New(),
			DestinationID:   uuid.New(),
			TripStatus:      TripStatusNotStarted,
			StartTime:       time.Now(),
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
		{
			TripID:          uuid.New(),
			PassengerID:     uuid.New(),
			MySourceID:      uuid.New(),
			MyDestinationID: uuid.New(),
			MyStatus:        TripStatusNotStarted,
			DriverID:        uuid.New(),
			PassengerLimit:  4,
			SourceID:        uuid.New(),
			DestinationID:   uuid.New(),
			TripStatus:      TripStatusNotStarted,
			StartTime:       time.Now(),
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
	}

	mockStorer.On("QueryMyTrip", mock.Anything, userID, filter, orderBy, pageNumber, rowsPerPage).Return(expectedTrips, nil)

	trips, err := core.QueryMyTrip(context.Background(), userID, filter, orderBy, pageNumber, rowsPerPage)

	assert.NoError(t, err)
	for i, trip := range trips {
		assert.Equal(t, expectedTrips[i].TripID, trip.TripID)
		assert.Equal(t, expectedTrips[i].PassengerID, trip.PassengerID)
		assert.Equal(t, expectedTrips[i].MySourceID, trip.MySourceID)
		assert.Equal(t, expectedTrips[i].MyDestinationID, trip.MyDestinationID)
		assert.Equal(t, expectedTrips[i].MyStatus, trip.MyStatus)
		assert.Equal(t, expectedTrips[i].DriverID, trip.DriverID)
		assert.Equal(t, expectedTrips[i].PassengerLimit, trip.PassengerLimit)
		assert.Equal(t, expectedTrips[i].SourceID, trip.SourceID)
		assert.Equal(t, expectedTrips[i].DestinationID, trip.DestinationID)
		assert.Equal(t, expectedTrips[i].TripStatus, trip.TripStatus)
		assert.Equal(t, expectedTrips[i].StartTime, trip.StartTime)
	}
	mockStorer.AssertExpectations(t)
}

func TestQueryMyTripError(t *testing.T) {
	mockStorer := new(MockStorer)
	core := NewCore(mockStorer)

	userID := uuid.New()
	filter := QueryFilterByUser{} // Set appropriate filter
	orderBy := order.By{}         // Set appropriate order
	pageNumber := 1
	rowsPerPage := 10

	mockStorer.On("QueryMyTrip", mock.Anything, userID, filter, orderBy, pageNumber, rowsPerPage).Return([]UserTrip{}, errors.New("query my trip error"))

	_, err := core.QueryMyTrip(context.Background(), userID, filter, orderBy, pageNumber, rowsPerPage)
	assert.Error(t, err)
	fmt.Println(err.Error())
	assert.Contains(t, err.Error(), "query my trip error")
	mockStorer.AssertExpectations(t)
}

func TestCreateRating(t *testing.T) {
	mockStorer := new(MockStorer)
	core := NewCore(mockStorer)

	newRating := NewRating{
		TripID:  uuid.New(),
		Rating:  5.0,
		Comment: "some comment",
	}

	mockStorer.On("CreateRating", mock.Anything, mock.AnythingOfType("Rating")).Return(nil)

	rating, err := core.CreateRating(context.Background(), newRating.TripID, newRating)

	assert.NoError(t, err)
	assert.Equal(t, newRating.TripID, rating.TripID)
	assert.Equal(t, newRating.Rating, rating.Rating)
	assert.Equal(t, newRating.Comment, rating.Comment)
	mockStorer.AssertExpectations(t)
}

func TestCreateRatingError(t *testing.T) {
	mockStorer := new(MockStorer)
	core := NewCore(mockStorer)

	newRating := NewRating{
		TripID:  uuid.New(),
		Rating:  5.0,
		Comment: "some comment",
	}

	mockStorer.On("CreateRating", mock.Anything, mock.AnythingOfType("Rating")).Return(errors.New("create rating error"))

	_, err := core.CreateRating(context.Background(), newRating.TripID, newRating)
	assert.Error(t, err)
	fmt.Println(err.Error())
	assert.Contains(t, err.Error(), "create rating error")
	mockStorer.AssertExpectations(t)
}

func TestCreateTripPassenger(t *testing.T) {
	mockStorer := new(MockStorer)
	core := NewCore(mockStorer)

	newTripPassenger := NewTripPassenger{
		PassengerID:   uuid.New(),
		SourceID:      uuid.New(),
		DestinationID: uuid.New(),
	}

	expectedTripPassenger := TripPassenger{
		TripID:        uuid.New(),
		PassengerID:   newTripPassenger.PassengerID,
		SourceID:      newTripPassenger.SourceID,
		DestinationID: newTripPassenger.DestinationID,
		Status:        StatusPending,
		CreatedAt:     time.Now(),
	}

	mockStorer.On("Join", mock.Anything, mock.AnythingOfType("TripPassenger")).Return(nil)

	tripPassenger, err := core.Join(context.Background(), expectedTripPassenger.TripID, newTripPassenger)

	assert.NoError(t, err)
	assert.Equal(t, expectedTripPassenger.TripID, tripPassenger.TripID)
	assert.Equal(t, expectedTripPassenger.PassengerID, tripPassenger.PassengerID)
	assert.Equal(t, expectedTripPassenger.SourceID, tripPassenger.SourceID)
	assert.Equal(t, expectedTripPassenger.DestinationID, tripPassenger.DestinationID)
	assert.Equal(t, expectedTripPassenger.Status, tripPassenger.Status)
	mockStorer.AssertExpectations(t)
}

func TestCreateTripPassengerError(t *testing.T) {
	mockStorer := new(MockStorer)
	core := NewCore(mockStorer)

	newTripPassenger := NewTripPassenger{
		PassengerID:   uuid.New(),
		SourceID:      uuid.New(),
		DestinationID: uuid.New(),
	}

	mockStorer.On("Join", mock.Anything, mock.AnythingOfType("TripPassenger")).Return(errors.New("create trip passenger error"))

	_, err := core.Join(context.Background(), uuid.New(), newTripPassenger)
	assert.Error(t, err)
	fmt.Println(err.Error())
	assert.Contains(t, err.Error(), "create trip passenger error")
	mockStorer.AssertExpectations(t)
}

func TestQueryPassengers(t *testing.T) {
	mockStorer := new(MockStorer)
	core := NewCore(mockStorer)

	tripID := uuid.New()

	expectedTripDetails := TripDetails{
		TripID:               tripID,
		DriverID:             uuid.New(),
		DriverName:           "Test Driver",
		DriverImageURL:       "https://www.google.com",
		DriverBrand:          "Toyota",
		DriverModel:          "Camry",
		DriverColor:          "White",
		DriverPlate:          "ABC1234",
		SourceName:           "Test Location",
		SourcePlaceID:        "Place123",
		SourceLatitude:       10.0,
		SourceLongitude:      20.0,
		DestinationName:      "Test Location",
		DestinationPlaceID:   "Place123",
		DestinationLatitude:  10.0,
		DestinationLongitude: 20.0,
	}

	mockStorer.On("QueryPassengers", mock.Anything, tripID).Return(expectedTripDetails, nil)

	tripDetails, err := core.QueryPassengers(context.Background(), tripID)

	assert.NoError(t, err)
	assert.Equal(t, expectedTripDetails.TripID, tripDetails.TripID)
	assert.Equal(t, expectedTripDetails.DriverID, tripDetails.DriverID)
	assert.Equal(t, expectedTripDetails.DriverName, tripDetails.DriverName)
	assert.Equal(t, expectedTripDetails.DriverImageURL, tripDetails.DriverImageURL)
	assert.Equal(t, expectedTripDetails.DriverBrand, tripDetails.DriverBrand)
	assert.Equal(t, expectedTripDetails.DriverModel, tripDetails.DriverModel)
	assert.Equal(t, expectedTripDetails.DriverColor, tripDetails.DriverColor)
	assert.Equal(t, expectedTripDetails.DriverPlate, tripDetails.DriverPlate)
	assert.Equal(t, expectedTripDetails.SourceName, tripDetails.SourceName)
	assert.Equal(t, expectedTripDetails.SourcePlaceID, tripDetails.SourcePlaceID)
	assert.Equal(t, expectedTripDetails.SourceLatitude, tripDetails.SourceLatitude)
	assert.Equal(t, expectedTripDetails.SourceLongitude, tripDetails.SourceLongitude)
	assert.Equal(t, expectedTripDetails.DestinationName, tripDetails.DestinationName)
	assert.Equal(t, expectedTripDetails.DestinationPlaceID, tripDetails.DestinationPlaceID)
	assert.Equal(t, expectedTripDetails.DestinationLatitude, tripDetails.DestinationLatitude)
	assert.Equal(t, expectedTripDetails.DestinationLongitude, tripDetails.DestinationLongitude)
	mockStorer.AssertExpectations(t)
}

func TestQueryPassengersError(t *testing.T) {
	mockStorer := new(MockStorer)
	core := NewCore(mockStorer)

	tripID := uuid.New()

	mockStorer.On("QueryPassengers", mock.Anything, tripID).Return(TripDetails{}, errors.New("query passengers error"))

	_, err := core.QueryPassengers(context.Background(), tripID)
	assert.Error(t, err)
	fmt.Println(err.Error())
	assert.Contains(t, err.Error(), "query passengers error")
	mockStorer.AssertExpectations(t)
}
