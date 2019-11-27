/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package plugin

import "github.com/spf13/pflag"

type Plugin interface {
	Version() string
}

type InjectsCommandName interface {
	InjectCommandName(commandName string)
}

type Deprecated interface {
	DeprecationWarning() string
}

type ProjectScaffolder interface {
	Plugin
	ProjectHelp() string
	ProjectExample() string
	BindProjectFlags(*pflag.FlagSet)
	ScaffoldProject() error
}

type APIScaffolder interface {
	Plugin
	APIHelp() string
	APIExample() string
	BindAPIFlags(*pflag.FlagSet)
	ScaffoldAPI() error
}

type WebhookScaffolder interface {
	Plugin
	WebhookHelp() string
	WebhookExample() string
	BindWebhookFlags(flag *pflag.FlagSet)
	ScaffoldWebhook() error
}
