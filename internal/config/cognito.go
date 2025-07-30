package config

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/google/uuid"
)

type User struct {
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	Birthdate string `json:"birthdate" binding:"required"` // ✅ Added birthdate
}

type UserConfirmation struct {
	Email string `json:"email" binding:"required,email"`
	Code  string `json:"code" binding:"required"`
}

type UserLogin struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CognitoInterface interface {
	SignUp(user *User) error
	ConfirmAccount(user *UserConfirmation) error
	SignIn(user *UserLogin) (string, error)
}

type cognitoClient struct {
	coginitoClient *cognito.CognitoIdentityProvider
	appClientID    string
}

func NewCognitoClient(appClientID string) CognitoInterface {
	region := os.Getenv("AWS_REGION")
	config := &aws.Config{Region: aws.String(region)}
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	client := cognito.New(sess)

	return &cognitoClient{
		coginitoClient: client,
		appClientID:    appClientID,
	}
}

func (c *cognitoClient) SignUp(user *User) error {
	userCognito := &cognito.SignUpInput{
		ClientId: aws.String(c.appClientID),
		Username: aws.String(user.Email),
		Password: aws.String(user.Password),
		UserAttributes: []*cognito.AttributeType{
			{
				Name:  aws.String("name"),
				Value: aws.String(user.Name),
			},
			{
				Name:  aws.String("email"),
				Value: aws.String(user.Email),
			},
			{
				Name:  aws.String("birthdate"),
				Value: aws.String(user.Birthdate),
			},
			{
				Name:  aws.String("custom:custom_id"),
				Value: aws.String(uuid.NewString()),
			},
		},
	}

	_, err := c.coginitoClient.SignUp(userCognito)
	return err
}

func (c *cognitoClient) ConfirmAccount(user *UserConfirmation) error {
	confirmationInput := &cognito.ConfirmSignUpInput{
		Username:         aws.String(user.Email),
		ConfirmationCode: aws.String(user.Code),
		ClientId:         aws.String(c.appClientID),
	}
	_, err := c.coginitoClient.ConfirmSignUp(confirmationInput)
	return err
}

func (c *cognitoClient) SignIn(user *UserLogin) (string, error) {
	authInput := &cognito.InitiateAuthInput{
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		AuthParameters: aws.StringMap(map[string]string{
			"USERNAME": user.Email,
			"PASSWORD": user.Password,
		}),
		ClientId: aws.String(c.appClientID),
	}

	result, err := c.coginitoClient.InitiateAuth(authInput)
	if err != nil {
		return "", err
	}

	return *result.AuthenticationResult.AccessToken, nil
}
