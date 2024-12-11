up:
	goose -dir internal/db/migrations mysql "root:@tcp(localhost:3306)/go_rbac_db" up
down:
	goose -dir internal/db/migrations mysql "root:@tcp(localhost:3306)/go_rbac_db" down
