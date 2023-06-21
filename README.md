#### GET /todos
This request is used to retrieve all todos. The response example below indicates a successful request (status: 200) and returns an array of three todo items.
```json
{
    "status": 200,
    "todos": [
        {
            "id": "1",
            "item": "Clean Room",
            "completed": false
        },
        {
            "id": "2",
            "item": "Read Book",
            "completed": false
        },
        {
            "id": "3",
            "item": "Record Video",
            "completed": false
        }
    ]
}
```

#### GET /todos/2
This request is used to retrieve the details of a specific todo item with an ID of 2. The response example below indicates a successful request and returns the details of the todo item.
```json
{
    "status": 200,
    "todo": {
        "id": "2",
        "item": "Read Book",
        "completed": false
    }
}
```
#### POST /todos body
This request is used to create a new todo item. The request body should contain the details of the todo item. The example below represents a new todo item with an ID of 4 and the description "buy bread".
```json 
{
    "id": "4",
    "item": "buy bread",
    "completed": false
}
```

#### POST /todos 
This request creates a new todo item and returns its details. The response example below indicates a successful request (status: 201) and returns the details of the newly created todo item.
```json 
{
    "status": 201,
    "todo": {
        "id": "4",
        "item": "buy bread",
        "completed": false
    }
}
```
#### for all bad request and returns http.StatusBadRequest
For any other incorrect requests, the API responds with an HTTP 400 Bad Request status. The example below represents the response for a bad request.
```json 
{
    "status": 400
}
```
