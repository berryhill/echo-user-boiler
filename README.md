# echo-user-boiler
## API

#### POST/user
Creates a User with a form
Data:
+username {string}
+password {string}
#### POST/login
Logs a User in with a form
Data:
+username {string}
+password {string}
#### GET/user/:username
Gets a User with the Restful Parameter
+username {string{
