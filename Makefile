tw:
	@npx tailwindcss -i input.css -o static/css/tw.css --watch

dev:
	@templ generate -watch -proxy="http://localhost:4000" -open-browser=false -cmd="go run cmd/web/main.go"