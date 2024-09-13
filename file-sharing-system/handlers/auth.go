package handlers

import (
    "context"
    "encoding/json"
    "net/http"
    "file-sharing-system/models"
    "file-sharing-system/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)

    _, err := utils.Db.Exec(context.Background(),
        "INSERT INTO users (email, password) VALUES ($1, $2)", user.Email, user.Password)
    if err != nil {
        http.Error(w, "Error registering user", http.StatusInternalServerError)
        return
    }

    w.Write([]byte("User registered successfully"))
}

func Login(w http.ResponseWriter, r *http.Request) {
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)

    var dbPassword string
    err := utils.Db.QueryRow(context.Background(),
        "SELECT password FROM users WHERE email = $1", user.Email).Scan(&dbPassword)

    if err != nil || dbPassword != user.Password {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    token, _ := utils.GenerateJWT(user.Email)
    w.Write([]byte(token))
}
