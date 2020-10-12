package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type LoginSecurity struct {
	EnforcePasswordExpiration bool     `json:"enforcePasswordExpiration"`
	PasswordExpirationDays    int      `json:"passwordExpirationDays"`
	EnforceDifferentPasswords bool     `json:"enforceDifferentPasswords"`
	NumDifferentPasswords     int      `json:"numDifferentPasswords"`
	EnforceStrongPasswords    bool     `json:"enforceStrongPasswords"`
	EnforceAccountLockout     bool     `json:"enforceAccountLockout"`
	AccountLockoutAttempts    int      `json:"accountLockoutAttempts"`
	EnforceIdleTimeout        bool     `json:"enforceIdleTimeout"`
	IdleTimeoutMinutes        int      `json:"idleTimeoutMinutes"`
	EnforceTwoFactorAuth      bool     `json:"enforceTwoFactorAuth"`
	EnforceLoginIPRanges      bool     `json:"enforceLoginIpRanges"`
	LoginIPRanges             []string `json:"loginIpRanges"`
}


// Returns The Login Security Settings For An Organization
func GetLoginSecurity(organizationId string) (LoginSecurity, interface{}) {
	baseurl := fmt.Sprintf("%s/organizations/%s/licenses", api.BaseUrl(),
		organizationId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)

	var results = LoginSecurity{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}