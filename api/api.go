package api

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"cave/api/db"
	"cave/api/handler"
	"cave/config"

	"github.com/fasthttp/router"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"golang.org/x/net/context"
)

// Api has the mongo database and router instances
type Api struct {
	Router *mux.Router
	DB     *mongo.Database
}


const (
	AcceptJson  = "application/json"
	AcceptRest  = "application/vnd.pgrst.object+json"
	ContentText = "text/plain; charset=utf8"
	ContentRest = "application/vnd.pgrst.object+json; charset=utf-8"
	ContentJson = "application/json; charset=utf-8"
)

type webServer struct {
	Config config.Configuration
	Addr   string
	ln     net.Listener
	router *router.Router
	debug  bool
}

// ConfigAndRunApp will create and initialize Api structure. Api factory function.
func ConfigAndRunApp(config *config.Configuration) {
	api := new(Api)
	api.Initialize(config)
	api.Run(config.Address)
}

// Initialize initialize the api with
func (api *Api) Initialize(config *config.Configuration) {
	api.DB = db.InitialConnection(config.DBName, config.MongoURL)
	api.createIndexes()

	api.Router = mux.NewRouter()
	api.UseMiddleware(handler.JSONContentTypeMiddleware)
	api.setRouters()
}

// SetupRouters will register routes in router
func (api *Api) setRouters() {
	api.Post("/person", api.handleRequest(handler.CreatePerson))
	api.Patch("/person/{id}", api.handleRequest(handler.UpdatePerson))
	api.Put("/person/{id}", api.handleRequest(handler.UpdatePerson))
	api.Get("/person/{id}", api.handleRequest(handler.GetPerson))
	api.Get("/person", api.handleRequest(handler.GetPersons))
	api.Get("/person", api.handleRequest(handler.GetPersons), "page", "{page}")
}

// UseMiddleware will add global middleware in router
func (api *Api) UseMiddleware(middleware mux.MiddlewareFunc) {
	api.Router.Use(middleware)
}

// createIndexes will create unique and index fields.
func (api *Api) createIndexes() {
	// username and email will be unique.
	keys := bsonx.Doc{
		{Key: "email", Value: bsonx.Int32(1)},
	}
	people := api.DB.Collection("people")
	db.SetIndexes(people, keys)
}

// Get will register Get method for an endpoint
func (api *Api) Get(path string, endpoint http.HandlerFunc, queries ...string) {
	api.Router.HandleFunc(path, endpoint).Methods("GET").Queries(queries...)
}

// Post will register Post method for an endpoint
func (api *Api) Post(path string, endpoint http.HandlerFunc, queries ...string) {
	api.Router.HandleFunc(path, endpoint).Methods("POST").Queries(queries...)
}

// Put will register Put method for an endpoint
func (api *Api) Put(path string, endpoint http.HandlerFunc, queries ...string) {
	api.Router.HandleFunc(path, endpoint).Methods("PUT").Queries(queries...)
}

// Patch will register Patch method for an endpoint
func (api *Api) Patch(path string, endpoint http.HandlerFunc, queries ...string) {
	api.Router.HandleFunc(path, endpoint).Methods("PATCH").Queries(queries...)
}

// Delete will register Delete method for an endpoint
func (api *Api) Delete(path string, endpoint http.HandlerFunc, queries ...string) {
	api.Router.HandleFunc(path, endpoint).Methods("DELETE").Queries(queries...)
}

// Run will start the http server on host that you pass in. host:<ip:port>
func (api *Api) Run(host string) {
	// use signals for shutdown server gracefully.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt, os.Kill)
	go func() {
		log.Fatal(http.ListenAndServe(host, api.Router))
	}()
	log.Printf("Server is listning on http://%s\n", host)
	sig := <-sigs
	log.Println("Signal: ", sig)

	log.Println("Stoping MongoDB Connection...")
	api.DB.Client().Disconnect(context.Background())
}

// RequestHandlerFunction is a custome type that help us to pass db arg to all endpoints
type RequestHandlerFunction func(db *mongo.Database, w http.ResponseWriter, r *http.Request)

// handleRequest is a middleware we create for pass in db connection to endpoints.
func (api *Api) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(api.DB, w, r)
	}
}
