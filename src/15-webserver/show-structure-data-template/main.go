package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}

/* prints:

<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>My Peeps</title>
</head>
<body>
<ul>

	<li>Gandhi</li>

	<li>MLK</li>

	<li>Buddha</li>

	<li>Jesus</li>

	<li>Muhammad</li>

</ul>
</body>
</html>

// from Todd McLeod: https://github.com/GoesToEleven/golang-web-dev
*/
