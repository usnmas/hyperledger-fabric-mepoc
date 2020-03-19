`Query`

- curl http://130.195.10.173:8080/api/query/CAR4
- curl http://130.195.10.173:8080/api/queryallcars


`Update car owner (both below are same)`

- curl -d '{"owner":"ABC"}' -H "Content-Type: application/json" -X PUT http://130.195.10.173:8080/api/changeowner/CAR4
- curl --location --request PUT '130.195.10.173:8080/api/changeowner/CAR4' --header 'Content-Type: application/json' --data-raw '{"owner":"ZQQ"}'
