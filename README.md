# warung-pintar

I use existing libs :

 - Chi Router
 - Ozzo Validation, for input request validation
 - Godotenv, for env loader
 - Gorm, for ORM

# Solution must consist of a minimum 3 microservices inside one repository (monorepo).
I have no idea what service I have to build, I apologize in advance, because suddenly I got an urgent task at work, I tried to make some simple services as follows:
> checkout-service
- checkout service is a main service here. this service will return the summary of prices
> item-service
- this service will return the products
> promo-service
- this service will return the promos of the products

# sequence diagram
the flow explained in warung-pintar.png

# For setup after cloning the repo:
- cd warung-pintar
- cd checkout-service, go mod tidy
- cd item-service, go mod tidy
- cd promo-service, go mod tidy

# to do a unit test :
- go to the package you want to testing then run a command "go test"
- you can see the coverage testing in each package by open the project with vscode, choose the testing file, right click then choose "Go:Toogle Test Coverage in Current Package"

# summary of unit test 
I have done the unit test and here are the result :
> Running tool: /usr/local/go/bin/go test -timeout 30s -coverprofile=/var/folders/h_/tjhvlj3n32sc9lvvfbm8x9ym0000gn/T/vscode-goeBpI0D/go-code-cover warung-pintar/checkout-service/src/app/use_cases/checkout
>> ok  	warung-pintar/checkout-service/src/app/use_cases/checkout	0.673s	coverage: 100.0% of statements

> Running tool: /usr/local/go/bin/go test -timeout 30s -coverprofile=/var/folders/h_/tjhvlj3n32sc9lvvfbm8x9ym0000gn/T/vscode-goeBpI0D/go-code-cover warung-pintar/checkout-service/src/interface/rest/handlers/checkout
>> ok  	warung-pintar/checkout-service/src/interface/rest/handlers/checkout	0.665s	coverage: 100.0% of statements

> Running tool: /usr/local/go/bin/go test -timeout 30s -coverprofile=/var/folders/h_/tjhvlj3n32sc9lvvfbm8x9ym0000gn/T/vscode-goeBpI0D/go-code-cover warung-pintar/item-service/src/app/use_cases/item
>> ok  	warung-pintar/item-service/src/app/use_cases/item	0.627s	coverage: 100.0% of statements

> Running tool: /usr/local/go/bin/go test -timeout 30s -coverprofile=/var/folders/h_/tjhvlj3n32sc9lvvfbm8x9ym0000gn/T/vscode-goeBpI0D/go-code-cover warung-pintar/item-service/src/interface/rest/handlers/item
>> ok  	warung-pintar/item-service/src/interface/rest/handlers/item	0.746s	coverage: 100.0% of statements

> Running tool: /usr/local/go/bin/go test -timeout 30s -coverprofile=/var/folders/h_/tjhvlj3n32sc9lvvfbm8x9ym0000gn/T/vscode-goeBpI0D/go-code-cover warung-pintar/promo-service/src/app/use_cases/item
>> ok  	warung-pintar/promo-service/src/app/use_cases/item	0.605s	coverage: 100.0% of statements

> Running tool: /usr/local/go/bin/go test -timeout 30s -coverprofile=/var/folders/h_/tjhvlj3n32sc9lvvfbm8x9ym0000gn/T/vscode-goeBpI0D/go-code-cover warung-pintar/promo-service/src/interface/rest/handlers/item
>> ok  	warung-pintar/promo-service/src/interface/rest/handlers/item	0.664s	coverage: 100.0% of statements


# for db table :
> I use Postgresql for DB
>> in folder db, there are 2 .sql files with the create table command and insert command. you can run the command in your sql editor page.

# to running the project
after clone and do some set up that explained before, do this following actions :
- set database credential in warung-pintar/.env
- set database credential in each service
item-service :
DB_HOST=database (recommend to literally use "database" according to the docker-compose.yaml)
DB_PORT=5432  
DB_NAME=projek/your_db_name
DB_USERNAME=your_db_user
DB_PASSWORD=your_db_password
DB_SCHEMA=warung_pintar
DB_SSL_MODE=disable

promo-service :
DB_HOST=database (recommend to literally use "database" according to the docker-compose.yaml)
DB_PORT=5432  
DB_NAME=projek/your_db_name
DB_USERNAME=your_db_user
DB_PASSWORD=your_db_password
DB_SCHEMA=warung_pintar
DB_SSL_MODE=disable

- cd warung-pintar, docker-compose up
- go to you postgresql db editor (pgAdmin, etc)
- make a new connection to 0.0.0.0
- make a new database, in this project I make a db named "projek"
- in db "projek" make a new schema named "warung-pintar"
- do all command to make the table and insert, you can see the command in db/items.sql and db/promo.sql
- project ready to use

# the endpoint
> here is the curl for the endpoint :
- checkout
> curl --location --request GET 'http://localhost:8080/api/ping'
> curl --location --request GET 'http://localhost:8080/api/health'
> curl --location --request POST 'http://localhost:8080/api/checkout' \
--header 'x-api-key: warung-pintar' \
--header 'Content-Type: application/json' \
--data-raw '{
    "data": [
        {
            "sku": "120P90",
            "qty": 6
        }
    ] 
}'

- item list
curl --location --request POST 'http://localhost:8081/api/items' \
--header 'x-api-key: warung-pintar' \
--header 'Content-Type: application/json' \
--data-raw '{
    "data": [
        {
            "sku": "120P90"
        }
    ] 
}'

- promo list
curl --location --request POST 'http://localhost:8082/api/promos' \
--header 'x-api-key: warung-pintar' \
--header 'Content-Type: application/json' \
--data-raw '{
    "data": [
        {
            "sku": "120P90"
        }
    ] 
}'
> by the way, I really apologize, due to my urgent task at my work that I explained before, I just use POST method for all endpoints so the request can be write on the request body. I do this, so I can focus on speed in delivery. once again i really apologize.

> i use x-api-key in header, for now you can fill it with anything in example I fill it with "warung-pintar". it just an example if in assumsion the endpoint using a header middleware

> here is the postman link if you want to use postman instead : 
> https://www.getpostman.com/collections/bb13471853f252e7e3a9
