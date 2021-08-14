package main

// Import our dependencies. We'll use the standard HTTP library as well as the gorilla router for this app
import (
	// "encoding/json"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

/* We will first create a new type called Product
   This type will contain information about VR experiences */
type Product struct {
	Id          int
	Name        string
	Slug        string
	Description string
}

/* We will create our catalog of VR experiences and store them in a slice. */
var products = []Product{
	{Id: 1, Name: "World of Authcraft", Slug: "world-of-authcraft", Description: "Battle bugs and protect yourself from invaders while you explore a scary world with no security"},
	{Id: 2, Name: "Ocean Explorer", Slug: "ocean-explorer", Description: "Explore the depths of the sea in this one of a kind underwater experience"},
	{Id: 3, Name: "Dinosaur Park", Slug: "dinosaur-park", Description: "Go back 65 million years in the past and ride a T-Rex"},
	{Id: 4, Name: "Cars VR", Slug: "cars-vr", Description: "Get behind the wheel of the fastest cars in the world."},
	{Id: 5, Name: "Robin Hood", Slug: "robin-hood", Description: "Pick up the bow and arrow and master the art of archery"},
	{Id: 6, Name: "Real World VR", Slug: "real-world-vr", Description: "Explore the seven wonders of the world in VR"},
}

func main() {
	// Here we are instantiating the gorilla/mux router
	r := mux.NewRouter()

	// On the default page we will simply serve our static index page.
	r.Handle("/", http.FileServer(http.Dir("./views/")))

	initialiseRoutes(r)

	// We will setup our server so we can serve static assest like images,
	// css from the /static/{file} route
	r.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Our application will run on port 8080. Here we declare the port and
	// pass in our router.
	http.ListenAndServe(":8080", r)
}

func initialiseRoutes(r *mux.Router) {

	// Status of server - check that API is working
	r.Handle("/status", StatusHandler).Methods("GET")

	// List of products - the main GET endpoint
	r.Handle("/products", ProductsHandler).Methods("GET")

	// Route for updating products - the main POST endpoint
	r.Handle("/products/{slug}/feedback", AddFeedbackHandler).Methods("POST")
}

// Create NotImplemented Function
func _NotImplmentedFunction(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
}

// Create Status Function
func _StatusFunction(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running"))
}

// Create Products Handler Function
func _ProductsFunction(w http.ResponseWriter, r *http.Request) {
	payload, _ := json.Marshal(products)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
}

// Create Feedback Function
func _AddFeedbackFunction(w http.ResponseWriter, r *http.Request) {
	var product Product
	vars := mux.Vars(r)
	slug := vars["slug"]

	for _, p := range products {
		if p.Slug == slug {
			product = p
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if product.Slug != "" {
		payload, _ := json.Marshal(product)
		w.Write([]byte(payload))
	} else {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

// Wrappers
var NotImplemented = http.HandlerFunc(_NotImplmentedFunction)
var StatusHandler = http.HandlerFunc(_StatusFunction)
var ProductsHandler = http.HandlerFunc(_ProductsFunction)
var AddFeedbackHandler = http.HandlerFunc(_AddFeedbackFunction)
