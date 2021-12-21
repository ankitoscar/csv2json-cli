# csv2json-cli
This is a CLI tool made in GO to convert a csv file to json  

## Build the tool 


Run the command :

`go build csv2json.go`


## Convert a file 

Run the following command : 

`.\csvjson <csvFile>`

### Flags 

1. *pretty* : To get indented and formatted JSON output. Default - `false`
2. *separator*: Mention whether the separator is a comma or a semicolon. Default - `comma`
