package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCat_Existing(t *testing.T) {
	// 1. PRÉPARER : une requête pour un chat qui existe au démarrage (Toto)
	req := httptest.NewRequest("GET", "/cats/id1", nil)
	req.SetPathValue("catId", "id1")

	// 2. APPELER la fonction testée
	status, _ := getCat(req)

	// 3. VÉRIFIER le code de statut
	if status != http.StatusOK {
		t.Errorf("attendu %d, obtenu %d", http.StatusOK, status)
	}
}

func TestGetCat_AddedCatIsFound(t *testing.T) {
	// 1. PRÉPARER : un chat ajouté pour ce test, et son nettoyage
	catsDatabase["test-get"] = Cat{Name: "Romeo"}
	t.Cleanup(func() {
		delete(catsDatabase, "test-get")
	})

	req := httptest.NewRequest("GET", "/cats/test-get", nil)
	req.SetPathValue("catId", "test-get")

	// 2. APPELER la fonction testée
	status, _ := getCat(req)

	// 3. VÉRIFIER le code de statut
	if status != http.StatusOK {
		t.Errorf("attendu %d, obtenu %d", http.StatusOK, status)
	}
}

func TestGetCat_NotFound(t *testing.T) {
	// 1. PRÉPARER : une requête pour un chat qui n'existe pas
	req := httptest.NewRequest("GET", "/cats/chat-inexistant", nil)
	req.SetPathValue("catId", "chat-inexistant")

	// 2. APPELER la fonction testée
	status, _ := getCat(req)

	// 3. VÉRIFIER le code de statut
	if status != http.StatusNotFound {
		t.Errorf("attendu %d, obtenu %d", http.StatusNotFound, status)
	}
}

func TestDeleteCat_Existing(t *testing.T) {
	// 1. PRÉPARER : un chat ajouté pour ce test, et son nettoyage
	catsDatabase["test-delete"] = Cat{Name: "Felix"}
	t.Cleanup(func() {
		delete(catsDatabase, "test-delete")
	})

	req := httptest.NewRequest("DELETE", "/cats/test-delete", nil)
	req.SetPathValue("catId", "test-delete")

	// 2. APPELER la fonction testée
	status, _ := deleteCat(req)

	// 3.a VÉRIFIER le code de statut
	if status != http.StatusNoContent {
		t.Errorf("attendu %d, obtenu %d", http.StatusNoContent, status)
	}

	// 3.b VÉRIFIER que le chat a bien été supprimé de la base
	_, exist := catsDatabase["test-delete"]
	if exist {
		t.Errorf("le chat test-delete devrait avoir été supprimé, mais il est toujours dans la base")
	}
}

func TestDeleteCat_NotFound(t *testing.T) {
	// 1. PRÉPARER : une requête pour supprimer un chat qui n'existe pas
	req := httptest.NewRequest("DELETE", "/cats/chat-inexistant", nil)
	req.SetPathValue("catId", "chat-inexistant")

	// 2. APPELER la fonction testée
	status, _ := deleteCat(req)

	// 3. VÉRIFIER le code de statut
	if status != http.StatusNotFound {
		t.Errorf("attendu %d, obtenu %d", http.StatusNotFound, status)
	}
}