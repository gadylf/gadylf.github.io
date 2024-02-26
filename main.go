package main

import (
	"fmt"
	"groupie-tracker/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Création d'un routeur
	router := mux.NewRouter()

	router.HandleFunc("/", controllers.HandleIndex)
	router.HandleFunc("/artistes", controllers.HandleArtistes)
	router.HandleFunc("/lieux", controllers.HandleLieux)
	router.HandleFunc("/dates", controllers.HandleDates)
	router.HandleFunc("/apropos", controllers.HandleApropos)

	// Démarrage du serveur HTTP
	http.Handle("/", router)

	handleHtml := http.FileServer(http.Dir("templates/html"))
	http.Handle("/html/", http.StripPrefix("/html/", handleHtml))
	handleCss := http.FileServer(http.Dir("templates/css"))
	http.Handle("/css/", http.StripPrefix("/css/", handleCss))
	handleMedia := http.FileServer(http.Dir("templates/media"))
	http.Handle("/media/", http.StripPrefix("/media/", handleMedia))

	// Définition du gestionnaire d'erreur 404
	router.NotFoundHandler = http.HandlerFunc(controllers.NotFoundHandler)

	port := ":8080"

	fmt.Println("http://localhost:8080")
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Erreur lors de l'écoute du serveur:", err)
	}
}
