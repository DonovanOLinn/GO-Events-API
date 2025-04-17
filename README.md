# Events

A Go-powered "Event Booking" REST API

### TODO
```bash
- Unit Testing
- Dockerize
- Deploy on AWS EC2 instance
- Connect AWS RDS Database
```

### Routes
```bash
GET /events                     => Get a list of available events
GET /events/<id>                => Get a list of available events
POST /events                    => Create a new bookable event
PUT /events/<id>                => Update an event
DELETE /events/<id>             => Delete an event
Post /signup                    => Create new user
Post /login                     => AUthenticate user
Post /events/<id>/register      => Register User for Event
Delete /events/<id>/register    => Cancel Registration
```