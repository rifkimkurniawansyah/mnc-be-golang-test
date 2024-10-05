package handlers

import (
    "encoding/json"
    "golang-api-test/models"
    "golang-api-test/utils"
    "net/http"
	"fmt"
)

func Payment(w http.ResponseWriter, r *http.Request) {
    var payment models.Payment
    if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    if !loggedInUsers[payment.Username] {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

	utils.Log(payment.Username + " made a payment of " + fmt.Sprintf("%.2f", payment.Amount))
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Payment successful"})
}
