# Basic app

This is a basic web application. I did this to see how I could structure a go backend.
Supports register, login and whoami.

This is not production ready! It is just for learning.


## Run

```bash
go run main.go serve | fblog -d
```


## Manual testing
```bash
http POST localhost:8080/register email=test@test.de password=fkrb

# Extract cookie from this
http POST localhost:8080/login email=test@test.de password=fkrb


# use cookie here
http GET localhost:8080/whoami Cookie:<COOKIE_VALUE>

```
