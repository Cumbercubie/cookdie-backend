#### Progress

08/19/25:
- Create restaurant servcices & API
- TODO:
    - Finish sqlc gen for create & get restaurant, finish restaurant schema
    - Follow https://www.udemy.com/course/backend-master-class-golang-postgresql-kubernetes/learn/lecture/25822346#overview
    
Structure:
```
yelp-sim/
├── README.md
├── docker-compose.yml
├── .env
├── Makefile
│
├── services/
│   ├── api-gateway/
│   │   ├── cmd/
│   │   │   └── main.go
│   │   ├── internal/
│   │   │   ├── handlers/
│   │   │   ├── routes/
│   │   │   └── middleware/
│   │   ├── Dockerfile
│   │   └── go.mod
│
│   ├── user-service/
│   │   ├── cmd/
│   │   ├── internal/
│   │   ├── migrations/
│   │   ├── Dockerfile
│   │   └── go.mod
│
│   ├── business-service/
│   │   ├── cmd/
│   │   ├── internal/
│   │   ├── migrations/
│   │   ├── Dockerfile
│   │   └── go.mod
│
│   ├── event-service/
│   │   ├── cmd/
│   │   ├── internal/
│   │   │   ├── producer/
│   │   │   ├── consumer/
│   │   │   └── models/
│   │   ├── Dockerfile
│   │   └── go.mod
│
├── simulator/
│   ├── load-generator/
│   │   ├── main.py
│   │   └── requirements.txt
│   └── traffic-models/
│       └── zipf.py
│
├── libs/
│   ├── logger/
│   ├── config/
│   ├── db/
│   └── messaging/
│
├── infra/
│   ├── kafka/
│   │   └── docker-compose.yml
│   ├── postgres/
│   └── redis/
│
└── scripts/
    ├── seed_data.sh
    ├── migrate.sh
    └── start_local.sh

```