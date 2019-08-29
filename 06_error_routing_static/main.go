package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var gae *gaeInfo

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI != "/" {
		return404(w, r)
		return
	}
	fmt.Fprint(w, "Working Golang App\n")
}

// Display the 404 file
func return404(w http.ResponseWriter, _ *http.Request) {
	content, _ := ioutil.ReadFile("public/404.html")
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "%s", content)
}

// First We Intercept the WriteHeader Request for 404

// intercept404Writer would be extend / wrap the
//  http.ResponseWriter interface
type intercept404Writer struct {
	req             *http.Request       // Store the Current Request
	writer          http.ResponseWriter // Store the Original Response Writer
	notFoundHandler http.Handler        // Initiate Not Found Response and Ignore writes
	header          http.Header         // Used as the Virtual Header Storage
}

// Header would return the current 'Response Header' stored in the
// wrapped 'i.writer' member.
// In case we have set our custom header in the
// 'i.header' member then it would directly return that
// instead.
func (i *intercept404Writer) Header() http.Header {
	// If we have Processed Header Writing and received Final Call
	//  to Get the Header. Since all of the earlier calls
	//  had updated the internal member 'i.header' we do
	//  a copy operation in the wrapped WriteHeader function
	if i.header == nil && i.writer != nil {
		return i.writer.Header()
	}
	// Return the Virtual Header Storage for the Header till the full write
	//  is completed
	return i.header
}

func (i *intercept404Writer) WriteHeader(statusCode int) {
	// First copy all the material from Virtual Header Storage to Real Header
	for key, valuearr := range i.header {
		for _, val := range valuearr {
			i.writer.Header().Add(key, val)
		}
	}
	// We have Detected Found the 404
	if statusCode == http.StatusNotFound {
		// Start the NotFound Handler or the Not Found Page
		i.notFoundHandler.ServeHTTP(i.writer, i.req)
		// Remove the The Actual Writer as 404 has been served
		i.writer = nil
	} else {
		// Anything other than 404
		// Finally We write the Actual Status Code
		i.writer.WriteHeader(statusCode)
	}
	// Indicate we are completed with Filling up all headers
	i.header = nil
}

// Write here would bypass all calls in case the 'i.found404' flag
//  is set. Otherwise it would function like the normal write.
func (i *intercept404Writer) Write(b []byte) (int, error) {
	// In Normal Case
	if i.writer != nil {
		return i.writer.Write(b)
	}
	// Else Return a dummy output
	return len(b), nil
}

// Wrap the Handler
func static404Handler(h, h404Handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wr := &intercept404Writer{
			req:             r,
			writer:          w,
			notFoundHandler: h404Handler,
			header:          make(http.Header),
		}
		h.ServeHTTP(wr, r)
	})
}

// handlerFor404 Wrap Handler type for 404 displaying function
type handlerFor404 struct{}

// ServeHTTP to satisfy the http.Handler interface
func (h *handlerFor404) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Call the Custom 404 display Function
	return404(w, r)
}

func main() {

	http.HandleFunc("/", indexHandler)
	// Create the File Server
	hfileServer := http.FileServer(http.Dir("public"))
	// Create the 404 Page Handler
	h404 := &handlerFor404{}
	// Finally the Wrapping of File Server and 404 Handler
	hfileServerWith404 := static404Handler(hfileServer, h404)
	// http.Handle("/home/", http.StripPrefix("/home/", hfileServer))
	http.Handle("/home/", http.StripPrefix("/home/", hfileServerWith404))
	// Load the Environment Info from Appengine
	gae = getAppengineEnv()

	log.Printf("Starting server on :%s", gae.PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", gae.PORT), nil))
}
