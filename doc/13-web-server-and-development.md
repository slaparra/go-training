# Web development

## Templates

*Go’s [html/template] package provides a rich templating language for HTML templates. It is mostly used in web applications to display data in a structured way in a client’s browser. One great benefit of Go’s templating language is the automatic escaping of data.*

*Package [text/template] implements data-driven templates for generating textual output.  
To generate HTML output, see package [html/template], which has the same interface as this package but automatically secures HTML output against certain attacks.*

[text/template]: https://godoc.org/text/template
[html/template]: https://godoc.org/html/template

A data structure in our code:
```
data := TodoPageData{
	PageTitle: "My TODO list",
	Todos: []Todo{
		{Title: "Task 1", Done: false},
		{Title: "Task 2", Done: true},
		{Title: "Task 3", Done: true},
	},
}
```

some_tpl.gohtml:
```
<h1>{{.PageTitle}}<h1>
<ul>
    {{range .Todos}}
        {{if .Done}}
            <li class="done">{{.Title}}</li>
        {{else}}
            <li>{{.Title}}</li>
        {{end}}
    {{end}}
</ul>
```

## Links
- https://gowebexamples.com/templates/