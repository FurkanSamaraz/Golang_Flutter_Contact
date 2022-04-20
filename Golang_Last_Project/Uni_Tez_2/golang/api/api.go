package api

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"main.go/beeeb"
	"main.go/block"
	"main.go/checkError"
	"main.go/login"
)

type Data struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var db *sql.DB

func GetJson(w http.ResponseWriter, r *http.Request) {
	db, _ = sql.Open("sqlite3", "veri.db")

	vt, err := db.Query("SELECT * FROM data")
	checkError.ErrorContr(err, "10")
	var dataUs Data
	var peopleData []Data

	for vt.Next() {

		vt.Scan(&dataUs.Id, &dataUs.Name)

		peopleData = append(peopleData, dataUs)

	}
	var dataStr string

	jsonResp, _ := json.MarshalIndent(peopleData, "", "\n")
	dataStr = string(jsonResp)

	bytes := make([]byte, 32) //AES-256 için rastgele bir 32 bayt anahtar oluşturun.
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}
	key := hex.EncodeToString(bytes) //anahtarı bayt cinsinden kodlayın ve gizli olarak saklayın, bir kasaya koyun
	//	fmt.Printf("KEY(anahtar) => %s\n\n", key)

	encrypted := block.Encrypt(dataStr, key)
	//	fmt.Printf("ENCRYPTED(şifreli) => %s\n\n", encrypted)

	if login.Control == true {

		beeeb.BebMessage("😊", "Data Decrypted✅")

		decrypted := block.Decrypt(encrypted, key)
		//fmt.Printf("decrypted(şifre çözüm) => %s\n", "https://localhost:8080/protected")

		mySlice := []byte(decrypted)

		defer db.Close()
		defer vt.Close()

		w.Write(mySlice)
	}

}
