package cms

import (
	"net/http"
	"time"
)

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title: "Go Project",
		Content: "Hello there!",
		Posts: []*Post{
			&Post{
				Title: "First post",
				Content: "This is my first very interresting post",
				DatePublished: time.Now(),
			},
		},
	}

	Tmpl.ExecuteTemplate(w, "page", p)
}
