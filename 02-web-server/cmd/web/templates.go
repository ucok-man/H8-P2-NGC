package main

import "fmt"

func BaseTMPL(child string) string {
	var Base = fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<title>Aveggers</title>
	</head>
	<body>
		%v
	</body>
	</html>
	`, child)

	return Base
}
