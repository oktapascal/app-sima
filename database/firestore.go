package database

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/oktapascal/app-barayya/bootstraps"
	"github.com/oktapascal/app-barayya/utils"
)

func NewFirestoreClient(config bootstraps.Config, ctx context.Context) *firestore.Client {
	var projectID = config.Get("FIRESTORE_PROJECT_ID")

	client, err := firestore.NewClient(ctx, projectID)
	utils.PanicIfError(err)

	return client
}
