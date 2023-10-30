package domain

type Color string

const (
	Green  Color = "#97ca00"
	Red    Color = "#e05d44"
	Yellow Color = "#dfb317"
)

type TemplateValues struct {
	Color Color
	Env   string
	Tag   string
}

type AppInfo struct {
	Version   string
	CommitSha string
}
