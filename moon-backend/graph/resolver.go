package graph

import (
	"moon/ent"
	"moon/pkg/s3"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

type Resolver struct {
	Client   *ent.Client
	S3Client *s3.Client
}
