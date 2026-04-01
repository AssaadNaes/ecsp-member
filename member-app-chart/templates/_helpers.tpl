{{/*
Expand the name of the chart.
*/}}
/*{{- define "member-app-chart.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}*/

{{/*
Loop threw all labels for the frontend deployment
*/}}
{{- define "frontend.labels" -}}
{{- range .Values.frontend.labels }}
    {{- range $key, $value := . }}
{{ $key }}: {{ $value | quote }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Loop threw all labels for the backend deployment
*/}}
{{- define "backend.labels" -}}
{{- range .Values.backend.labels }}
    {{- range $key, $value := . }}
{{ $key }}: {{ $value | quote }}
{{- end }}
{{- end }}
{{- end }}
