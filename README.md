# Project toggl-card-game

This is a coding challenge project given by Toggl.com hiring team.

## NOTE: 
The data is stored in memory and not persisted into a DB or FS.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

1) First clone the repository to your local machine.
2) Execute the following command to run api server locally:

```bash
make run
```
## Testing

test create new default deck endpoint
```bash
curl -X POST http://localhost:8080/api/deck
```

test create new full shuffled deck endpoint 
```bash
curl -X POST -G 'http://localhost:8080/api/deck' -d 'shuffle=true'
``` 

test create new deck with custom cards endpoint
```bash
curl -X POST -G 'http://localhost:8080/api/deck' -d 'cards=2C,3D,10H,KC'
```

test open existing deck endpoint (replace <deck_id> with actual deck id)
```bash
curl -X GET http://localhost:8080/api/deck/<deck_id>
```

test draw card from deck endpoint (replace <deck_id> with actual deck id and count number)
```bash
curl -X PUT http://localhost:8080/api/deck -d '{"deck_id": "<deck_id>", "count": 2}'
```


## Makefile Commands Description

build the application
```bash
make build
```

run the application
```bash
make run
```

run the test suite
```bash
make test
```

run test coverage
```bash
make test-cover
```

clean up binary from the last build
```bash
make clean
```
