package mailuquadlet

const containerTmpl = `[Unit]{{ $uuid := .Annotations.uuid }}
Description=Mailu {{ .Name }} container

[Container]
Image={{ .Image }}
AutoUpdate=registry

{{ range .Volumes }}Volume={{ . }}
{{ end }}
EnvironmentFile=/etc/mailu/{{ $uuid }}.env
{{ if eq .Name "admin" }}Environment=I_KNOW_MY_SETUP_DOESNT_FIT_REQUIREMENTS_AND_WONT_FILE_ISSUES_WITHOUT_PATCHES=true
{{ end }}
HostName={{ .Name }}
{{ range $key, $value := .Networks }}Network={{ $uuid }}-{{ $key }}.network
{{ end }}
{{ if .Annotations.ip }}IP={{ .Annotations.ip }}{{ end }}
{{ range .Ports }}PublishPort={{ if contains .HostIP ":" }}[{{ .HostIP }}]{{ else }}{{.HostIP}}{{ end }}:{{ .Target }}:{{ .Published }}
{{ end }}
{{ range .DNS }}DNS={{ . }}{{ end }}

[Service]
Restart=always
TimeoutStartSec=900

[Install]
WantedBy=default.target
`

const networkTmpl = `[Network]
{{ range .Ipam.Config }}Subnet={{ .Subnet }}
{{ end }}
{{ if .Internal }}Internal=true{{ end }}
{{ if .EnableIPv6 }}IPv6=true{{ end }}
`

const volumeTmpl = `[Volume]`
