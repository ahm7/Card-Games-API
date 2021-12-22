package ApiHelper



import(
	"net/http"
)


func NotFound(w http.ResponseWriter){

	w.WriteHeader(http.StatusNotFound)
}

func ContentNotFOund(w http.ResponseWriter){

	w.WriteHeader(http.StatusNoContent)
}

func Success(w http.ResponseWriter){

	w.WriteHeader(http.StatusAccepted)
}
func BadRequest(w http.ResponseWriter){

	w.WriteHeader(http.StatusBadRequest)
}