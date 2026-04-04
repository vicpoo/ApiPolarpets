// habito_con_estado.go - NUEVO ARCHIVO en src/registro_habito/domain/entities/
package entities

import "time"

type HabitoConEstado struct {
	IDHabito        int32      `json:"id_habito"`
	Titulo          string     `json:"titulo"`
	Descripcion     string     `json:"descripcion"`
	Puntos          int32      `json:"puntos"`
	Completado      bool       `json:"completado"`
	FechaRealizada  *time.Time `json:"fecha_realizada,omitempty"`
	PuntosGenerados *int32     `json:"puntos_generados,omitempty"`
}