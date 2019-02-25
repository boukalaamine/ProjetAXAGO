package main

import (
	"fmt"
	"log"
	"reflect"

	"baliance.com/gooxml/document"
)

func projetTPPC(gar garantie, cl client, inter intermediaire) {
	//fmt.Println("Salut")
	doc, err := document.Open("form.docx")
	if err != nil {
		log.Fatalf("error opening form: %s", err)
	}

	// FindAllFields is a helper function that traverses the document
	// identifying fields
	fields := doc.FormFields()
	fmt.Println("found", len(fields), "fields")

	for _, fld := range fields {

		fmt.Println("- Name:", fld.Name(), "Type:", fld.Type(), "Value:", fld.Value())

		//Informations du client
		fmt.Println(reflect.TypeOf(fld))

		if fld.Name() == "nomClient" {
			fld.SetValue(cl.nomClient)
		}

		if fld.Name() == "adresseClient" {
			fld.SetValue(cl.adresseClient)
		}

		if fld.Name() == "numContrat" {
			fld.SetValue(cl.numContrat)
		}

		if fld.Name() == "numClient" {
			fld.SetValue(cl.numClient)
		}

		//Informations du client

		if fld.Name() == "nomAgent" {
			fld.SetValue(inter.nomAgent)
		}

		if fld.Name() == "adresseAgent" {
			fld.SetValue(inter.adresseAgent)
		}

		if fld.Name() == "telAgent" {
			fld.SetValue(inter.telAgent)
		}

		if fld.Name() == "faxAgent" {
			fld.SetValue(inter.faxAgent)
		}

		if fld.Name() == "activite" {
			fld.SetValue(gar.activite)
		}

		if fld.Name() == "sensibilite1" {
			clauseMarchandiseSensibleVol(gar, fld)
		}

		if fld.Name() == "sensibilite2" {
			clauseMarchandiseSensibleVol(gar, fld)
		}

		if fld.Name() == "typeContrat" {
			typeContrat(gar, fld)
		}

		if fld.Name() == "titreFrigo" {
			titreFrigo(gar, fld)
		}

		if fld.Name() == "frigo" {
			frigo(gar, fld)
		}

		if fld.Name() == "citerne" {
			citerne(gar, fld)
		}

		if fld.Name() == "titreCiterne" {
			titreCiterne(gar, fld)
		}

		if fld.Name() == "animaux" {
			animaux(gar, fld)
		}

		if fld.Name() == "titreAnimaux" {
			titreAnimaux(gar, fld)
		}

		if fld.Name() == "titreExpositions" {
			titreExpositions(gar, fld)
		}

		if fld.Name() == "expositions" {
			expositions(gar, fld)
		}

		if fld.Name() == "titreVehiculesEtRemo" {
			titreVehiculesEtRemorques(gar, fld)
		}

		if fld.Name() == "vehiculesEtRemorques" {
			vehiculesEtRemorques(gar, fld)
		}

		//Informations sur les garanties
	}
	doc.SaveToFile("filled-form.docx")

}

func clauseMarchandiseSensibleVol(gar garantie, fld document.FormField) {
	sensible := ""
	if gar.typeMarchandise == "marchandiseNonSensible" {
		sensible = "non"
	}
	fld.SetValue(sensible)

}

func typeContrat(gar garantie, fld document.FormField) {
	if gar.typeContrat == "contratTemporaire" {
		data := "Le contrat est souscrit pour une durée temporaire et cessera tous ses effets à compter du : " + gar.dateFin
		fld.SetValue(data)
	}
	if gar.typeContrat == "contratAnnuel" {
		data := "Le contrat est souscrit jusqu'à la date de la première échéance principale et est dès lors reconduit tacitement d'année en année, sauf résiliation par l'une ou l'autre des parties dans les cas, formes et délais prévus aux Conditions générales."
		fld.SetValue(data)
	}
}

func frigo(gar garantie, fld document.FormField) {
	if gar.frigo == "on" {
		fld.SetValue("Cette garantie est acquise aux conditions exprimées à  l'article 3.4.1 « Autres garanties» des Conditions Générales.")

	}
}

func titreFrigo(gar garantie, fld document.FormField) {
	if gar.frigo == "on" {

		fld.SetValue("Transport Frigorifique")
	}
}

func citerne(gar garantie, fld document.FormField) {
	if gar.citerne == "on" {

		fld.SetValue("Cette garantie est acquise aux conditions exprimées à  l'article 3.4.3 « Autres garanties» des Conditions Générales.")
	}
}

func titreCiterne(gar garantie, fld document.FormField) {
	if gar.citerne == "on" {

		fld.SetValue("Transport en Citerne")

	}
}

func animaux(gar garantie, fld document.FormField) {
	if gar.animauxVivants == "on" {
		fld.SetValue("Cette garantie est acquise aux conditions exprimées à  l'article 3.4.1 « Autres garanties» des Conditions Générales.")
	}
}

func titreAnimaux(gar garantie, fld document.FormField) {
	if gar.animauxVivants == "on" {
		fld.SetValue("Transport d'animaux vivants")
	}
}

func expositions(gar garantie, fld document.FormField) {
	if gar.expositions == "on" {
		data := "Pour les marchandises transportées au moyen des véhicules agréés au contrat, la garantie est acquise conformément aux articles 1, 2.3, 2.7, 2.8, 5 et 6 des Conditions Générales Multirisque Exposition dans la limite de " + gar.nombreExpositions + " expositions par année d'assurance.\r\n" +
			"\r\n" +
			"La garantie est étendue aux opérations de déplacement des marchandises entre le véhicule et le stand d'exposition, effectuées avec des moyens techniques appropriés.  \r\n" +
			"\r\n" +
			"L'Assuré déclare que pendant les heures d'ouverture, la surveillance est permanente et qu'en dehors des heures d'ouverture, les locaux sont gardiennés.\r\n" +
			"\r\n" +
			"La durée de la garantie par foire, salon et/ou exposition ne pourra excéder quinze jours.\r\n" +
			"\r\n" +
			""
		fld.SetValue(data)
	}
}

func titreExpositions(gar garantie, fld document.FormField) {
	if gar.expositions == "on" {
		fld.SetValue("Garantie expositions")
	}
}

func vehiculesEtRemorques(gar garantie, fld document.FormField) {
	if gar.vehiculesEtRemorques == "on" {
		data := "Vous nous avez déclaré lors de la souscription l'état de votre parc automobile et ses caractéristiques conformes à  celles figurant sur les cartes grises de chaque véhicule.\r\n" +
			"\r\n" +
			"En fonction de cet état un capital garanti a été fixé d'un commun accord au vu des marchandises que vous transportez.\r\n" +
			"\r\n" +
			"Nous vous dispensons de déclarer en cours d'exercice toute mise en service de nouveau véhicule ou retrait étant entendu :\r\n" +
			"\r\n" +
			"Que notre garantie vous est acquise au cours de l'exercice concerné,\r\n" +
			"\r\n" +
			"Que vous vous engagez à  déclarer l'état détaillé de votre parc arràªté à  la date d'échéance anniversaire de votre contrat.\r\n" +
			"\r\n" +
			"à € la suite de cette déclaration, nous établirons un avenant sur l'exercice antérieur dont le montant de la régularisation [cotisation ou ristourne] sera calculé à  raison de 50 % de la différence entre l'ancienne et la nouvelle cotisation annuelle.\r\n" +
			"\r\n" +
			"à € défaut de cette déclaration, nous pourrons vous mettre en demeure, par lettre recommandée, de satisfaire à  cette obligation dans les 10 jours à  partir de l'envoi de cette lettre, le cachet de la Poste faisant foi.\r\n" +
			"Si passé ce délai, vous ne nous avez pas adressé l'état de votre parc, nous émettrons un avenant ressortant une cotisation provisoire calculée sur la base de votre dernière déclaration majorée de 50 %.\r\n" +
			"\r\n" +
			"Nous pouvons également, à  toute époque, vous demander la production de toutes pièces justificatives afin de procéder à  une vérification de vos déclarations.\r\n" +
			""
		fld.SetValue(data)
	}
}

func titreVehiculesEtRemorques(gar garantie, fld document.FormField) {
	if gar.vehiculesEtRemorques == "on" {
		fld.SetValue("Véhicules et remorques")
	}
}
