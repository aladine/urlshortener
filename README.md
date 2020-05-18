# URL Shortener Service

# Technical detail
- Database: connect with hosted postgres cloud service.
- Testing: one integration test for 1 endpoint
- Third party libraries used: 
   + "github.com/gorilla/mux": golang router
   + "github.com/go-pg/pg/v10": postgres database connector

# Assumptions
- If many users submit the same url, we could get different shorten url values. If the same user submit the same url, he should get an error that the url is already submitted(not implemented).
- From the example from the problem description, I assume the shorten url length unique value has 4 chars. And the unique hash need to 
- There will be a limit of the number of urls all the users can submit. If we choose the hash in range 0-9a-z then the maximum limit is N=36^4. To increase the limit we need to change the design for >=6 char instead of 4.

# API
- GET /: return the list of url and shorten urls by user id. In this project userID is hard code = 2.
- POST /url: submit a url and expect a shortened url return with status code 201
- GET /url/[:hash]: when user clicks the shortened link, it should redirect to the original URL
