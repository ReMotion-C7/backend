package config

import (
	storage_go "github.com/supabase-community/storage-go"
)

var StorageClient *storage_go.Client

func SetUpStorage() {

	projectEndpoint := LoadEnvConfig("PROJECT_ENDPOINT")
	apiKey := LoadEnvConfig("API_KEY")

	storageClient := storage_go.NewClient(projectEndpoint, apiKey, nil)

	StorageClient = storageClient
}

func GetStorageClient() *storage_go.Client {

	return StorageClient

}
