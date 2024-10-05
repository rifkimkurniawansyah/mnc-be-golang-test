package handlers

import (
    "encoding/json"
    "golang-api-test/models"
    "golang-api-test/utils"
    "net/http"
)

var (
    customers []models.Customer
    loggedInUsers = make(map[string]bool)
)

func LoadCustomers() {
    utils.ReadJsonFile("data/customers.json", &customers)
}

func Login(w http.ResponseWriter, r *http.Request) {
    var user models.Customer
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    for _, customer := range customers {
        if customer.Username == user.Username && customer.Password == user.Password {
            loggedInUsers[user.Username] = true
            utils.Log(user.Username + " logged in")
            w.WriteHeader(http.StatusOK)
            json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
            return
        }
    }

    http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func Logout(w http.ResponseWriter, r *http.Request) {
    var user models.Customer
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    if _, ok := loggedInUsers[user.Username]; ok {
        delete(loggedInUsers, user.Username)
        utils.Log(user.Username + " logged out")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]string{"message": "Logout successful"})
        return
    }

    http.Error(w, "Unauthorized", http.StatusUnauthorized)
}
