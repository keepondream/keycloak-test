package main

import (
	"context"
	"fmt"

	"github.com/Nerzal/gocloak/v13"
)

func main() {
	// ...
	client := gocloak.NewClient("https://....")

	ctx := context.Background()
	token, err := client.LoginAdmin(ctx, "admin", "password", "master")
	// token, err := client.Login(ctx, "account", "f89210b1-1431-4336-92fb-bf6c44ddafd7", "master", "testt", "testt")

	if err != nil {
		fmt.Println(err)
		panic("Something wrong with the credentials or url")
	}
	_ = token

	// userID, err := RegisterUser(ctx, client, token.AccessToken, "master", "password", gocloak.User{
	// 	Username:               gocloak.StringP("registerUser"),
	// 	Enabled:                gocloak.BoolP(true),
	// 	FirstName:              gocloak.StringP("shixiang"),
	// 	LastName:               gocloak.StringP("wang"),
	// 	Email:                  gocloak.StringP("qwer@123.com"),
	// 	ClientRoles:            &map[string][]string{},
	// 	RealmRoles:             &[]string{},
	// 	Groups:                 &[]string{},
	// 	ServiceAccountClientID: new(string),
	// 	Credentials:            &[]gocloak.CredentialRepresentation{},
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic("Something wrong with the credentials or url")
	// }
	// fmt.Println(userID)

	userID := "81a932f6-e37a-4189-beee-bbac3208bcfa"
	fmt.Println(GetUserInfo(ctx, client, token.AccessToken, "master", userID))

	// 客户端客户登录需要开启 Direct Access Grants Enabled
	token, err = client.Login(ctx, "account", "f89210b1-1431-4336-92fb-bf6c44ddafd7", "master", "registeruser", "password")
	if err != nil {
		fmt.Println(err)
		panic("Something wrong with the credentials or url")
	}
	fmt.Println("=====================================")
	fmt.Println(token.AccessToken)
	fmt.Println("=====================================")

	token.AccessToken = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJDWmVLNUpxNThHUnhfd1Baa0swVElnc0hITGwxTUQzVDAxeGh0X3ZQWkVjIn0.eyJleHAiOjE2ODA1MTkxNDgsImlhdCI6MTY4MDUxOTA4OCwianRpIjoiZmU3ZGE0Y2MtNjc2MS00ZTdkLThlNGEtOWY5NTcyOTE3Njc3IiwiaXNzIjoiaHR0cHM6Ly9rZXljbG9hay1rZWVwb25kcmVhbS5jbG91ZC5va3RldG8ubmV0L2F1dGgvcmVhbG1zL21hc3RlciIsInN1YiI6IjgxYTkzMmY2LWUzN2EtNDE4OS1iZWVlLWJiYWMzMjA4YmNmYSIsInR5cCI6IkJlYXJlciIsImF6cCI6ImFjY291bnQiLCJzZXNzaW9uX3N0YXRlIjoiZWQ3NTUyNzQtMGFkNi00NTMyLTg3ZWUtNjcxZWNiYzM5NmEyIiwiYWNyIjoiMSIsInJlYWxtX2FjY2VzcyI6eyJyb2xlcyI6WyJvZmZsaW5lX2FjY2VzcyIsInVtYV9hdXRob3JpemF0aW9uIl19LCJyZXNvdXJjZV9hY2Nlc3MiOnsiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJvcGVuaWQgZW1haWwgcHJvZmlsZSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwibmFtZSI6InNoaXhpYW5nIHdhbmciLCJwcmVmZXJyZWRfdXNlcm5hbWUiOiJyZWdpc3RlcnVzZXIiLCJnaXZlbl9uYW1lIjoic2hpeGlhbmciLCJmYW1pbHlfbmFtZSI6IndhbmciLCJlbWFpbCI6InF3ZXJAMTIzLmNvbSJ9.R_W1IXulNEJ4o1fnC2TFRPCIZyvmdKztA0fe5laaor5V3qhH72lGwpOMSwBYu0DAlNgSx4VxQMqHtcuUIyy5kPLE9ZpQzOnN7_1cdL5CU4j86n9e7Q5koXtATqp2bpRgt6KHxJ2cAvDexkLcXaWtai9y1Cn6j65u8l691F9WcAiWTFelhK1hYjYTueCFeL5rJ21x4s04E4mG1DSBwl46oIR5QSt3W2XqVg1wHeJSnZvNbHXT3P1pf_RkwM54hGxJ0gxBdcnKX7E9oudkoG6dETP0oaMfkfEQOlQvX1acO2_PHfC4JAABurU6LmWLDqQzxnXJOlrFZc31x6yUo0oz3g"
	// 验证客户身份
	fmt.Println(client.RetrospectToken(ctx, token.AccessToken, "account", "f89210b1-1431-4336-92fb-bf6c44ddafd7", "master"))

}

// invoice服务用户注册
func RegisterUser(ctx context.Context, adminClient *gocloak.GoCloak, token, realm, password string, user gocloak.User) (string, error) {
	userId, err := adminClient.CreateUser(ctx, token, realm, user)
	if err != nil {
		return "", err
	}
	err = adminClient.SetPassword(ctx, token, userId, realm, password, false)
	if err != nil {
		return "", err
	}
	return userId, nil
}

// 根据userId获取用户信息
func GetUserInfoByID(ctx context.Context, adminClient *gocloak.GoCloak, token, realm, userId string) (*gocloak.User, error) {
	return adminClient.GetUserByID(ctx, token, realm, userId)
}
