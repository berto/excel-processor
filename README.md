# Excel Processor

Containerized go web server that processes an excel file(s) into a database.

## Requirements:

- Simple front-end with a drop-zone or file upload component
- Tests
- Readme
- Docker file and/or docker compose

## Extra credit:

- Process a zip file of excel files

## ERD

![erd](./erd.png)

## Usage

- GET '/': displays client with upload form for excel data (look inside `db/sample-data` for an example)
- GET '/migrate': truncates/migrates db. Run before seeding or uploading
- GET '/seed': seeds db with ship data inside `db/sample-data`
- GET '/budget/:code': display uploaded budget data from the db
- POST '/ship': route to parse excel data from form upload

## TODOs and Tracker

- Information about next features and bugs, check out the [Trello](https://trello.com/b/heHh1P0r/seaspan-ship-excel-processor) board.

## Deployment

### Heroku

[Heroku](https://excel-processor.herokuapp.com/)

### Docker

**Clone and navigate to repo**
```
go get -d github.com/berto/excel-processor
cd $GOPATH/src/github.com/berto/excel-processor
```

**Build Docker image**
```
docker build -t excel-processor .
```

**Start Docker container**
```
docker run -i -t -p 3000:3000 -v "$PWD":/go/src/github.com/berto/excel-processor excel-processor
```

**Build Application**
```
go build
```

**Run Application**
```
./excel-processor
```
