# Getir coding assigment (case study)

This assignment required me to create a REST/JSON API and provide two endpoints:

* `/fetch` has only a POST method and its purpose consists of searching records in a MongoDB collection;
* `/in-memory` has a POST and a GET method, to insert into and read from an in-memory database;

The only constraint is not to use a web framework, therefore I use only the net/http package, from standard library.

## Architecture

`main.go` loads the app configuration, creates all the dependencies and wire them into the few objects needed for the project.
(in a real-world app I'd use google/wire or a custom dependencies' injection mechanism)
Then it sets up the routes to the two endpoint handlers, and runs an http server. 

The endpoint handlers can be found in the `api` package. Usually I don't do anything in handlers but parse the request, 
validate it, forward it to a service (in which the business logic is decoupled from the infrastructure concerns, such as a http API),
and eventually transform and send the results back to a client. 
In this case it sounded overkill to me, thus my handler uses directly facilities that are usually downstream, such as...

The `pkg` package contains anything that could be useful to other services, for example a DB reader (with a MongoDB implementation), 
a key-value storage (with a trivial implementation based on a hashmap), and other stuff.

## Documentation and... a quirk!

I do believe in the Agile Manifesto statement 'Working software over comprehensive documentation', though I really like 
to give hints to fellow developers on crucial parts of my work. For example, while developing the MongoDB integration for
this assignment, I stumbled upon an issue that's beyond the scope of the assigment itself (at least, I believe so).

If you look closely at `pkg/db/mongo/reader.go` you'll find a double implementation of the 'query' used by /fetch. 
One of them, uncommented and active, is awful: it just queries the records by date range, then filter out them by 
count range... manually! Why? The function `aggregateRecords` is instead the correct one, but can't be used for the 
assigment, since it relies on Mongo's aggregation framework and its `$sum` operator, which apparently is not available 
on the Atlas cluster provided on the case study's description (at least not in the current tier). 

Despite that unfortunate workaround, you will find my API docs at `/swagger`

## Tests

Sometimes I use TDD for core logic components, but in this case I spent too much time on writing the MongoDB query and 
other parts, thus I only added two unit tests: one for `api/fetch/handler.go` and the other for `pkg/kv/map_storage.go`.
The first one is an example of how to test a http handler and the second one shows how I usually write table-based tests.
I'm sorry ^^

## How to build and deploy

```
docker build -t getir-fabiobozzo .
docker run -p8080:8080 -e MONGODB_URI=<the URI> getir-fabiobozzo
```

And that's it. The API is available at `http://localhost:8080`.
There's also a live version deployed at `https://getir-fabiobozzo.herokuapp.com` (`/` isn't mapped, try `/swagger`)

Please contact me for any questions. :-)