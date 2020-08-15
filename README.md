# Simple Golang Api

For now implemented for cars only.
For routing used basic http library tried gorilla/mux but did not like it

## Routes

| Method    | Route                         | Content-Type              |
|-----------|-------------------------------|------------------|
| GET       | /cars                          | no                   |
| POST      | /cars                          | body(form-data)      |
| DELETE    | /cars?id={int}                 | query params(raw query)  |
| PUT       | /cars?id={int}                 | query params(raw query)  |
| GET       | /cars?id={int}                 | query params(raw query)  |
| GET       | /cars?id={int}&status=                 | query params(raw query)  |
| GET       | /statuses                      | no                       |
| POST      | /statuses                      | body(form-data)               |
| DELETE    | /statuses?id={int}             | query params(raw query)       |
| PUT       | /statuses?id={int}             | query params(raw query)       |
| GET       | /statuses?id={int}             | query params(raw query)       |

## Usage
1. To start type ``go run main.go`` or build it if you want.
2. All dummy data created on storage creation
3. On delete all data is displayed (for convenience)
4. ``/cars?id={int}&status=`` to take data with connected status:
```
{
     "id": 2,
     "brand": "BMW",
     "model": "x3",
     "price": 23444,
     "status": {
         "id": 2,
         "name": "На складе"
     },
     "mileage": 3222
}
```

### Minuses
- Missed some edge cases (e.g. check for empty param, form-data etc)
- No validation whatsoever
- In-memory storage limited to 10 records