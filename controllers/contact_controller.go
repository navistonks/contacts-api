package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// TODO: Fix DELETE route response

type Contact struct {
	ID          int    `json:"id"`
	Name        string `json:"Name"`
	Number      uint   `json:"Number"`
	Description string `json:"Description"`
}

// Map id to contacts
var (
	Contacts     = map[int]*Contact{}
	sequence int = 1
)

func GetContacts(c echo.Context) error {
	return c.JSON(http.StatusOK, Contacts)
}

func GetContact(c echo.Context) error {

	// Convert id parameter into an integer
	id, _ := strconv.Atoi(c.Param("id"))

	return c.JSON(http.StatusOK, Contacts[id])
}

func CreateContact(c echo.Context) error {

	contact := &Contact{
		ID: sequence,
	}

	// Bind request into Contact struct
	if err := c.Bind(contact); err != nil {
		return err
	}

	// Add contact to array
	Contacts[contact.ID] = contact

	sequence++
	return c.JSON(http.StatusCreated, contact)

}

func UpdateContact(c echo.Context) error {
	contact := new(Contact)

	// Bind request to Contact struct
	if err := c.Bind(contact); err != nil {
		return err
	}

	// Convert id parameter into an integer
	id, _ := strconv.Atoi(c.Param("id"))

	Contacts[id].Name = contact.Name
	Contacts[id].Description = contact.Description
	Contacts[id].Number = contact.Number

	return c.JSON(http.StatusOK, Contacts[id])
}

func DeleteContact(c echo.Context) error {
	// Convert id string to int
	id, _ := strconv.Atoi(c.Param("id"))

	contact := Contacts[id]
	delete(Contacts, id)

	// Return deleted contact
	return c.JSON(http.StatusOK, contact)
}
