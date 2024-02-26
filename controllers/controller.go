package controllers

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/models"
	"html/template"
	"log"
	"net/http"
)

var ApiArtistes = "https://groupietrackers.herokuapp.com/api/artists"

var ApiLocations = "https://groupietrackers.herokuapp.com/api/locations"

var ApiDates = "https://groupietrackers.herokuapp.com/api/dates"

var ApiRelations = "https://groupietrackers.herokuapp.com/api/relation"

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	artistsData, err := loadArtistsFromAPi()
	fmt.Println(artistsData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tmpl, err := template.ParseFiles("templates/html/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//fmt.Println(artistsData)
	err = tmpl.Execute(w, artistsData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func HandleArtistes(w http.ResponseWriter, r *http.Request) {
	artistsData, err := loadArtistsFromAPi()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tmpl, err := template.ParseFiles("templates/html/artistes.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, artistsData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func HandleLieux(w http.ResponseWriter, r *http.Request) {
	artistsData, err := loadArtistsFromAPi()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tmpl, err := template.ParseFiles("templates/html/lieux.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, artistsData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func HandleDates(w http.ResponseWriter, r *http.Request) {
	artistsData, err := loadArtistsFromAPi()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tmpl, err := template.ParseFiles("templates/html/dates.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, artistsData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func HandleApropos(w http.ResponseWriter, r *http.Request) {
	artistsData, err := loadArtistsFromAPi()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tmpl, err := template.ParseFiles("templates/html/apropos.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, artistsData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/html/404.html")
	// Utilisez models.Entreprises pour accéder à la variable
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("Erreur lors de l'exécution du template: %v", err)
		return
	}
}

func loadDataFromAPi() (map[string]string, error) {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		fmt.Println("Erreur lors du chargement des données :", err)
		return nil, err
	}
	defer res.Body.Close()

	var data map[string]string
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		fmt.Println("Erreur lors du décodage des données :", err)
		return nil, err
	}
	return data, nil
}

func loadArtistsFromAPi() (interface{}, error) {
	res, err := http.Get(ApiArtistes)
	if err != nil {
		fmt.Println("Erreur lors du chargement des données :", err)
		return nil, err
	}
	defer res.Body.Close()

	var artists []models.Artist
	err = json.NewDecoder(res.Body).Decode(&artists)
	if err != nil {
		fmt.Println("Erreur lors du décodage des données :", err)
		return nil, err
	}
	return artists, nil
}

// func loadDatesFromAPi(url string) (interface{}, error) {
// 	res, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println("Erreur lors du chargement des données :", err)
// 		return nil, err
// 	}
// 	defer res.Body.Close()

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println("Erreur lors de la lecture des données :", err)
// 		return nil, err
// 	}
// 	var data map[string]interface{}
// 	err = json.Unmarshal(body, &data)
// 	if err != nil {
// 		fmt.Println("Erreur lors du décodage des données :", err)
// 		return nil, err
// 	}
// 	fmt.Println(data)
// 	return data, nil
// }
