package main

import (
	"context"
	"fmt"
	"log"
	"os"

	accessapproval "cloud.google.com/go/accessapproval/apiv1"
	pb "cloud.google.com/go/accessapproval/apiv1/accessapprovalpb"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {
	projectId := os.Getenv("GOOGLE_CLOUD_PROJECT")
	ctx := context.Background()
	req := &pb.ListApprovalRequestsMessage{
		Parent: "projects/" + projectId,
	}

	gcpClient, err := accessapproval.NewClient(ctx, option.WithTelemetryDisabled())
	if err != nil {
		log.Fatal(err)
	}

	it := gcpClient.ListApprovalRequests(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(resp.Name)
	}
}
