package authboss

import (
	"html/template"
	"io"
	"io/ioutil"
	"net/smtp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	layoutTpl      = "layout.tpl"
	layoutEmailTpl = "layoutEmail.tpl"
)

// Cfg is the singleton instance of Config
var Cfg *Config = NewConfig()

// Config holds all the configuration for both authboss and it's modules.
type Config struct {
	// MountPath is the path to mount authboss's routes at (eg /auth).
	MountPath string
	// ViewsPath is the path to search for overridden templates.
	ViewsPath string
	// HostName is the host of the web application (eg https://www.happiness.com:8080) for e-mail url generation.
	HostName string
	// BCryptCost is the cost of the bcrypt password hashing function.
	BCryptCost int

	// PrimaryID is the primary identifier of the user. Set to one of:
	// authboss.StoreEmail, authboss.StoreUsername (StoreEmail is default)
	PrimaryID string

	Layout          *template.Template
	LayoutEmail     *template.Template
	LayoutDataMaker ViewDataMaker

	AuthLogoutRoute       string
	AuthLoginSuccessRoute string

	RecoverRedirect             string
	RecoverInitiateSuccessFlash string
	RecoverTokenDuration        time.Duration
	RecoverTokenExpiredFlash    string
	RecoverFailedErrorFlash     string

	Policies      []Validator
	ConfirmFields []string

	ExpireAfter time.Duration

	LockAfter    int
	LockWindow   time.Duration
	LockDuration time.Duration

	EmailFrom          string
	EmailSubjectPrefix string
	SMTPAddress        string
	SMTPAuth           smtp.Auth

	XSRFName  string
	XSRFMaker XSRF

	Storer            Storer
	CookieStoreMaker  CookieStoreMaker
	SessionStoreMaker SessionStoreMaker
	LogWriter         io.Writer
	Callbacks         *Callbacks
	Mailer            Mailer
}

func NewConfig() *Config {
	return &Config{
		MountPath:  "/",
		ViewsPath:  "/",
		HostName:   "localhost:8080",
		BCryptCost: bcrypt.DefaultCost,

		PrimaryID: StoreEmail,

		Layout:      template.Must(template.New("").Parse(`<html><body>{{template "authboss" .}}</body></html>`)),
		LayoutEmail: template.Must(template.New("").Parse(`<html><body>{{template "authboss" .}}</body></html>`)),

		AuthLogoutRoute:       "/login",
		AuthLoginSuccessRoute: "/",

		Policies: []Validator{
			Rules{
				FieldName:       "username",
				Required:        true,
				MinLength:       2,
				MaxLength:       4,
				AllowWhitespace: false,
			},
			Rules{
				FieldName: "password",
				Required:  true,
				MinLength: 4,
				MaxLength: 8,

				AllowWhitespace: false,
			},
		},
		ConfirmFields: []string{
			StoreEmail, ConfirmPrefix + StoreEmail,
			StorePassword, ConfirmPrefix + StorePassword,
		},

		RecoverRedirect:             "/login",
		RecoverInitiateSuccessFlash: "An email has been sent with further insructions on how to reset your password",
		RecoverTokenDuration:        time.Duration(24) * time.Hour,
		RecoverTokenExpiredFlash:    "Account recovery request has expired.  Please try agian.",
		RecoverFailedErrorFlash:     "Account recovery has failed.  Please contact tech support.",

		LogWriter: ioutil.Discard,
		Callbacks: NewCallbacks(),
		Mailer:    LogMailer(ioutil.Discard),
	}
}
