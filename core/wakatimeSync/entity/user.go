package entity

import "time"

type UserRequest struct{}

type UserResponse struct {
	Data struct {
		Bio                    interface{} `json:"bio"`
		City                   interface{} `json:"city"`
		ColorScheme            string      `json:"color_scheme"`
		CreatedAt              time.Time   `json:"created_at"`
		DateFormat             string      `json:"date_format"`
		DefaultDashboardRange  string      `json:"default_dashboard_range"`
		DisplayName            string      `json:"display_name"`
		DurationsSliceBy       string      `json:"durations_slice_by"`
		Email                  string      `json:"email"`
		FullName               interface{} `json:"full_name"`
		GithubUsername         interface{} `json:"github_username"`
		HasPremiumFeatures     bool        `json:"has_premium_features"`
		HumanReadableWebsite   interface{} `json:"human_readable_website"`
		Id                     string      `json:"id"`
		InvoiceIdFormat        string      `json:"invoice_id_format"`
		IsEmailConfirmed       bool        `json:"is_email_confirmed"`
		IsEmailPublic          bool        `json:"is_email_public"`
		IsHireable             bool        `json:"is_hireable"`
		IsOnboardingFinished   bool        `json:"is_onboarding_finished"`
		LanguagesUsedPublic    bool        `json:"languages_used_public"`
		LastHeartbeatAt        time.Time   `json:"last_heartbeat_at"`
		LastPlugin             string      `json:"last_plugin"`
		LastPluginName         string      `json:"last_plugin_name"`
		LastProject            string      `json:"last_project"`
		LinkedinUsername       interface{} `json:"linkedin_username"`
		Location               interface{} `json:"location"`
		LoggedTimePublic       bool        `json:"logged_time_public"`
		MeetingsOverCoding     bool        `json:"meetings_over_coding"`
		ModifiedAt             interface{} `json:"modified_at"`
		NeedsPaymentMethod     bool        `json:"needs_payment_method"`
		Photo                  string      `json:"photo"`
		PhotoPublic            bool        `json:"photo_public"`
		Plan                   string      `json:"plan"`
		ProfileUrl             string      `json:"profile_url"`
		ProfileUrlEscaped      string      `json:"profile_url_escaped"`
		PublicEmail            interface{} `json:"public_email"`
		PublicProfileTimeRange string      `json:"public_profile_time_range"`
		ShareAllTimeBadge      interface{} `json:"share_all_time_badge"`
		ShareLastYearDays      interface{} `json:"share_last_year_days"`
		ShowMachineNameIp      bool        `json:"show_machine_name_ip"`
		TimeFormat24Hr         interface{} `json:"time_format_24hr"`
		TimeFormatDisplay      string      `json:"time_format_display"`
		Timeout                int         `json:"timeout"`
		Timezone               string      `json:"timezone"`
		TwitterUsername        interface{} `json:"twitter_username"`
		Username               interface{} `json:"username"`
		Website                interface{} `json:"website"`
		WeekdayStart           int         `json:"weekday_start"`
		WritesOnly             bool        `json:"writes_only"`
	} `json:"data"`
}
