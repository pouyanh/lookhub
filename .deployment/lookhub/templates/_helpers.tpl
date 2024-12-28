{{/*
Expand the name of the chart.
*/}}
{{- define "lookhub.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "lookhub.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "lookhub.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "lookhub.labels" -}}
helm.sh/chart: {{ include "lookhub.chart" . }}
{{ include "lookhub.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "lookhub.selectorLabels" -}}
app.kubernetes.io/name: {{ include "lookhub.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Expand the namespace of the release.
Allows overriding it for multi-namespace deployments in combined charts.
*/}}
{{- define "lookhub.namespace" -}}
{{- default .Release.Namespace .Values.namespaceOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "lookhub.image" -}}
{{- printf "%s/%s:%s" .Values.image.registry .Values.image.repository (default .Chart.AppVersion .Values.image.tag) }}
{{- end -}}
