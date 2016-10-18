# echo-user-boiler
## API

#### POST/user
Creates a User with a form
JSON
+ username {string}
+ password {string}

#### GET/user/:id
Gets a User with the Restful Parameter
+ username {string}

#### GET/user/name/:username
Gets a User with the Restful Parameter
+ username {string}

#### PUT/user/:id
Updates a User with the Restful Parameter ':id' (not yet implemented)
+ username {string}

#### GET/users
Gets all Users

#### POST/login
Logs a User in with a form
JSON
+ username {string}
+ password {string}

#### DELETE/user/:id
Deletes a User with the Restful Parameter ':id'
+ id {string}
