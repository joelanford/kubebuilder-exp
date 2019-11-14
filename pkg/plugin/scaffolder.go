package plugin

import (
	"github.com/spf13/pflag"
)

type ProjectScaffolder interface {
	Plugin
	ProjectHelp() string
	ProjectExample() string
	ScaffoldProject() error
	BindProjectFlags(*pflag.FlagSet)
}

type APIScaffolder interface {
	Plugin
	APIHelp() string
	APIExample() string
	ScaffoldAPI() error
	BindAPIFlags(*pflag.FlagSet)
}

type WebhookScaffolder interface {
	Plugin
	WebhookHelp() string
	WebhookExample() string
	BindWebhookFlags(flag *pflag.FlagSet)
	ScaffoldWebhook() error
}
