package main

// 这里可能不太好懂，但已经写好案例了，可以去案例理解
const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}`

// |可以理解为占位符嘛，换了个顺序而已
