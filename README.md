# Simple Golang Api

For now implemented for cars only.
For routing used basic http library tried gorilla/mux but did not like it

## Routes

| Method    | Route                         | Content-Type              |
|-----------|-------------------------------|------------------|
| GET       | /cars                         | no           |
| POST      | /cars                         | body(form-data)   |
| DELETE    | /cars?id={int}             | query params(raw query)  |
| PUT       | /cars?id={int}             | query params(raw query)  |
| GET       | /cars?id={int}             | query params(raw query)  |
| GET       | /statuses                      | no       |
| POST      | /statuses                      | body(form-data)       |

### Minuses
- Missed some edge cases (e.g. check for empty param, form-data etc)
- No validation whatsoever
- In-memory storage limited to 10 records