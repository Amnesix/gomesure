// Le Package mesure implémentre de façon très simple un système permettant de
// mesurer le temps d'exécution d'une routine
package gomesure

import "time"

// Le type Mesure est la seule structure nécessaire.
type Mesure struct {
	start  int64
	mesure float64
}

// Starttime permet d'initialiser une mesure
func (t *Mesure) Starttime() {
	t.start = time.Now().UnixNano()
}

// Elapsedtime permet d'arrêter la mesure
func (t *Mesure) Elapsedtime() {
	t.mesure = float64(time.Now().UnixNano()-t.start) / 1000000.
}

// GetMillisecondes permet de récupérer le résultat en mili seconde.
func (t *Mesure) GetMillisecondes() float64 {
	return t.mesure
}

// GetSecondes permet de récupérer le résultat en mili seconde.
func (t *Mesure) GetSecondes() float64 {
	return t.mesure / 1000.
}
