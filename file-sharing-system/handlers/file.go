package handlers

import (
    "context"
    "net/http"
    "os"
    "file-sharing-system/utils"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {

    file, handler, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Error uploading file", http.StatusBadRequest)
        return
    }
    defer file.Close()


    localPath := "./uploads/" + handler.Filename
    f, err := os.Create(localPath)
    if err != nil {
        http.Error(w, "Error saving file", http.StatusInternalServerError)
        return
    }
    defer f.Close()
    _, err = f.ReadFrom(file)

 
    _, err = utils.Db.Exec(context.Background(),
        "INSERT INTO files (user_id, file_name, size, s3_url) VALUES ($1, $2, $3, $4)",
        1, handler.Filename, handler.Size, localPath)
    if err != nil {
        http.Error(w, "Error saving file metadata", http.StatusInternalServerError)
        return
    }

    w.Write([]byte("File uploaded successfully"))
}
