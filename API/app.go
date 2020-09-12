package main
import (
	"encoding/json"

	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
	."./dao"
	."./config"
	."./models"

 
)


var config = Config{}
var dao = MoviesDAO{}

//GET
//ENDPOINT: http:localhost:8080/movies
func allMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
 //fmt.Fprintln(w, "Muestras todas las peliculas")
	movies, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, movies)
}

//GET
//ENDPOINT: http:localhost:8080/movies/{id}
//ENDPOINT: http:localhost:8080/movies/4
func findMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}
	respondWithJson(w, http.StatusOK, movie)
}

// POST
// ENDPOINT: http:localhost:8080/movies
func createMovieEndPoint(w http.ResponseWriter, r *http.Request) {
 //fmt.Fprintln(w, "Registra una película")
	defer r.Body.Close()
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	movie.ID = bson.NewObjectId()
	if err := dao.Insert(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, movie)
}

// PUT
// ENDPOINT: http:localhost:8080/movies
func updateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
 //fmt.Fprintln(w, "Actualiza una película")
	defer r.Body.Close()
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
	
}
// DELETE
// ENDPOINT: http:localhost:8080/movies/{id}
func deleteMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "Elimina una película segun id")
	defer r.Body.Close()
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})

}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}


func main() {
	r := mux.NewRouter()

	r.HandleFunc("/movies", allMoviesEndPoint).Methods("GET")
	r.HandleFunc("/movies/{id}", findMovieEndPoint).Methods("GET")
	r.HandleFunc("/movies", createMovieEndPoint).Methods("POST")
	r.HandleFunc("/movies", updateMovieEndPoint).Methods("UPDATE")
	r.HandleFunc("/movies/{id	}", deleteMovieEndPoint).Methods("DELETE")
	if err := http.ListenAndServe(":9000", r); err != nil {
		log.Fatal(err)
	}
}