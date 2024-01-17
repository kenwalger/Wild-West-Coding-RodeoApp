### Rodeo App for Wild West Coding Livestream


[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

### Project Resources

+ [Postman](https://www.postman.com)
+ [MongoDB Atlas](https://www.mongodb.com/atlas/database) as the database platform
+ [MongoDB Compass](https://www.mongodb.com/products/compass) GUI for data exploration
+ [GoLand](https://www.jetbrains.com/go)
+ [GoSwagger](https://goswagger.io)
+ [GoDotEnv](https://github.com/joho/godotenv)
+ [Go JWT](https://github.com/golang-jwt/jwt)
+ [Gin Sessions]( https://github.com/gin-contrib/sessions)
+ [GoLang xid – Globally Unique ID Generator]( https://pkg.go.dev/github.com/rs/xid)
+ [Argon2 hashing](https://pkg.go.dev/golang.org/x/crypto/argon2)


#### `instance/.env` File Format

```
# MongoDB Atlas Configuration
MONGODB_URI=<<MongoDB Connection String>>
MONGODB_DATABASE=<<MongoDB database name>>
RODEO_COLLECTION=<<MongoDB collection name for rodeos>>
USERS_COLLECTION=<<MongoDB collection name for users>>

# Server Port
PORT=<<Port number>>

# Authorization Strings
JWT_SECRET=<<random 16 character string>> 
```

#### Swagger
Generate spec
`swagger generate spec -o ./swagger.json`

Start documentation server (port 53065 by default)
`swagger serve ./swagger.json`

Different format (port 64761 by default)
`swagger serve -F swagger ./swagger.json`

Generate spec in markdown
`swagger generate markdown -f ./swagger.json --output swagger.md`

### Recordings
#### Episode 1
Project setup

+ [YouTube](https://www.youtube.com/watch?v=_BFUo-nQ3dE&list=PL2k86RlAekM-15R1CeiACQDQ6imxFToIF&index=2)
+ [LinkedIn](https://www.linkedin.com/events/wildwestcoding7092169384751763456/comments/)

#### Episode 2
Version Control, struct tags, documentation, database connection

+ [YouTube](https://www.youtube.com/watch?v=jtVn8ObZbUo&list=PL2k86RlAekM-15R1CeiACQDQ6imxFToIF&index=1&t=1745s)
+ [LinkedIn](https://www.linkedin.com/events/parsing-integrations7092232101378338816/comments/)

#### Episode 3
Adding a database and debugging

+ [YouTube]( https://www.youtube.com/watch?v=9bCvgMmJ97s)
+ [LinkedIn]( https://www.linkedin.com/events/wwc3-databasestoapisecurity7092272450637332481/comments/)

#### Episode 4
API Routes

+ [YouTube](https://www.youtube.com/watch?v=nsYZB5jamMw)
+ [LinkedIn](https://www.linkedin.com/events/wwc4-apiroutesanddocumentation7105588023962058752/comments/)

#### Episode 5
Finish API endpoints, implement JWT authorization

+ [YouTube]( https://www.youtube.com/watch?v=oYqZSAlTPs4&list=PL2k86RlAekM-15R1CeiACQDQ6imxFToIF&index=2)  
+ [LinkedIn]( https://www.linkedin.com/events/securingtheapi7112464363583676418/comments/)

#### Episode 6
Users & Password Hashing
+ [YouTube](https://www.youtube.com/watch?v=G5AMUFErcgw&list=PL2k86RlAekM-15R1CeiACQDQ6imxFToIF&index=7)  
+ [LinkedIn](https://www.linkedin.com/events/users-passwordhashing7120080186661900288/comments/)

#### Episode 7
Inedible Cookies
+ [YouTube](https://www.youtube.com/watch?v=anqXlL4EkVc&list=PL2k86RlAekM-15R1CeiACQDQ6imxFToIF&index=1&t=1s)
+ [LinkedIn](https://www.linkedin.com/events/inediblecookies7122360676488019969/comments/)

#### Episode 8
Templates & Data Passing
+ [YouTube]()
+ [LinkedIn](https://www.linkedin.com/events/7137830678917746688/comments/)

#### Episode 9
Year End Celebration - **NOT** a Rodeo App episode, but still fun.
+ [YouTube](https://www.youtube.com/watch?v=C3hkqD1kmb8&t=1s)
+ [LinkedIn](https://www.linkedin.com/events/yearendcelebration-wwc20237139034900170493952/comments/)