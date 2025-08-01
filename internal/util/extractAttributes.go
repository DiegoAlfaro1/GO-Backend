package util

import (
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func ExtractAttributes(output *cognito.AdminGetUserOutput) map[string]string {
	attrs := make(map[string]string)
	for _, attr := range output.UserAttributes {
		attrs[*attr.Name] = *attr.Value
	}
	return attrs
}
