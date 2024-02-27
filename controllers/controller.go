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
	artistesChanel := make(chan []models.Artist)
	errorCh := make(chan error)
	locationsChanel := make(chan []models.Location)
	datesChanel := make(chan []models.Date)
	relationsChanel := make(chan []models.Relation)

	go loadArtistsFromAPi(artistesChanel, errorCh)
	go loadLocationsFromAPi(locationsChanel)
	go loadDatesFromAPi(datesChanel)
	go loadRelationsFromAPi(relationsChanel)

	select {
	case artistsData := <-artistesChanel:
		select {
		case locationsData := <-locationsChanel:
			select {
			case datesData := <-datesChanel:
				select {
				case relationsData := <-relationsChanel:

					var artistLocationData []models.ArtistLocation

					for i, v := range artistsData {
						var artistLocation models.ArtistLocation
						artistLocation.Artist = v
						artistLocation.Location = locationsData[i]
						artistLocationData = append(artistLocationData, artistLocation)
					}

					tmpl, err := template.ParseFiles("templates/html/index.html")
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					err = tmpl.Execute(w, map[string]interface{}{"Artistes": artistsData, "Locations": locationsData, "Dates": datesData, "Relations": relationsData, "ArtisteLocation": artistLocationData})
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
				case err := <-errorCh:
					http.Error(w, "Erreur lors de l'appel à l'API: "+err.Error(), http.StatusInternalServerError)
					return
				}
			case err := <-errorCh:
				http.Error(w, "Erreur lors de l'appel à l'API: "+err.Error(), http.StatusInternalServerError)
				return
			}

		case err := <-errorCh:
			http.Error(w, "Erreur lors de l'appel à l'API: "+err.Error(), http.StatusInternalServerError)
			return
		}

	case err := <-errorCh:
		http.Error(w, "Erreur lors de l'appel à l'API: "+err.Error(), http.StatusInternalServerError)
		return
	}

}

func HandleArtistes(w http.ResponseWriter, r *http.Request) {
	artistesChanel := make(chan []models.Artist)
	errorCh := make(chan error)

	go loadArtistsFromAPi(artistesChanel, errorCh)
	artistsData := <-artistesChanel

	// artistesLocationsChanel :=  make(chan []models.Location)

	// go loadOneLocationFromAPi(artistesLocationsChanel,artistsData)

	locationsChanel := make(chan []models.Location)

	go loadLocationsFromAPi(locationsChanel)
	locationsData := <-locationsChanel

	datesChanel := make(chan []models.Date)

	go loadDatesFromAPi(datesChanel)
	datesData := <-datesChanel

	tmpl, err := template.ParseFiles("templates/html/artistes.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, map[string]interface{}{"Artistes": artistsData, "Locations": locationsData, "Dates": datesData})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func HandleLieux(w http.ResponseWriter, r *http.Request) {
	artistesChanel := make(chan []models.Artist)
	errorCh := make(chan error)

	go loadArtistsFromAPi(artistesChanel, errorCh)
	artistsData := <-artistesChanel

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

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
	artistesChanel := make(chan []models.Artist)
	errorCh := make(chan error)

	go loadArtistsFromAPi(artistesChanel, errorCh)
	artistsData := <-artistesChanel

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

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
	artistesChanel := make(chan []models.Artist)
	errorCh := make(chan error)

	go loadArtistsFromAPi(artistesChanel, errorCh)
	artistsData := <-artistesChanel

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

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

// func loadDataFromAPi() (map[string]string, error) {
// 	res, err := http.Get("https://groupietrackers.herokuapp.com/api")
// 	if err != nil {
// 		fmt.Println("Erreur lors du chargement des données :", err)
// 		return nil, err
// 	}
// 	defer res.Body.Close()

// 	var data map[string]string
// 	err = json.NewDecoder(res.Body).Decode(&data)
// 	if err != nil {
// 		fmt.Println("Erreur lors du décodage des données :", err)
// 		return nil, err
// 	}
// 	return data, nil
// }

func loadArtistsFromAPi(resultChan chan<- []models.Artist, errorCh chan<- error) (interface{}, error) {
	res, err := http.Get(ApiArtistes)
	if err != nil {
		errorCh <- err
		return nil, err
	}
	defer res.Body.Close()

	var artists []models.Artist
	err = json.NewDecoder(res.Body).Decode(&artists)
	if err != nil {
		errorCh <- err
		return nil, err
	}

	resultChan <- artists
	return artists, nil
}

func loadLocationsFromAPi(resultChan chan<- []models.Location) (interface{}, error) {
	res, err := http.Get(ApiLocations)
	if err != nil {
		fmt.Println("Erreur lors du chargement des données :", err)
		return nil, err
	}
	defer res.Body.Close()

	var result map[string][]models.Location
	var locations []models.Location
	err = json.NewDecoder(res.Body).Decode(&result)
	locations = result["index"]
	if err != nil {
		fmt.Println("Erreur lors du décodage des données :", err)
		return nil, err
	}
	resultChan <- locations
	return locations, nil
}

func loadOneLocationFromAPi(resultChan chan<- []models.Location, url string) (interface{}, error) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Erreur lors du chargement des données :", err)
		return nil, err
	}
	defer res.Body.Close()

	var locations []models.Location
	err = json.NewDecoder(res.Body).Decode(&locations)
	if err != nil {
		fmt.Println("Erreur lors du décodage des données :", err)
		return nil, err
	}
	resultChan <- locations
	return locations, nil
}

func loadDatesFromAPi(resultChan chan<- []models.Date) (interface{}, error) {
	res, err := http.Get(ApiDates)
	if err != nil {
		fmt.Println("Erreur lors du chargement des données :", err)
		return nil, err
	}
	defer res.Body.Close()

	var result map[string][]models.Date
	var dates []models.Date
	err = json.NewDecoder(res.Body).Decode(&result)
	dates = result["index"]
	if err != nil {
		fmt.Println("Erreur lors du décodage des données :", err)
		return nil, err
	}
	resultChan <- dates
	return dates, nil
}

func loadRelationsFromAPi(resultChan chan<- []models.Relation) (interface{}, error) {
	res, err := http.Get(ApiRelations)
	if err != nil {
		fmt.Println("Erreur lors du chargement des données :", err)
		return nil, err
	}
	defer res.Body.Close()

	var result map[string][]models.Relation
	var relations []models.Relation
	err = json.NewDecoder(res.Body).Decode(&result)
	relations = result["index"]
	if err != nil {
		fmt.Println("Erreur lors du décodage des données :", err)
		return nil, err
	}
	resultChan <- relations
	return relations, nil
}
