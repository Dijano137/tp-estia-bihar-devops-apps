package main

import (
	"slices"
	"testing"
)

func TestListMapKeys(t *testing.T) {
	// 1. PRÉPARER : une map à nous, avec deux chats
	aMap := map[string]Cat{
		"id-a": {Name: "Felix"},
		"id-b": {Name: "Romeo"},
	}

	// 2. APPELER la fonction testée
	keys := listMapKeys(aMap)

	// 3.a VÉRIFIER la présence de chaque clé dans le résultat (quel que soit l'ordre)
	if !slices.Contains(keys, "id-a") {
		t.Errorf("la clé id-a est absente du résultat")
	}
	if !slices.Contains(keys, "id-b") {
		t.Errorf("la clé id-b est absente du résultat")
	}

	// 3.b VÉRIFIER le nombre de clés renvoyées
	if len(keys) != 2 {
		t.Errorf("attendu 2 clés, obtenu %d", len(keys))
	}
}