package types

const (
	RoleUnknown Role = iota
	RoleAdmin

	// RoleBasic is the default role. Leaving a little space for future roles.
	RoleBasic Role = 10
)

type Role int

func (u Role) HasRole(role Role) bool {
	return u != RoleUnknown && role >= u
}

type User struct {
	Metadata
	Username                   string `json:"username,omitempty"`
	Role                       Role   `json:"role,omitempty"`
	ExplicitAdmin              bool   `json:"explicitAdmin,omitempty"`
	Email                      string `json:"email,omitempty"`
	IconURL                    string `json:"iconURL,omitempty"`
	Timezone                   string `json:"timezone,omitempty"`
	CurrentAuthProvider        string `json:"currentAuthProvider,omitempty"`
	LastActiveDay              Time   `json:"lastActiveDay,omitzero"`
	Internal                   bool   `json:"internal,omitempty"`
	DailyPromptTokensLimit     int    `json:"dailyPromptTokensLimit,omitempty"`
	DailyCompletionTokensLimit int    `json:"dailyCompletionTokensLimit,omitempty"`
	DisplayName                string `json:"displayName,omitempty"`
}

type UserList List[User]
