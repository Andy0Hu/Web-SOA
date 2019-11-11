# Web-SOA

## Language, Database and Environment

### Language and Databse

The program languages are `Java` and `Golang`.

The Databases are `MySQL` and `mongoDB`.

### Environment

#### Golang

 Version 1.13.4, prebuilt, the ELF can be run directly in Linux or macOS.

#### Java

**TODO**

#### MySQL

**TODO**

#### mongoDB

Version 4.2.1

### API Introduction

#### auth

This part is coded with `Golang`. We use JWT as a way for securely transmitting information between parties.

the database used is mongoDB.

1. Login

   1. method: **POST** 

   2. url: `api/v1/auth/sessions`

   3. intro:

       Function binds id and password in the Body of request. It checks the password and return the result of this login operation. It also returns the token generated as response.

2. Register

   1. method: **POST**

   2. url: `api/v1/auth/users`

   3. intro:

      Function binds id, password and username inthe Body of request. It checks whether this id is in the database and decides to give the permission to the user who is going to register, then insert the user messages to the database.

