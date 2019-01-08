package main

import "fmt"

func main() {
	name := "Todd McLeod"
	age := "36"

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	<h2>` + age + `</h2>
	</body>
	</html>
	`
	fmt.Println(tpl)
}
