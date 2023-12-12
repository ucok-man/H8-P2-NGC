package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func (app *application) GetVillains(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	villains, err := app.entity.Villain.GetAll()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var buff bytes.Buffer
	for _, villain := range villains {
		list := fmt.Sprintf(`
			<ul>
				<h2>Name 	: %s</h2>
				<p>Universe : %s</p>
				<img src="%s" alt="%s" width="200px" height="200px"/>
			</ul>
		`,
			villain.Name,
			villain.Universe,
			villain.ImageURL,
			fmt.Sprintf("image of %s", villain.Name),
		)

		buff.WriteString(list)
		buff.WriteString("\n<hr>")
	}

	fmt.Fprintf(w, "%v", BaseTMPL(buff.String()))
}
