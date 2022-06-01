package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

type IamClient interface {
	ListAccountAliases(ctx context.Context,
		params *iam.ListAccountAliasesInput,
		optFns ...func(*iam.Options)) (*iam.ListAccountAliasesOutput, error)

	ListUsers(ctx context.Context,
		params *iam.ListUsersInput,
		optFns ...func(*iam.Options)) (*iam.ListUsersOutput, error)

	GetUser(ctx context.Context,
		params *iam.GetUserInput,
		optFns ...func(*iam.Options)) (*iam.GetUserOutput, error)
}
