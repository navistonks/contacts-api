package controllers

import (
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/navistonks/contacts-app/api/v1/configs"
	"github.com/navistonks/contacts-app/api/v1/models"
	"github.com/navistonks/contacts-app/api/v1/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

var contactCollection *mongo.Collection = configs.GetCollection(configs.DB, "contacts")
var validate = validator.New()

func GetContacts(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var contacts []models.Contact
	defer cancel()

	results, err := contactCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ContactResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error reading contacts from database.",
		})
	}

	// Get all contacts and append to list
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleContact models.Contact
		if err = results.Decode(&singleContact); err != nil {
			return c.JSON(http.StatusInternalServerError, responses.ContactResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error unmarshalling documents from database.",
			})
		}

		contacts = append(contacts, singleContact)
	}

	return c.JSON(http.StatusOK, contacts)
}

func GetContact(c echo.Context) error {

	// Get id from parameters
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("id")
	var contact models.Contact
	defer cancel()

	// Get object id
	objId, _ := primitive.ObjectIDFromHex(id)

	// Get contact from database
	err := contactCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&contact)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ContactResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error reading contact from database.",
		})
	}
	return c.JSON(http.StatusOK, contact)
}

func CreateContact(c echo.Context) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var contact models.Contact
	defer cancel()

	// Validate request body

	if err := c.Bind(&contact); err != nil {
		return c.JSON(http.StatusBadRequest, responses.ContactResponse{
			Code:    http.StatusBadRequest,
			Message: "Error validating request body.",
		})
	}

	// Validate required fields

	if validationErr := validate.Struct(&contact); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.ContactResponse{
			Code:    http.StatusBadRequest,
			Message: "Error validating request fields.",
		})
	}

	newContact := models.Contact{
		ID:          primitive.NewObjectID(),
		Name:        contact.Name,
		Number:      contact.Number,
		Description: contact.Description,
	}

	result, err := contactCollection.InsertOne(ctx, newContact)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ContactResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error inserting contact on database.",
		})
	}

	return c.JSON(http.StatusCreated, result)
}

func UpdateContact(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("id")
	var contact models.Contact
	defer cancel()

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.ContactResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid object ID.",
		})
	}

	// Validate request body

	if err := c.Bind(&contact); err != nil {
		return c.JSON(http.StatusBadRequest, responses.ContactResponse{
			Code:    http.StatusBadRequest,
			Message: "Error validating request body.",
		})
	}

	// Validate fields
	if validationErr := validate.Struct(&contact); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.ContactResponse{
			Code:    http.StatusBadRequest,
			Message: "Error validating request fields.",
		})
	}

	// Update contact on database

	update := bson.M{"name": contact.Name, "number": contact.Number, "description": contact.Description}
	result, err := contactCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ContactResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error updating contact on database.",
		})
	}

	// Get updated contact
	var updatedContact models.Contact
	if result.MatchedCount == 1 {
		err := contactCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedContact)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.ContactResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error reading updated contact from database.",
			})
		}
	}

	return c.JSON(http.StatusOK, updatedContact)
}

func DeleteContact(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("id")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	result, err := contactCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			responses.ContactResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error deleting from database.",
			})
	}

	if result.DeletedCount < 1 {
		return c.JSON(http.StatusNotFound, responses.ContactResponse{
			Code:    http.StatusNotFound,
			Message: "User not found.",
		})
	}

	return c.NoContent(http.StatusNoContent)
}
