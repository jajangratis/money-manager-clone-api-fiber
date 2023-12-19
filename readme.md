
# Money Manager Clone API

Create api for money manager apps using fiber and postgreSQL


## Tech Stack


**Server:** 

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Nginx](https://img.shields.io/badge/nginx-%23009639.svg?style=for-the-badge&logo=nginx&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)

## Installation

Install in local with native golang or using docker compose

```bash
go mod download && go mod verify
go build -v -o -ldflags="-s -w" -o <YOUR_APP_NAME>
  ./<YOUR_APP_NAME>
```



Run the migration to generate the database you can create the database in  advance then run the migration with this command
```bash
migrate -database ${POSTGRESQL_URL} -path db/migrations up
```

## Deployment

To deploy this project run

```bash
docker-compose up --build
```


## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`DB_USERNAME`
`DB_PASSWORD`
`DB_HOST`
`DB_PORT`
`DB_SELECT`
`IP_HOST`
`ENVIROMENT``

`LOGIN_IV`
`KEY_LOGIN`
`KEY_JWT `

## Authors
![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)
- [@jajangratis](https://www.github.com/jajangratis)
  


## License

[MIT](https://choosealicense.com/licenses/mit/)


## Acknowledgements

- [Awesome Readme Templates](https://awesomeopensource.com/project/elangosundar/awesome-README-templates)
- [Awesome README](https://github.com/matiassingers/awesome-readme)
- [How to write a Good readme](https://bulldogjob.com/news/449-how-to-write-a-good-readme-for-your-github-project)

