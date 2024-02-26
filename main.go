package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func home(w http.ResponseWriter, r *http.Request) {
	// Charge le fichier HTML
	tmpl, err := template.ParseFiles("Templates/home.html")
	if err != nil {
		handleError(w, "error-500.html", http.StatusInternalServerError)
		return
	}
	// Rendre le template avec un modèle vide pour l'instant
	if err := tmpl.Execute(w, nil); err != nil {
		handleError(w, "error-500.html", http.StatusInternalServerError)
	}
	/*if r.URL.Path != "/" && r.URL.Path != "/ascii-art" {
	      handleError(w, "error-404.html", http.StatusNotFound)
	      return
	      /*w.WriteHeader(http.StatusNotFound) // le code d'état 404 (Not Found)
	      return
	  }
	  http.ServeFile(w, r, "Templates/index.html")*/
}

func partenariat(w http.ResponseWriter, r *http.Request) {
	// Charge le fichier HTML
	tmpl, err := template.ParseFiles("Templates/partenariat.html")
	if err != nil {
		handleError(w, "error-500.html", http.StatusInternalServerError)
		return
	}
	// Rendre le template avec un modèle vide pour l'instant
	if err := tmpl.Execute(w, nil); err != nil {
		handleError(w, "error-500.html", http.StatusInternalServerError)
	}
	//http.ServeFile(w, r, "Templates/guide.html")
}

func programme(w http.ResponseWriter, r *http.Request) {
	// Charge le fichier HTML
	tmpl, err := template.ParseFiles("Templates/programme.html")
	if err != nil {
		handleError(w, "error-500.html", http.StatusInternalServerError)
		return
	}
	// Rendre le template avec un modèle vide pour l'instant
	if err := tmpl.Execute(w, nil); err != nil {
		handleError(w, "error-500.html", http.StatusInternalServerError)
	}
	//http.ServeFile(w, r, "Templates/guide.html")
}

func blog(w http.ResponseWriter, r *http.Request) {
	// Charge le fichier HTML
	tmpl, err := template.ParseFiles("Templates/blog.html")
	if err != nil {
		handleError(w, "error-500.html", http.StatusInternalServerError)
		return
	}
	// Rendre le template avec un modèle vide pour l'instant
	if err := tmpl.Execute(w, nil); err != nil {
		handleError(w, "error-500.html", http.StatusInternalServerError)
	}
	//http.ServeFile(w, r, "Templates/guide.html")
}

func aboutus(w http.ResponseWriter, r *http.Request) {
	// Charge le fichier HTML
	tmpl, err := template.ParseFiles("Templates/aboutus.html")
	if err != nil {
		handleError(w, "error-500.html", http.StatusInternalServerError)
		return
	}
	// Rendre le template avec un modèle vide pour l'instant
	if err := tmpl.Execute(w, nil); err != nil {
		handleError(w, "error-500.html", http.StatusInternalServerError)
	}
	//http.ServeFile(w, r, "Templates/guide.html")
}

func main() {
	// Définit la route pour le fichier HTML
	http.HandleFunc("/", home)
	http.HandleFunc("/partenariat", partenariat)
	http.HandleFunc("/programme", programme)
	http.HandleFunc("/aboutus", aboutus)
	http.HandleFunc("/blog", blog)

	// Définir la gestion des fichiers statiques
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))
	// http.Handle("/templates/assets/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))
	// http.Handle("/Templates/assets/img", http.StripPrefix("/Templ/", http.FileServer(http.Dir("Templates"))))

	// Démarre le serveur sur le port 8080
	fmt.Println("Serveur démarré sur le port :http://localhost:8080")
	http.ListenAndServe(":8081", nil)
}

func handleError(w http.ResponseWriter, templateName string, statusCode int) {
	w.WriteHeader(statusCode)
	tmpl, err := template.ParseFiles("templates/" + templateName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
