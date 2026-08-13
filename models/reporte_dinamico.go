package models

type ReporteEstructura struct {
	Proyecto          int64    `json:"Proyecto"`
	Periodo           int64    `json:"Periodo"`
	TipoReporte       int64    `json:"Reporte"`
	Columnas          []string `json:"Columnas"`
	TipoInscripcion   int64    `json:"TipoInscripcion,omitempty"`
	EstadoInscripcion string   `json:"EstadoInscripcion,omitempty"`
}

type ReciboResponse struct {
	ReciboCollection ReciboCollection `json:"reciboCollection"`
}

type ReciboCollection struct {
	Recibo []Recibo `json:"recibo"`
}

type Recibo struct {
	Estado              string  `json:"estado"`
	Ano                 int     `json:"ano,string"`
	Cuota               int     `json:"cuota,string"`
	Periodo             int     `json:"periodo,string"`
	FechaPagado         string  `json:"fecha_pagado"`
	Secuencia           int     `json:"secuencia,string"`
	Documento           string  `json:"documento"`
	FechaOrdinario      string  `json:"fecha_ordinario"`
	Pago                string  `json:"pago"`
	Nombre              string  `json:"nombre"`
	Fecha               string  `json:"fecha"`
	ValorExtraordinario float64 `json:"valor_extraordinario,string"`
	Observaciones       string  `json:"observaciones"`
	Carrera             int     `json:"carrera,string"`
	ValorPagado         float64 `json:"valor_pagado,string"`
	FechaExtraordinario string  `json:"fecha_extraordinario"`
	ValorOrdinario      float64 `json:"valor_ordinario,string"`
}
