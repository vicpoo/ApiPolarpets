// registro_habito.go
package entities

import "time"

type RegistroHabito struct {
	IDRegistroHabito int32     `json:"id_registro_habito" gorm:"column:id_registro_habito;primaryKey;autoIncrement"`
	IDHabito         int32     `json:"id_habito" gorm:"column:id_habito"`
	FechaRealizada   time.Time `json:"fecha_realizada" gorm:"column:fecha_realizada"`
	PuntosGenerados  int32     `json:"puntos_generados" gorm:"column:puntos_generados"`
}

// Setters
func (rh *RegistroHabito) SetIDRegistroHabito(id int32) {
	rh.IDRegistroHabito = id
}

func (rh *RegistroHabito) SetIDHabito(idHabito int32) {
	rh.IDHabito = idHabito
}

func (rh *RegistroHabito) SetFechaRealizada(fechaRealizada time.Time) {
	rh.FechaRealizada = fechaRealizada
}

func (rh *RegistroHabito) SetPuntosGenerados(puntosGenerados int32) {
	rh.PuntosGenerados = puntosGenerados
}

// Getters
func (rh *RegistroHabito) GetIDRegistroHabito() int32 {
	return rh.IDRegistroHabito
}

func (rh *RegistroHabito) GetIDHabito() int32 {
	return rh.IDHabito
}

func (rh *RegistroHabito) GetFechaRealizada() time.Time {
	return rh.FechaRealizada
}

func (rh *RegistroHabito) GetPuntosGenerados() int32 {
	return rh.PuntosGenerados
}

// Constructor
func NewRegistroHabito(idHabito int32, fechaRealizada time.Time, puntosGenerados int32) *RegistroHabito {
	return &RegistroHabito{
		IDHabito:        idHabito,
		FechaRealizada:  fechaRealizada,
		PuntosGenerados: puntosGenerados,
	}
}