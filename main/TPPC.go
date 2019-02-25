package main

import (
	"fmt"
	"log"
	"net/http"
)

type client struct {
	nomClient     string
	adresseClient string
	numContrat    string
	numClient     string
}

type intermediaire struct {
	nomAgent     string
	adresseAgent string
	telAgent     string
	faxAgent     string
}

type garantie struct {
	animauxVivants              string
	frigo                       string
	vehiculesEtRemorques        string
	citerne                     string
	typeContrat                 string
	activite                    string
	flotte                      string
	typeCotisation              string
	typeMarchandise             string
	expositions                 string
	nombreExpositions           string
	IAC                         string
	HIAC                        string
	garantieVol                 string
	capital                     string
	nombreDeSinistres           string
	franchiseToutCas            string
	franchiseParVehiculeRoulant string
	cotisationParVehicule       string
	cotisationParChiffreAffaire string
	dateFin                     string
}

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":

		http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
		http.ServeFile(w, r, "TPPC.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)

		nomAgent := r.FormValue("nomAgent")
		adresseAgent := r.FormValue("adresseAgent")
		telAgent := r.FormValue("telAgent")
		faxAgent := r.FormValue("faxAgent")

		agent := intermediaire{nomAgent: nomAgent, adresseAgent: adresseAgent, telAgent: telAgent, faxAgent: faxAgent}
		fmt.Fprintf(w, "Agent = ", agent)

		nomClient := r.FormValue("nomClient")
		adresseClient := r.FormValue("adresseClient")
		numContrat := r.FormValue("numeroContrat")
		//numClient := ficheClientServlet.numClient

		cl := client{nomClient: nomClient, adresseClient: adresseClient, numContrat: numContrat, numClient: "123"}
		fmt.Fprintf(w, "Client = ", cl)

		animauxVivants := r.FormValue("animauxVivants")
		frigo := r.FormValue("frigo")
		vehiculesEtRemorques := r.FormValue("vehiculesEtRemorques")
		citerne := r.FormValue("citerne")
		typeContrat := r.FormValue("typeContrat")
		activite := r.FormValue("activite")
		flotte := r.FormValue("flotte")
		typeCotisation := r.FormValue("typeCotisation")
		typeMarchandise := r.FormValue("typeMarchandise")
		expositions := r.FormValue("expositions")
		nombreExpositions := r.FormValue("nombreExpositions")
		IAC := r.FormValue("IAC")
		HIAC := r.FormValue("HIAC")
		garantieVol := r.FormValue("garantieVol")
		capital := r.FormValue("capital")
		nombreDeSinistres := r.FormValue("nombreDeSinistres")
		franchiseToutCas := r.FormValue("franchiseToutCas")
		franchiseParVehiculeRoulant := r.FormValue("franchiseParVehiculeRoulant")
		cotisationParVehicule := r.FormValue("cotisationVehicule")
		cotisationParChiffreAffaire := r.FormValue("cotisationCA")
		dateFin := r.FormValue("dateFin")
		ga := garantie{animauxVivants: animauxVivants, frigo: frigo, vehiculesEtRemorques: vehiculesEtRemorques, citerne: citerne, typeContrat: typeContrat, activite: activite, flotte: flotte, typeCotisation: typeCotisation, typeMarchandise: typeMarchandise, expositions: expositions, nombreExpositions: nombreExpositions, IAC: IAC, HIAC: HIAC, garantieVol: garantieVol, capital: capital, nombreDeSinistres: nombreDeSinistres, franchiseToutCas: franchiseToutCas, franchiseParVehiculeRoulant: franchiseParVehiculeRoulant, cotisationParVehicule: cotisationParVehicule, cotisationParChiffreAffaire: cotisationParChiffreAffaire, dateFin: dateFin}
		fmt.Fprintf(w, "Garantie = ", ga)

		projetTPPC(ga, cl, agent)

		/*dateEffet  := ficheClientServlet.dateEffet;
		if(typeContrat.equals("contratTemporaire"))
		{
			  dateFin   := r.FormValue("dateFin");




		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Address = %s\n", address) */
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	http.HandleFunc("/", hello)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
