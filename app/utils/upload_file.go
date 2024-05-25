package utils

import (
    "encoding/base64"
    "io/ioutil"
    "os"
    "path/filepath"
)

func UploadFile(base64String, uploadDirectory, newFileName string) error {
    // Decode the base64 string into binary data
    decoded, err := base64.StdEncoding.DecodeString(base64String)
    if err != nil {
        return err
    }

    // Create the upload directory if it doesn't exist
    if err := os.MkdirAll(uploadDirectory, 0755); err != nil {
        return err
    }

    // Write the binary data to a file
    filePath := filepath.Join(uploadDirectory, newFileName)
    if err := ioutil.WriteFile(filePath, decoded, 0644); err != nil {
        return err
    }

    return nil
}
