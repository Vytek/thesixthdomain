L'equazione del radar, nota anche come l'equazione del radar di Friss, è utilizzata per calcolare la potenza del segnale ricevuto da un'antenna radar. La forma base dell'equazione del radar monostatico (trasmettitore e ricevitore nella stessa posizione) è:

\[ P_r = \frac{P_t G_t G_r \lambda^2 \sigma}{(4\pi)^3 R^4} \]

Dove:
- \( P_r \) è la potenza ricevuta,
- \( P_t \) è la potenza trasmessa,
- \( G_t \) è il guadagno dell'antenna trasmittente,
- \( G_r \) è il guadagno dell'antenna ricevente,
- \( \lambda \) è la lunghezza d'onda del segnale trasmesso,
- \( \sigma \) è la sezione radar equivalente dell'obiettivo (RCS),
- \( R \) è la distanza tra il radar e l'obiettivo.

Ecco una funzione in Go che implementa l'equazione del radar:

```go
package main

import (
	"fmt"
	"math"
)

// RadarEquation calcola la potenza ricevuta da un'antenna radar.
// Parametri:
// - Pt: Potenza trasmessa (in Watt)
// - Gt: Guadagno dell'antenna trasmittente (dimensionless)
// - Gr: Guadagno dell'antenna ricevente (dimensionless)
// - lambda: Lunghezza d'onda (in metri)
// - sigma: Sezione radar equivalente dell'obiettivo (in metri quadri)
// - R: Distanza tra il radar e l'obiettivo (in metri)
func RadarEquation(Pt, Gt, Gr, lambda, sigma, R float64) float64 {
	// Costante (4π)^3
	constant := math.Pow(4*math.Pi, 3)
	
	// Calcolo della potenza ricevuta
	Pr := (Pt * Gt * Gr * math.Pow(lambda, 2) * sigma) / (constant * math.Pow(R, 4))
	
	return Pr
}

func main() {
	// Esempio di utilizzo della funzione RadarEquation
	Pt := 1.0         // Potenza trasmessa in Watt
	Gt := 30.0        // Guadagno dell'antenna trasmittente (dimensionless)
	Gr := 30.0        // Guadagno dell'antenna ricevente (dimensionless)
	lambda := 0.03    // Lunghezza d'onda in metri (es. 10 GHz)
	sigma := 1.0      // Sezione radar equivalente dell'obiettivo in metri quadri
	R := 10000.0      // Distanza in metri

	Pr := RadarEquation(Pt, Gt, Gr, lambda, sigma, R)
	fmt.Printf("La potenza ricevuta è: %e Watt\n", Pr)
}
```

In questo esempio:
- La potenza trasmessa (\(P_t\)) è 1 Watt.
- Il guadagno dell'antenna trasmittente (\(G_t\)) è 30 (dimensionless).
- Il guadagno dell'antenna ricevente (\(G_r\)) è 30 (dimensionless).
- La lunghezza d'onda (\(\lambda\)) è 0.03 metri (equivalente a 10 GHz di frequenza).
- La sezione radar equivalente (\(\sigma\)) è 1 metro quadrato.
- La distanza (\(R\)) è 10000 metri (10 km).

La funzione `RadarEquation` calcola la potenza ricevuta utilizzando questi parametri e la stampa in output. Puoi modificare i valori degli input nel `main` per testare con diversi scenari.