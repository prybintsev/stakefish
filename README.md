# Ethereum API


## Running the application

To run the application, execute the following command in the terminal from the project root directory
```
docker-compose up --build
```

## Swagger

Swagger definitions are located under docs directory.
They can be generated automatically by running the following command:
```shell
make swagger
```

## API usage examples

Get application info:
```shell

```

Lookup IPv4 addresses of a specific domain
```shell
curl -X 'GET' 'localhost:3000/v1/tools/lookup?domain=google.com' -H 'accept: application/json' | jq
```

Response:
```shell
{
  "client_ip": "172.19.0.1",
  "created_at": 1666323692,
  "domain": "google.com",
  "addresses": [
    {
      "ip": "142.250.217.110"
    }
  ]
}
```



