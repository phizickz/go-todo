# run templ generation in watch mode to detect all .templ files and 
# re-create _templ.txt files on change, then send reload event to browser. 
# Default url: http://localhost:7331
live/templ:
	templ generate --watch --proxy="http://localhost:8080" --open-browser=false -v

# run air to detect any go file changes to re-build and re-run the server.
live/server:
	go run github.com/cosmtrek/air@v1.51.0 \
	--build.cmd "go build -o tmp/bin/main" --build.bin "tmp/bin/main" --build.delay "100" \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true cmd/main.go

# watch for any js or css change in the assets/ folder, then reload the browser via templ proxy.
live/sync_assets:
	go run github.com/cosmtrek/air@v1.51.0 \
	--build.cmd "templ generate --notify-proxy" \
	--build.bin "true" \
	--build.delay "100" \
	--build.exclude_dir "" \
	--build.include_dir "assets" \
	--build.include_ext "js,css"

live/db:
	docker run --name go-app-db \
		-e POSTGRES_USER=go_user \
		-e POSTGRES_PASSWORD=go_password \
		-e POSTGRES_DB=go_app_db \
		-p 5432:5432 \
		-d postgres:16

# Stop the PostgreSQL database
live/db_stop:
	docker stop go-app-db || true
	docker rm go-app-db || true

# Rebuild the database (useful for migrations or clean state testing)
live/db_reset: live/db_stop live/db
# start all 5 watch processes in parallel.
live: 
	make -j4 live/templ live/server live/sync_assets live/db_reset
