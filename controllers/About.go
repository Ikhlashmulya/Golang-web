package controllers

import (
    "net/http"
    "html/template"
    "path"
    "log"
    "latihan-web/models"
)

func About(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles(path.Join("views", "about.html"), path.Join("views", "layout.html"))
    if err != nil {
        log.Println(err)
        http.Error(w, "Something Went Wrong", http.StatusInternalServerError)
        return
    }
    
    data := models.GetAllData()
    
    err = tmpl.Execute(w, data)
    if err != nil {
        log.Println(err)
        http.Error(w, "Something Went Wrong", http.StatusInternalServerError)
        return
    }
}