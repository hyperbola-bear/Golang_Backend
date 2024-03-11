# Golang Backend API
This an API created using Golang for Govtech take-home assignment. I've included instructions below on how to run it.

## Pre-requisites 
1. Have Golang and Docker installed. 
2. Possess a copy of Makefile contents.

## set up

1. Git clone project into your IDE
```
> git clone https://github.com/hyperbola-bear/Golang_Backend
```
2. Then install dependencies needed in the project
```
> go mod tidy
```
3. Create a file named "Makefile" in root directory
4. Pull Postgres image to Docker
```
> docker pull postgres:16-alpine
```
5. install make (use scoop if using windows)
6. Run a docker container with postgres16
```
> make postgres
```
7. Create a Database in PSQL
```
> make createdb
```
8. migrate schema into PSQL
```
> make migrateup
```

## Testing 
1. Run the program which will be hosted on localhost:8080
```
> make server 
```
2. To run Test
```
> make test
```

## API Routes

### Register `GET`
registers student(s) to a teacher

`http://localhost:8080/api/register`

sample payload:
```
{
  "teacher": "teacherben@gmail.com",
  "students":
    [
      "studentron@gmail.com",
      "studenthon@gmail.com"
    ]
}
```
Expected result
```
204 No Content
```

### Common Students `GET`
Given a list of teacher(s) (given as query parameters), retrieve a list of students who
are registered to all the teachers in the list.

`http://localhost:8080/api/commonstudents?teacher=<teacher_name>&teacher=<teacher_name>`

sample payload:
```
http://localhost:8080/api/commonstudents?teacher=teacherben@gmail.com
```
Expected result
```
{
    "students":
    [
      "studentron@gmail.com",
      "studenthon@gmail.com"
    ]
}
```

sample payload:
```
http://localhost:8080/api/commonstudents?teacher=teacherben@gmail.com&teacher=mr@.com
```
Expected result
```
{
    "students": []
}
```

### Suspension `POST`
Suspend a student

`http://localhost:8080/api/suspend`

sample payload:
```
{
  "student": "student3"
}
```
Expected result
```
204 No Content
```

### Receive Notification `POST`
Send a notification from a teacher to students.
To receive notification:
- Students must *NOT* be suspended
  and 
- Registered to teacher
  or
- Mentioned in the notification with a @  


`http://localhost:8080/api/retrievefornotifications`

sample payload:
```
{
  "teacher":  "teacherben@gmail.com",
  "notification": "tomorrow exam@lecture venue @jay@email.com @student3"
}
```
Expected result
```
{
    "recipients": [
        "studentron@gmail.com",
        "studenthon@gmail.com",
        "jay@email.com"
    ]
}
```
student3 is not in the list of recipients because he is suspended.

