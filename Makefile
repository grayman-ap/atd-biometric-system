postgres:
	docker run --name attendance-db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=studentsecret -d postgres

createdb:
	docker exec -it attendance-db createdb --username=root --owner=root student_attendance

dropdb: 
	docker exec -it attendance-db dropdb student_attendance

migrateup:
	migrate -path db/migration  -database "postgresql://root:studentsecret@localhost:5432/student_attendance?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:studentsecret@localhost:5432/student_attendance?sslmode=disable" -verbose down

sqlc:
	sqlc generate	

mod:
	go mod init github.com/grayman-ap/student_attendance

.PHONY: postgres createdb dropdb migrateup migratedown sqlc mod