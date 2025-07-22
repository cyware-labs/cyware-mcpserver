package co

const user_details_endpoint = "/cpapi/rest-auth/user-details/"

type LoggedInUserDetailsResponse struct {
	Email                string   `json:"email"`
	FullName             string   `json:"full_name"`
	ID                   string   `json:"id"`
	IsActive             bool     `json:"is_active"`
	PermissionGroups     []string `json:"permission_groups"`
	PermissionGroupsData []struct {
		Created     string `json:"created"`
		CreatedBy   string `json:"created_by"`
		Description string `json:"description"`
		ID          string `json:"id"`
		IsActive    bool   `json:"is_active"`
		IsEditable  bool   `json:"is_editable"`
		MetaInfo    struct {
			Cftr struct {
				Enabled bool `json:"enabled"`
			} `json:"cftr"`
			Co struct {
				Enabled     bool `json:"enabled"`
				Permissions struct {
					ApproveExecutionViaEmailPlaybook  string `json:"approve_execution_via_email_playbook"`
					ApproveExecutionViaMobilePlaybook string `json:"approve_execution_via_mobile_playbook"`
					CreateOpenapi                     string `json:"create_openapi"`
				} `json:"permissions"`
			} `json:"co"`
			Csap struct {
				Enabled bool `json:"enabled"`
			} `json:"csap"`
			Csapm struct {
				Enabled bool `json:"enabled"`
			} `json:"csapm"`
			Ctix struct {
				Enabled bool `json:"enabled"`
			} `json:"ctix"`
		} `json:"meta_info"`
		Modified    string   `json:"modified"`
		ModifiedBy  string   `json:"modified_by"`
		Name        string   `json:"name"`
		Permissions []string `json:"permissions"`
		TenantID    string   `json:"tenant_id"`
	} `json:"permission_groups_data"`
	Permissions []string `json:"permissions"`
	Phone       struct {
		CountryCode string `json:"country_code"`
		Ext         int    `json:"ext"`
		PhNumber    int    `json:"ph_number"`
	} `json:"phone"`
	Preferences struct {
		DarkMode     bool   `json:"dark_mode"`
		DateFormat   string `json:"date_format"`
		DateFormatUI string `json:"date_format_ui"`
		ShowSynopsis bool   `json:"show_synopsis"`
		TimeFormat   string `json:"time_format"`
		TimeZone     string `json:"time_zone"`
	} `json:"preferences"`
	PreferredWorkspace struct {
		Code string `json:"code"`
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"preferred_workspace"`
	PreferredWorkspaceID string `json:"preferred_workspace_id"`
	Username             string `json:"username"`
	TenantID             string `json:"tenant_id"`
}

func GetLoggedInUserDetails() *LoggedInUserDetailsResponse {
	user_details_resp := LoggedInUserDetailsResponse{}
	CO_CLIENT.MakeRequest("GET", user_details_endpoint, nil, &user_details_resp, nil, nil)
	return &user_details_resp
}
