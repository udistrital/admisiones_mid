package models

type Tag struct {
	Selected bool
	Required bool
}

type SuiteInscripcion struct {
	Id                float64 `json:"Id"`
	Activo            bool    `json:"Activo"`
	DependenciaId     float64 `json:"DependenciaId"`
	ListaTags         string  `json:"ListaTags"`
	PeriodoId         float64 `json:"PeriodoId"`
	TipoInscripcionId float64 `json:"TipoInscripcionId"`
	FechaCreacion     string  `json:"FechaCreacion,omitempty"`
	FechaModificacion string  `json:"FechaModificacion,omitempty"`
}

var TagsInscripcionPrograma = map[string]Tag{
	"info_persona": {
		Selected: false,
		Required: false,
	},
	"formacion_academica": {
		Selected: false,
		Required: false,
	},
	"idiomas": {
		Selected: false,
		Required: false,
	},
	"experiencia_laboral": {
		Selected: false,
		Required: false,
	},
	"produccion_academica": {
		Selected: false,
		Required: false,
	},
	"documento_programa": {
		Selected: false,
		Required: false,
	},
	"descuento_matricula": {
		Selected: false,
		Required: false,
	},
	"propuesta_grado": {
		Selected: false,
		Required: false,
	},
	"perfil": {
		Selected: false,
		Required: false,
	},
}
