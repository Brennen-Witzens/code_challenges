package main

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
)

func main() {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "test-project")
	if err != nil {
	}

	defer client.Close()

	nm := client.Doc("States")
	err = client.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		// WithReadOptions returns a transaction
		/*
		* So we start our transaction with RunTransaction.
		* We set it to Read Only (using firestore.ReadOnly - which I think is true)
		* We then set readOptions - IE the time we want to read the documents from
		* Then go forward with the transaction as before BUT just using the readOptions as our "starter" -> readOptions.Get(), etc
		 */
		readOptions := tx.WithReadOptions(firestore.ReadOption(firestore.ReadTime(time.Now())))
		doc, err := readOptions.Get(nm)
		if err != nil {
			return err
		}
		_, err = doc.DataAt("pop")
		if err != nil {
			return err
		}
		return nil
	}, firestore.ReadOnly)
	if err != nil {
	}

}
