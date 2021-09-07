package auth

import ("")

func checkToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	  header := r.Header
	  
	  if _, ok := header["Token"]; ok{
		next(w, r)
		return
	  }
	  respondJSON(w, 401, "Unauthorized.")
	}
}