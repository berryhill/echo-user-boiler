# echo-user-boiler
## API

#### POST/user
<<<<<<< HEAD
Creates a User with a form
<<<<<<< HEAD
JSON
=======
Creates a User with JSON
>>>>>>> 96fd4a69b5294fddd309bc5532551539fc358336
```
{
  "username": {string},
  "password": {string}
}
```

#### GET/user/:id
Gets a User with the Restful Parameter
+ username {string}
=======
+ username {string}
+ password {string}

#### POST/login
Logs a User in with a form
+ username {string}
+ password {string}
>>>>>>> efc7c7da2819ac8ca9c1582e2dc47805a44e045a

#### GET/user/name/:username
Gets a User with the Restful Parameter
+ username {string}

<<<<<<< HEAD
#### GET/me (not implemented yet)
Gets current User
+ username {string}

#### PUT/user/:id (not yet implemented)
Updates a User with the Restful Parameter
+ username {string}

#### GET/users
Gets all Users

#### POST/login
User Login Portal with JSON
```
{
  "username": {string},
  "password": {string}
}
```
Returns a TokenedUser JSON
```
{
  "username": {string},
  "id": {string},
  "token": {string}
}
```

#### DELETE/user/:id
Deletes a User with the Restful Parameter
+ id {string}
=======

#### GET/user/:id
Gets a User with the Restful Parameter
+ username {string}
>>>>>>> efc7c7da2819ac8ca9c1582e2dc47805a44e045a
