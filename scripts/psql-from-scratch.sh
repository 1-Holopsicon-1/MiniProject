docker run  -e POSTGRES_USER=userTest -e POSTGRES_PASSWORD=pwTest -e POSTGRES_DB=test --name MiniProject -dp 1515:5432 postgres:latest