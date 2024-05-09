run:
	@wgo -file=.go -file=.templ -xfile=_templ.go npx tailwindcss -i ./tailwind.css -o static/styles.css :: templ generate :: go run .