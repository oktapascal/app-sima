package database

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/oktapascal/app-sima/bootstraps"
	"github.com/oktapascal/app-sima/utils"
)

func NewFirestoreClient(config bootstraps.Config, ctx context.Context) *firestore.Client {
	var projectID = config.Get("FIRESTORE_PROJECT_ID")

	client, err := firestore.NewClient(ctx, projectID)
	utils.PanicIfError(err)

	return client
}
