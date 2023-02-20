package global

{{- if .HasGlobal }}

import "github.com/defeng-hub/ByOfficeAutomatic/server/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}