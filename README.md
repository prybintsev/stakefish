# Stakefish


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

### Get application info
```shell
curl -X 'GET' 'localhost:3000' -H 'accept: application/json' | jq
```

Response:
```shell
{
  "version": "v0.0.2",
  "date": 1666324435,
  "kubernetes": false
}
```

### Lookup IPv4 addresses of a specific domain
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
### Validate an IPv4 address

```shell
curl -X 'POST' 'http://localhost:3000/v1/tools/validate' -H 'accept: application/json' -H 'Content-Type: application/json' -d '{"ip": "1.1.1.1"}' | jq
```

Response:
```shell
{
  "status": true
}
```

## Retrieve lookups history

```shell
curl -X 'GET' 'http://localhost:3000/v1/history'  -H 'accept: application/json' | jq
```

Response:
```shell
[
  {
    "client_ip": "172.19.0.1",
    "created_at": 1666323692,
    "domain": "google.com",
    "addresses": [
      {
        "ip": "142.250.217.110"
      }
    ]
  },
  {
    "client_ip": "172.19.0.1",
    "created_at": 1666323687,
    "domain": "google.com",
    "addresses": [
      {
        "ip": "142.250.217.110"
      }
    ]
  },
  {
    "client_ip": "172.19.0.1",
    "created_at": 1666323107,
    "domain": "reddit.com",
    "addresses": [
      {
        "ip": "151.101.65.140"
      },
      {
        "ip": "151.101.1.140"
      },
      {
        "ip": "151.101.129.140"
      },
      {
        "ip": "151.101.193.140"
      }
    ]
  }
]
```