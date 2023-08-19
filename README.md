# Transfers API

## How to run
```
docker-compose up -d
```

## End points

### Create account
`POST /accounts/`

####  Request
    curl --request POST \
        --url http://localhost:3000/accounts \
        --header 'Content-Type: application/json' \
        --data '{
            "name": "ronnie coleman",
            "cpf": "65223068084",
            "secret": "123",
            "balance": 1000.00
        }'

####  Response
    {
	    "id": "c6b0b0bc-927c-4d2a-8f9f-c15186cb48d4",
	    "name": "ronnie coleman",
	    "cpf": "65223068084",
	    "balance": 1000,
	    "created_at": "2023-08-18 01:22:10"
    }

### Get all accounts
`GET /accounts/`

####  Request
    curl --request GET \
        --url http://localhost:3000/accounts

####  Response
    [
	    {
		    "id": "c6b0b0bc-927c-4d2a-8f9f-c15186cb48d4",
		    "name": "ronnie coleman",
		    "cpf": "65223068084",
		    "balance": 1000,
		    "created_at": "2023-08-18 01:22:10"
	    },
        {
		    "id": "1903775c-f754-4230-8415-493b17623cd6",
		    "name": "chris bumstead",
		    "cpf": "84576735055",
		    "balance": 1000,
		    "created_at": "2023-08-19 00:19:35"
	    }
    ]
### Get balance
`GET /accounts/{id}/balance`

####  Request

    curl --request GET \
        --url http://localhost:3000/accounts/c6b0b0bc-927c-4d2a-8f9f-c15186cb48d4/balance

####  Response
    {
	    "balance": 1000
    }

### Login
`POST /login`

####  Request
    curl --request POST \
        --url http://localhost:3000/login \
        --header 'Content-Type: application/json' \
        --data '{
            "cpf": "14283939005",
            "secret": "1233"
        }'

####  Response
    {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcGYiOiI2NTIyMzA2ODA4NCIsInNlY3JldCI6IjEyMyIsImV4cCI6MTY5MjQwNDQ4Nn0.XWRYTY1ALlxDojE4Xl1HEGmLrvdxttXESyQYZvjjmK4"
    }

### Create transfer

`POST /tranfers`

####  Request
    curl --request POST \
        --url http://localhost:3000/tranfers \
        --header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcGYiOiI2NTIyMzA2ODA4NCIsInNlY3JldCI6IjEyMyIsImV4cCI6MTY5MjQwNDA5N30.SWF6aLn_w23rS2nSVHFVP-UINheeray_59plarQLW_o' \
        --header 'Content-Type: application/json' \
        --data '{
            "account_destination_id": "1903775c-f754-4230-8415-493b17623cd6",
            "amount": 150.50
        }'

####  Response

    {
        "id": "92a632b0-4d6a-4d85-ae16-03c091840dfb",
        "account_origin_id": "c6b0b0bc-927c-4d2a-8f9f-c15186cb48d4",
        "account_destination_id": "1903775c-f754-4230-8415-493b17623cd6",
        "amount": 150.5,
        "created_at": "2023-08-19 00:22:03"
    }

### Get transfer

`GET /tranfers`

####  Request
    curl --request GET \
        --url http://localhost:3000/tranfers \
        --header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcGYiOiI2NTIyMzA2ODA4NCIsInNlY3JldCI6IjEyMyIsImV4cCI6MTY5MjQwNDU4M30.bNfGiDN2c3MteGAwjkC2TccJkpAt4mD0d_8nLw1D0tQ'

####  Response

    [
        {
            "id": "92a632b0-4d6a-4d85-ae16-03c091840dfb",
            "account_origin_id": "c6b0b0bc-927c-4d2a-8f9f-c15186cb48d4",
            "account_destination_id": "1903775c-f754-4230-8415-493b17623cd6",
            "amount": 150.5,
            "created_at": "2023-08-19 00:22:03"
        }
    ]