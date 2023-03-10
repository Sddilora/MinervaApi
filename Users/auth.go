package auth

import (
	//auth "github.com/ScafTeam/firebase-go-client/auth"

	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

func SigninHandler(c *fiber.Ctx) error {

	//_, _, appfire := create.NewFireStore()

	var UserSignin struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&UserSignin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	k := authenticationHandler(UserSignin.Email, UserSignin.Password)
	log.Print("bakalım", k)

	resp := fiber.Map{
		"message": "Login access permitted ",
	}
	return c.Status(fiber.StatusCreated).JSON(resp)

}

func authenticationHandler(email string, password string) map[string]interface{} {
	// Create Http request
	apiKey, err := os.ReadFile("Users/apikey.txt")
	if err != nil {
		panic(err)
	}
	url := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=%s", string(apiKey))
	requestData := map[string]string{
		"email":             email,
		"password":          password,
		"returnSecureToken": "true",
	}
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// HTTP isteğini gönder ve yanıtı işle
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Yanıtı oku
	var responseBody map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		panic(err)
	}

	return responseBody
}

// client, err := appfire.Auth(context.Background())
// if err != nil {
// 	log.Printf("error while signin : %v\n", err)
// }

// user := &auth.UserToImport{
// 	params: UserSignin.Email,
// }

//	user.PasswordHash([]byte(UserSignin.Password))

//u, err := client.GetUserByEmail(context.Background(), UserSignin.Email)
//client.ImportUsers([]*auth.UserToImport{user})
// u, err := client.GetUserByEmail(context.Background(), UserSignin.Email)
// if err != nil {
// 	log.Printf("error getting user by email: %s: %v\n", UserSignin.Email, err)
// }

// k, err := client.CustomToken(context.Background(), u.UID)
// if err != nil {
// 	log.Printf("error getting custom token: %s: %v\n", UserSignin.Email, err)
// }

// z, err := client.SAMLProviderConfig()(context.Background(), u.ProviderID)
// if err != nil {
// 	log.Printf("error getting OpenID Connect: %s: %v\n", UserSignin.Email, err)
// }

// err = client.Sign(ctx, email, password)
// if err != nil {
// 	log.Fatalf("error verifying password: %v\n", err)
// }
// log.Printf("Successfully fetched user data: %v\n", z, u.UID, u.ProviderID)

// func loginHandler(w http.ResponseWriter, r *http.Request) {

// 	decoder := json.NewDecoder(r.Body)
// 	var user User
// 	err := decoder.Decode(&user)
// 	if err != nil {
// 		log.Fatal("Request body cant decoded", err)
// 	}
// 	authParams := (&auth.UserToSignInWithEmailAndPassword{}).
// 		Email(user.Email).
// 		Password(user.Password)
// 	_, err = Client.EmailSignInLink(context.Background(), authParams)
// 	if err != nil {
// 		// hata işleme
// 	}

// 	// Başarılı bir yanıt gönderin
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Kullanıcı başarıyla doğrulandı"))
// }

// var _, _, ctx, err = create.NewFireStore()

// func Auth(App *fiber.App, appfire *firestore.Client) {

// 	App.Post("/login", func(c *fiber.Ctx) error {

// 		var login struct { // Parse request body
// 			Name     string `json:"name"`
// 			Email    string `json:"email"`
// 			Password string `json:"password"`
// 		}

// 		if err := c.BodyParser(&login); err != nil {
// 			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 				"message": "Invalid request body",
// 			})
// 		}

// 		documentPath := fmt.Sprint("Users/", login.Email)

// 		x, err := appfire.Doc(documentPath).Get(ctx)

// 		emailCr := x.Data()

// 		if err != nil {
// 			return err
// 		}

// 		incomePassword := login.Password
// 		truePassword := emailCr["password"]

// 		if incomePassword == truePassword {
// 			return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 				"message": "Acces permitted",
// 			})
// 		} else {
// 			return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 				"message": "Acces denied",
// 			})
// 		}

// 	})
// }
