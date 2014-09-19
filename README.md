weather
=======

An attempt at simplifying https://gist.github.com/kachayev/21e7fe149bc5ae0bd878

In this example, we implement ```weather.FindAll(cities)``` to find the weather in the specified cities:

* calls are made in parallel
* weather for each city is queried twice to reduce overall latency
* #FindAll returns once we have weather for each city
* upon completion, any outstanding calls are actively canceled

It would be interesting to see the corresponding implementation in another language.

## Running the Test

Executes the test and instructs the application to log all messages so you can see what it's doing internally.

```
export DEBUG='*'
go test -run=TestLive
```

## Sample Test Run

```
$ go test -run=TestLive
06:30:26.944 377us  377us  weather - GetWeather(San Francisco)
06:30:26.944 24us   24us   weather - GetWeather(San Francisco)
06:30:26.944 6us    6us    weather - GetWeather(Oakland)
06:30:26.944 5us    5us    weather - GetWeather(Oakland)
06:30:26.944 5us    5us    weather - GetWeather(Detroit)
06:30:26.944 5us    5us    weather - GetWeather(Detroit)
06:30:26.944 5us    5us    weather - GetWeather(London)
06:30:26.944 5us    5us    weather - GetWeather(London)
06:30:27.122 178ms  178ms  openweathermap - http - ok
06:30:27.122 178ms  178ms  weather - received weather => San Francisco
06:30:27.122 9us    9us    weather - weather reports received => 1
06:30:27.127 4ms    4ms    openweathermap - http - ok
06:30:27.127 4ms    4ms    weather - received weather => Detroit
06:30:27.127 5us    5us    weather - weather reports received => 2
06:30:27.128 1ms    1ms    openweathermap - http - ok
06:30:27.128 1ms    1ms    weather - received weather => London
06:30:27.128 3us    3us    weather - weather reports received => 3
06:30:27.128 72us   72us   openweathermap - http - ok
06:30:27.128 68us   68us   weather - received weather => San Francisco
06:30:27.128 2us    2us    weather - weather reports received => 3
06:30:27.130 1ms    1ms    openweathermap - http - ok
06:30:27.130 1ms    1ms    weather - received weather => Oakland
06:30:27.130 3us    3us    weather - weather reports received => 4
06:30:27.130 4us    4us    weather - finished
map[San Francisco:0xc20801a730 Detroit:0xc20801a910 London:0xc20801aaf0 Oakland:0xc20801a7d0]
06:30:27.130 46us   46us   openweathermap - http - CancelRequest
06:30:27.130 3us    3us    openweathermap - http - CancelRequest
06:30:27.130 2us    2us    openweathermap - http - CancelRequest
PASS
ok  	github.com/savaki/weather	0.195s
```