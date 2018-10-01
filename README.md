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

## Deployment

### Heroku

[Heroku](https://excel-processor.herokuapp.com/)

### Docker

**Build Docker image**
```
docker build -t excel-processor .
```

**Start Docker container**
```
docker run -i -t -p 3000:3000 -v "$PWD":/go/src/github.com/berto/excel-processor excel-processor
```
