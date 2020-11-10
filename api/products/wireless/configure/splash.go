package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
)

type SplashPageSettings struct {
	SsidNumber     int    `json:"ssidNumber"`
	SplashPage     string `json:"splashPage"`
	SplashURL      string `json:"splashUrl"`
	UseSplashURL   bool   `json:"useSplashUrl"`
	RedirectURL    string `json:"redirectUrl"`
	UseRedirectURL bool   `json:"useRedirectUrl"`
	WelcomeMessage string `json:"welcomeMessage"`
	SplashLogo     struct {
		Md5       string `json:"md5"`
		Extension string `json:"extension"`
	} `json:"splashLogo"`
	SplashImage struct {
		Md5 interface{} `json:"md5"`
	} `json:"splashImage"`
	SplashPrepaidFront struct {
		Md5 interface{} `json:"md5"`
	} `json:"splashPrepaidFront"`
}

// Display The Splash Page Settings For The Given SSID
func GetSplashPageSettings(networkId, number string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/ssids/%s/splash/settings",
		api.BaseUrl(), networkId, number)
	var datamodel SplashPageSettings
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
