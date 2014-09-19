weather
=======

An attempt at simplifying https://gist.github.com/kachayev/21e7fe149bc5ae0bd878

In this example, we implement ```weather.FindAll(cities)``` to find the weather in the specified cities:

* calls are made in parallel
* weather for each city is queried twice to reduce overall latency
* #FindAll returns once we have weather for each city
* upon completion, any outstanding calls are actively canceled

It would be interesting to see the corresponding implementation in another language.

