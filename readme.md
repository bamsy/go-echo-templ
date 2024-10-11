# Go Templ/Htmx Base Project

My base project for using Go, Echo, Templ, HTMX, and PicoCSS

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Environment](#environment)

## Installation

- Install and setup [wgo](https://github.com/bokwoon95/wgo)
- Install templ cli ```go install github.com/a-h/templ/cmd/templ@latest```
- run ```go mod download```
- run ```make run```

## Usage

This project is a base project using the following technologies:

- [GO](https://go.dev/)
- [Echo](https://echo.labstack.com/)
- [Templ](https://templ.guide/)
- [PicoCSS](picocss.com)

After running the project you can browse to ```localhost:8080``` to see a Landing page

## Environment

The following environment variables are used in this project:

| Variable | Description |
|----------|-------------|
| `DEV_MODE` | Set this to "true" to get the browser to automatically refresh on changes after running `make run`. |


