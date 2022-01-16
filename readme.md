# ToDoList App
- Todo List App is an application used for task management
- A mini project RESTful API by Group 4 Alterra Immersive Program Batch 5


Run project with: 
```
go run main.go
```

## Stack-tech :dart:
- [x] RESTful API Using Go, Echo, Gorm, MySQL
- [x] AWS for service api

## Open Endpoints

Open endpoints require no Authentication.

* Create user : `POST /users/register`
* Login : `POST /users/login/`

## Endpoints that require Authentication

Closed endpoints require a valid Token to be included in the header of the request. A Token can be acquired from the Login view above.

### Current User related

Each endpoint manipulates or displays information related to the User whose Token is provided with the request:

- Get user data by ID : `GET /users/:id`
- Register user : `POST /users/:id`
- Update user data by ID : `PUT /users/:id`
- Delete user data by ID : `DELETE /users/:id`

### Task related

Each endpoint manipulates or displays information related to the Task whose Token is provided with the request:

- Get task data by ID : `GET /tasks/:id`
- Get task data by UserID : `GET /tasks/:id`
- Create task by user ID : `POST /tasks/:id`
- Update task by user ID : `PUT /tasks/:id`
- Delete task data by user ID : `DELETE /tasks/:id`

### Project related

Each endpoint manipulates or displays information related to the Project whose Token is provided with the request:

- Get task data by ID : `GET /projects/:id`
- Get task data by UserID : `GET /projects/:id`
- Create task by user ID : `POST /projects/:id`
- Update task by user ID : `PUT /projects/:id`
- Delete task data by user ID : `DELETE /projects/:id`


