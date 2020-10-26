package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
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
func GetSplashPageSettings(networkId, number string) (SplashPageSettings, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/ssids/%s/splash/settings",
		api.BaseUrl(), networkId, number)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var ssid SplashPageSettings
	user_agent.UnMarshalJSON(session.Body, &ssid)
	traceback := user_agent.TraceBack(session)
	return ssid, traceback
}

