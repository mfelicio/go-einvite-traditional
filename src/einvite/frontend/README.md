frontend
==
This package contains the components for the frontend layer.

App users can interact with the web application through RESTful http services and websocket connections available in this package.

Api is described below:

## Users:

All operations require authentication

### /user/:id/profile
- GET  
	- Gets the user :id profile
- PUT 
	- Updates the current user profile
	- HTTP 403 if current user can't see user :id

### /user/:id/contacts
- GET
	- Gets the user :id contacts
	- HTTP 403 if current user can't see user :id

### /user/:id/events
- GET
	- Gets all events for user :id
	- HTTP 403 if current user can't see user :id
	
### /user/:id/events/search?q=:searchTerms
- GET
	- Searches all events for user :id with :searchTerms
	- HTTP 403 if current user can't see user :id

### /user/:id/message
- POST
	- Send message to user :id
	- HTTP 403 if current user can't see user :id

## Events

### /events
- POST
	- Creates a new event

### /event/:id
- GET
	- Gets the details of event with :id
	- HTTP 403 if the curent user can't see event :id
- PUT
	- Updates the event with :id
	- HTTP 403 if the current user can't see event :id

### 


## Users (phase 2)

### /user/:id/invite/:eventId
- PUT
	- Invites user to event :eventId
	- HTTP 403 if user can't invite for event :eventId