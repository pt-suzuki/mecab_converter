package auth

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestトランスレータでリクエストをIDに変換(t *testing.T) {
	translator := TranslatorImpl()

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authenticated", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Ik9FWTVPVVF6UmtSR05EUTJNemxDTmpVek5FVXhNVEV3TXpGRVFURkNSVGd6UmpFd05qQTVNUSJ9.eyJpc3MiOiJodHRwczovL2xvZ2luLmNhcmVlcnBvY2tldC5qcC8iLCJzdWIiOiJhdXRoMHw1ZWM0OGU3MDBkOTQ2YjBjYjg4ZDE4YjMiLCJhdWQiOlsiaHR0cHM6Ly9hc2lhLW5vcnRoZWFzdDEtcHJkLWNscy5jbG91ZGZ1bmN0aW9ucy5uZXQvIiwiaHR0cHM6Ly9wcmQtY2xzLmF1dGgwLmNvbS91c2VyaW5mbyJdLCJpYXQiOjE1OTk1NzQ1NjYsImV4cCI6MTU5OTU3ODA2NiwiYXpwIjoiUk5iOUxpNm5yaUNzVEJLZmFjQmZPZUZEOVNZYVdiZ0wiLCJzY29wZSI6Im9wZW5pZCBlbWFpbCBwcm9maWxlIG9mZmxpbmVfYWNjZXNzIiwiZ3R5IjoicGFzc3dvcmQifQ.Ubt55-9znqlHpOf-DnSax3y4b93_aTTpXXM3M_lMGuydSloxCoazXjrvgxx3ks7rktwQjFfJYe297T_r88LVkuQ_NR1OSMA6ItwTaIARCyoday1QWN4vOAUu2uUrQt83RWZPO7xFHZ-Tz-OFOUH4pEczP9Z5nAvPdH_YvyROB3ZTk8z1hHuqNKLaM65u7S6uSklFYxe99mrXCUIOuHr7BWCMlrtu19NRIRworYZBjujqCi3_kAadILeKfJIOlKPiRBlIzD08RdpdeK3ah1KEQJvQgcRFBmmUp4uhZqL2l1SFgzYpVMPxKGu8ooLzIKOhglGfOB5P3UTd517FcPH2-A")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	id := translator.TranslateRequestToAuthId(c)
	assert.Equal(t, "auth0|5ec48e700d946b0cb88d18b3", id)
}
