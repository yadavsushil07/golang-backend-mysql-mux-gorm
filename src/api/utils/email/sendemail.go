package email

// import (
// 	"fmt"
// 	"os"
// 	"testing"

// 	httptransport "github.com/go-openapi/runtime/client"
// )

// func getAPIClient(tb testing.TB) *client.sendinblue {
// 	tb.Helper()
// 	if testing.Short() {
// 		tb.Skip("short")
// 	}
// 	sib := client.NewHTTPClient(nil)
// 	sib.Transport.(*httptransport.Runtime).DefaultAuthentication = httptransport.APIKeyAuth("api-key", "header", getAPIKey(tb))
// 	fmt.Println(sib)
// 	return sib
// }

// const apiKeyEnvVariable = "API_KEY"

// func getAPIKey(tb testing.TB) string {
// 	tb.Helper()
// 	apiKey := os.Getenv(apiKeyEnvVariable)
// 	if apiKey == "" {
// 		tb.Skipf("environment variable %q is not defined", apiKeyEnvVariable)
// 	}
// 	return apiKey
// }
import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendEmail() {

	url := "https://api.sendinblue.com/v3/smtp/email"

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("api-key", "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")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
