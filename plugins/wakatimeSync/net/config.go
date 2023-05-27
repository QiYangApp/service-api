package net

const (
	// domain 基础路由
	BASE_URL = "https://wakatime.com/api/v1/"

	// oath 认证域名
	OAUTH_URL = "https://wakatime.com/oauth/token"

	// user info 用户信息
	USER_INFO = BASE_URL + "users/current"

	// 持续时间
	DURATION = BASE_URL + "users/current/durations"

	// 心跳
	HEART_BEATS = BASE_URL + "users/current/heartbeats"

	//项目
	PROJECTS       = BASE_URL + "users/current/projects"
	SUMMARY        = BASE_URL + "users/current/summaries"
	PROJECT_DETAIL = BASE_URL + "users/current/projects/"
)
