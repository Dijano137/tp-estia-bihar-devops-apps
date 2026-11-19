package main

import "net/http"

func getCat(req *http.Request) (int, any) {

	catId := req.PathValue("catId")
	cat, exist := catsDatabase[catId]

	if !exist {
		Logger.Info("Cat not found")
		return http.StatusNotFound, "Cat not found"
	}
	return http.StatusOK, cat
}

func deleteCat(req *http.Request) (int, any) {

	catId := req.PathValue("catId")
	_, exist := catsDatabase[catId]

	if exist {
		delete(catsDatabase, catId)
		return http.StatusNoContent, nil
	} else {
		Logger.Info("Cat not found")
		return http.StatusNotFound, "Cat not found"
	}
}
