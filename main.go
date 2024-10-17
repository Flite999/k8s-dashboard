package main

import (
	"html/template"
	"log"
	"net/http"
)

// Handler to render the dashboard
func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		ClusterName string
		Metrics     []PodMetrics
	}{
		ClusterName: "My Cluster",
		Metrics:     getClusterMetrics(), // Call the function from metrics.go
	}

	tmpl := template.Must(template.ParseFiles("templates/dashboard.html"))
	tmpl.Execute(w, data)
}

func redirectToDashboard(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/dashboard", http.StatusFound)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/dashboard", dashboardHandler)

	// Redirect root to /dashboard
	http.HandleFunc("/", redirectToDashboard)

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
