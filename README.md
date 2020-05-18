Pick one of these high level specs and absolutely own it! Ideally you won't spend more than a few hours - your time is precious!

# URL Shortener Service
Your challenge is to create a URL shortener service to take a nice long url and turn it into something like qka.co/f00b - similar to how twitters t.co and bit.ly work.
This could be a good chance to do some frontend if you want (you can use any JavaScript framework/library on the front-end such as React or Vue) but the focus will be on how the service works so if you want to do full server side renders we won’t judge.
Your user should be able to:
Shorten any valid URL
See a list of their shortened links, even after refreshing the browser (database!)
Maybe copy the shortened link to their clipboard in a single click
Visit a shortened link and have it redirect them to the main page
Any other nice things you can think of :) 


# Assumptions
- If many users submit the same url, we could get different shorten url values. If the same user submit the same url, he should get an error that the url is already submitted.
- From the example from the problem description, I assume the shorten url length unique value has 4 chars. And the unique hash need to 
- There will be a limit of the number of urls all the users can submit. If we choose the hash in range 0-9a-z then the maximum limit is N=36^4. To increase the limit we need to change the design for >=6 char instead of 4.


# API
- GET /: return the list of url and shorten urls by user id. In this project userID is hard code = 2.
- POST /url: submit a url and expect a shortened url return with status code 201
- GET /url/[:hash]: when user clicks the shortened link, it should redirect to the original URL
