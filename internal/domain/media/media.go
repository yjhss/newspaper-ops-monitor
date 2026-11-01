package media

import "github.com/yjhss/newspaper-ops-monitor/internal/domain/workflow"

type LayoutType string

const (
	LayoutNormal  LayoutType = "NORMAL"
	LayoutTabloid LayoutType = "TABLOID"
)

type IngressType string

const (
	IngressWebhardPDF IngressType = "WEBHARD_PDF"
	IngressInputPDF   IngressType = "INPUT_PDF"
	IngressInputTIFF  IngressType = "INPUT_TIFF"
	IngressInputMedia IngressType = "INPUT_MEDIA"
)

type Profile struct {
	ID      string
	Name    string
	Enabled bool
	Layout  LayoutType
	Ingress IngressType
	Paths   map[workflow.Stage]string
}
