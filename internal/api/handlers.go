package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/imzadev/nudl/internal/core"
)

func (a *App) List(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")
	jscontent := core.GetAllByKey(a.content, key)
	resp,err := json.Marshal(jscontent)	
	if err != nil {
		fmt.Println(err)
	}

	w.Write(resp)
}

func (a *App) GetByID(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")
	id := chi.URLParam(r, "id")
	convertedId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("error converitng id to number", err)
	}
	content := core.GetElementById(a.content, key, int(convertedId))
	resp, err := json.Marshal(content)
	if err != nil {
		fmt.Println(err)
	}
	
	w.Write(resp)
}
