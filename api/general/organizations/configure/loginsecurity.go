package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	"log"
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
func GetLoginSecurity(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/licenses", api.BaseUrl(),
		organizationId)

	var datamodel = LoginSecurity{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
