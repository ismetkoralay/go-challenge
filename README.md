# go-challenge

It is a simple application developed with golang and running on docker.

* First install [Docker Desktop](https://hub.docker.com/editions/community/docker-ce-desktop-mac/)
* Then pull the repository and run the command below at the root directory(where the docker-compose.yml file is)
    * docker-compose -f docker-compose.yml up --build

* To add a customer to database => curl -X POST -H "Content-Type: application/json"  -d '{"Email":"test@test.com", "Password":"12345"}'  http://localhost:5000/api/v1/customers
* To get a customer from db by id => curl http://localhost:8080/api/v1/customers/{id}
    
