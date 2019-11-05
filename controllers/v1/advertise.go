package controllers

import (
	"net/http"
	"time"

	"github.com/amiraliio/advertiselocator/helpers"
	"github.com/amiraliio/advertiselocator/models"
	"github.com/amiraliio/advertiselocator/repositories/v1"
	"github.com/amiraliio/advertiselocator/requests"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//TODO check access just person can add advertise
//TODO use controller service

func advertiseRepository() repositories.AdvertiseInterface {
	return new(repositories.AdvertiseRepository)
}

func AddAdvertise(request echo.Context) (err error) {
	authData := request.Get(models.AuthorizationHeaderKey)
	if !helpers.IsInstance(authData, (*models.Client)(nil)) {
		return helpers.ResponseError(
			request,
			http.StatusBadRequest,
			helpers.InsertTarget,
			http.StatusText(http.StatusBadRequest),
			"CA1001",
			"Insert Advertise",
			err.Error(),
		)
	}
	advertiseRequest := new(requests.Advertise)
	if err = request.Bind(advertiseRequest); err != nil {
		return helpers.ResponseError(
			request,
			http.StatusBadRequest,
			helpers.InsertTarget,
			http.StatusText(http.StatusBadRequest),
			"CA1002",
			"Insert Advertise",
			err.Error(),
		)
	}
	if err = request.Validate(advertiseRequest); err != nil {
		//TODO problem validation error response
		return helpers.ResponseError(
			request,
			http.StatusUnprocessableEntity,
			helpers.InsertTarget,
			http.StatusText(http.StatusUnprocessableEntity),
			"CA1003",
			"Insert Advertise",
			err.Error(),
		)
	}
	advertise := new(models.Advertise)
	advertise.Status = models.ActiveStatus
	advertiseID := primitive.NewObjectID()
	advertise.ID = advertiseID
	advertise.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	advertise.CreatedBy = authData.(*models.Client).UserID
	advertise.Location = advertiseRequest.Location
	advertise.Tags = advertiseRequest.Tags
	person := new(models.Person)
	person.ID = authData.(*models.Client).UserID
	advertise.Advertiser = person
	advertise.Radius = advertiseRequest.Radius
	advertise.Images = advertiseRequest.Images
	advertise.Title = advertiseRequest.Title
	advertise.Description = advertiseRequest.Description
	advertise.Visibility = advertiseRequest.Visibility
	result, err := advertiseRepository().InsertAdvertise(advertise)
	if err != nil {
		return helpers.ResponseError(
			request,
			http.StatusBadRequest,
			helpers.InsertTarget,
			http.StatusText(http.StatusBadRequest),
			"CA1004",
			"Insert Advertise",
			err.Error(),
		)
	}
	return helpers.ResponseOk(request, http.StatusCreated, result)
}

func ListOfAdvertises(request echo.Context) (err error) {
	auth := request.Get(models.AuthorizationHeaderKey)
	if !helpers.IsInstance(auth, (*models.Client)(nil)) {
		return helpers.ResponseError(
			request,
			http.StatusBadRequest,
			helpers.QueryTarget,
			http.StatusText(http.StatusBadRequest),
			"CA1006",
			"List Advertise",
			err.Error(),
		)
	}
	// queries := request.QueryParams()
	filter := new(models.AdvertiseFilter)
	filter.UserID = auth.(*models.Client).UserID
	results, err := advertiseRepository().ListOfAdvertise(filter)
	if err != nil {
		return helpers.ResponseError(
			request,
			http.StatusBadRequest,
			helpers.QueryTarget,
			http.StatusText(http.StatusBadRequest),
			"CA1005",
			"List Of Advertise",
			err.Error(),
		)
	}
	return helpers.ResponseOk(request, http.StatusOK, results)
}

func GetAdvertise(request echo.Context) (err error) {
	return nil
}

func DeleteAdvertise(request echo.Context) (err error) {
	return nil
}

func UpdateAdvertise(request echo.Context) (err error) {
	return nil
}
