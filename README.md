# Contacts API
A simple RESTful API I built to learn Go and the Echo library, where the user can do basic CRUD functionality on contacts. 

# Setup
1. Create a `.env` file 
2. Set a `MONGODB_URI` variable with your MongoDB uri
3. Run with `go run server.go` or build a binary file with `go build`

# Endpoints

Contacts must have a name, number, and description.

GET:
- `/api/v1/contacts` Returns all contacts
- `/api/v1/contacts/:id` Returns the contact with the specified id

POST:
- `/api/v1/contacts` Creates a new contact, returns the new contact.

PUT:

- `/api/v1/contacts/:id` Updates the contact with the specified id, returns the new contact.

DELETE:

- `/api/v1/contacts/:id` Deletes the contact with the specified id, returns no content.
