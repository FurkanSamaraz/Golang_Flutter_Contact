package protected

import (
	"net/http"

	"main.go/api"
	"main.go/beeeb"
)

func Protected(w http.ResponseWriter, r *http.Request) {
	api.GetJson(w, r)
	beeeb.BebMessage("😊", "Protected Successful✅\n 📲 Mobil App\n")

}
