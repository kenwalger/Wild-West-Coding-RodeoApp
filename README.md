### Rodeo App for Wild West Coding Livestream


[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
![GitHub Repo stars](https://img.shields.io/github/stars/kenwalger/Wild-West-Coding-RodeoApp)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/kenwalger/Wild-West-Coding-RodeoApp)
![GitHub last commit (branch)](https://img.shields.io/github/last-commit/kenwalger/Wild-West-Coding-RodeoApp/main)

### Project Resources

+ [Postman](https://www.postman.com)
+ [MongoDB Atlas](https://www.mongodb.com/atlas/database) as the database platform
+ [MongoDB Compass](https://www.mongodb.com/products/compass) GUI for data exploration
+ [GoLand](https://www.jetbrains.com/go) IDE for Go
+ [GoSwagger](https://goswagger.io)
+ [GoDotEnv](https://github.com/joho/godotenv)
+ [Go JWT](https://github.com/golang-jwt/jwt)
+ [Gin Sessions]( https://github.com/gin-contrib/sessions)
+ [GoLang xid – Globally Unique ID Generator]( https://pkg.go.dev/github.com/rs/xid)
+ [Argon2 hashing](https://pkg.go.dev/golang.org/x/crypto/argon2)


#### `instance/.env` File Format
The application relies on environment variables for the database connection,
and JWT generation. These variables are accessed through the GoDotEnv package
(see above for link) and stored in a `.env` file in an `instance` directory.
The format is as follows:

```text
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

### Episodes & Recording Links

| Episode Number | Episode Topics                                                                                                                           | YouTube                                                                                                     | LinkedIn | Application      |
|----------------|------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------|----------|------------------|
| 1              | Project Setup                                                                                                                            | [Link](https://www.youtube.com/watch?v=_BFUo-nQ3dE&list=PL2k86RlAekM-15R1CeiACQDQ6imxFToIF&index=2)         | [Link](https://www.linkedin.com/events/wildwestcoding7092169384751763456/comments/) | Rodeo App        |
| 2              | Version Control, struct tags, documentation, database connection                                                                         | [Link](https://www.youtube.com/watch?v=jtVn8ObZbUo&list=PL2k86RlAekM-15R1CeiACQDQ6imxFToIF&index=1&t=1745s) | [Link](https://www.linkedin.com/events/parsing-integrations7092232101378338816/comments/) | Rodeo App        |
| 3              | Adding a database and debugging                                                                                                          | [Link](https://www.youtube.com/watch?v=9bCvgMmJ97s)                                                         | [Link](https://www.linkedin.com/events/wwc3-databasestoapisecurity7092272450637332481/comments/) | Rodeo App        |
| 4              | API Routes                                                                                                                               | [Link](https://www.youtube.com/watch?v=nsYZB5jamMw)                                                         | [Link](https://www.linkedin.com/events/wwc4-apiroutesanddocumentation7105588023962058752/comments/) | Rodeo App        |
| 5              | Finish API endpoints, implement JWT authorization                                                                                        | [Link](https://www.youtube.com/watch?v=oYqZSAlTPs4&list=PL2k86RlAekM-15R1CeiACQDQ6imxFToIF&index=2)         | [Link](https://www.linkedin.com/events/securingtheapi7112464363583676418/comments/) | Rodeo App        |
| 6              | Users & Password Hashing                                                                                                                 | [Link](https://www.youtube.com/watch?v=G5AMUFErcgw&list=PL2k86RlAekM-15R1CeiACQDQ6imxFToIF&index=7)         | [Link](https://www.linkedin.com/events/users-passwordhashing7120080186661900288/comments/) | Rodeo App        |
| 7              | Inedible Cookies                                                                                                                         | [Link](https://www.youtube.com/watch?v=anqXlL4EkVc&list=PL2k86RlAekM-15R1CeiACQDQ6imxFToIF&index=1&t=1s)    | [Link](https://www.linkedin.com/events/inediblecookies7122360676488019969/comments/) | Rodeo App        |
| 8              | Templates & Data Passing                                                                                                                 | [Link](https://www.youtube.com/watch?v=8lnss4xnpwY&list=PL2k86RlAekM-15R1CeiACQDQ6imxFToIF&index=11)        | [Link](https://www.linkedin.com/events/7137830678917746688/comments/) | Rodeo App        |
| 9              | Year End Celebration                                                                                                                     | [Link](https://www.youtube.com/watch?v=C3hkqD1kmb8&t=1s)                                                    | [Link](https://www.linkedin.com/events/yearendcelebration-wwc20237139034900170493952/comments/) | None - Just Fun  |
| 10             | User Login & Logout Page & Routes                                                                                                        | [Link](https://www.youtube.com/watch?v=guGxaRcozvM&list=PL2k86RlAekM-15R1CeiACQDQ6imxFToIF)                 | [Link](https://www.linkedin.com/events/7153155292778647554/comments/) | Rodeo App        |
| 11             | The AI Surge in 2024: Reshaping Developer and DevOps with guest [John Capobianco](https://www.linkedin.com/in/john-capobianco-644a1515/) | [Link](https://www.youtube.com/watch?v=iaV6KDNQXEc)                                                         | [Link](https://www.linkedin.com/events/theaisurgein2024-reshapingdevel7150243689200537601/comments/) | None - Just Fun  |
| 12             | Adding Duo Authentication | Link - TBD                                                                                                  | Link - TBD | Rodeo App |
