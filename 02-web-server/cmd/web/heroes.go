package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func (app *application) GetHeroes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	heroes, err := app.entity.Hero.GetAll()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var buff bytes.Buffer
	for _, hero := range heroes {
		list := fmt.Sprintf(`
			<ul>
				<h2>Name 	: %s</h2>
				<p>Universe : %s</p>
				<p>Skill 	: %s</p>
				<img src="%s" alt="%s" width="200px" height="200px"/>
			</ul>
		`,
			hero.Name,
			hero.Universe,
			hero.Skill,
			hero.ImageURL,
			fmt.Sprintf("image of %s", hero.Name),
		)

		buff.WriteString(list)
		buff.WriteString("\n<hr>")
	}

	fmt.Fprintf(w, "%v", BaseTMPL(buff.String()))
}
