# Pointless Microservices

## date

Run date server:

    $ go run date.go

Retrieve today's date:

    $ curl localhost:3112/date

## time

Run time server:

    $ go run time.go

Retrieve the current time:

    $ curl localhost:2359/time

## dice

Run dice server:

    $ go run dice.go

Roll the dice:

    $ curl localhost:6666/dice

## isarbaloonline

Run isarbaloonline server:

    $ go run isarbaloonline.go

Find out whether or not `http://arbalo.ch` is already online:

    $ curl localhost:1234/isarbaloonline
