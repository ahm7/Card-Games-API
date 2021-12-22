# Card Games REST API in GO

> Simple RESTful API for Card games. with database implementation 

## Quick Start


``` bash
# Install mux router
cd /main
go get -u github.com/gorilla/mux
```
``` bash
# update mongo url
in DataBase/connector.go 
update var mongoUrl with your database user name & password

```

``` bash
# start the api server
cd /main
go run main.go
```
``` bash
# to run tests
cd /main/Test
go test -v
```

## Endpoints

### open deck
``` bash
GET localhost:8001/api/Deck/{id}
```

### Create Deck
``` bash
POST localhost:8001/api/Deck

# Request sample
# {
#   "shuffled" : true,
#   "cards":"AH,AS,AD,AC"
# }
```

### Draw a card
``` bash
POST localhost:8001/api/Deck/{id}/draw

# Request sample
# {
#   "count":2
# }
```

# App Structure explain

### 1. ApiHelpers
#### Basically contains the helper functions used in returning api responses, HTTP status codes, default messages etc.

### 2. Controller
#### Contains handler functions for particular route to be called when an api is called.

### 3. Helpers
#### Contains helper functions used in all apis
### 4. Model
#####  model structs
 
### 6. Routers
#### Resources define the routes of the project

### 7. Storage
#### It is generally for storage purpose.
### Database
#### 8. contains the database functions
