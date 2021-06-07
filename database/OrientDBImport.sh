#!/bin/sh

start="$(date -u +%s)"

## Relational
./oetl.sh customers.json
./oetl.sh products.json
./oetl.sh venders.json
./oetl.sh feedbacks.json

## JSON
./oetl.sh orders.JSON

elapsed2="$(($end_time2-$end_time1))"

## Graph
./oetl.sh cinps.json
./oetl.sh pknowps.json


end="$(date -u +%s)"
cost="$(($end-$start))"

echo "Data importing cost $cost sec"
