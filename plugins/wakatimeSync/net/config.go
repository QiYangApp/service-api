package net

const (
	// BaseUrl domain 基础路由
	BaseUrl = "https://wakatime.com/api/v1/"

	// OauthUrl oath 认证域名
	OauthUrl = "https://wakatime.com/oauth/token"

	// UserInfo 用户信息
	UserInfo = BaseUrl + "users/current"

	// DURATION 持续时间
	DURATION = BaseUrl + "users/current/durations"

	// HeartBeats 心跳
	HeartBeats = BaseUrl + "users/current/heartbeats"

	// PROJECTS 项目
	PROJECTS = BaseUrl + "users/current/projects"
	SUMMARY  = BaseUrl + "users/current/summaries"
)
