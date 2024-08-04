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

Ci sono diverse varianti dell'equazione del radar che possono essere utilizzate a seconda del tipo di radar e delle specifiche applicazioni. Alcune delle varianti più comuni includono:

1. **Equazione del radar bistatico**:
   Questa equazione viene utilizzata quando il trasmettitore e il ricevitore sono in posizioni diverse.

2. **Equazione del radar con perdite**:
   Tiene conto delle perdite di sistema e atmosferiche.

3. **Equazione del radar per un target non cooperativo**:
   Considera il fatto che il target potrebbe non riflettere il segnale in modo ideale.

### Equazione del radar bistatico

L'equazione del radar bistatico è una generalizzazione dell'equazione del radar monostatico, dove il trasmettitore e il ricevitore sono in posizioni diverse:

\[ P_r = \frac{P_t G_t G_r \lambda^2 \sigma}{(4\pi)^3 R_t^2 R_r^2} \]

Dove:
- \( R_t \) è la distanza dal trasmettitore al target.
- \( R_r \) è la distanza dal target al ricevitore.

### Equazione del radar con perdite

Per considerare le perdite di sistema (\(L_s\)) e atmosferiche (\(L_a\)):

\[ P_r = \frac{P_t G_t G_r \lambda^2 \sigma}{(4\pi)^3 R^4 L_s L_a} \]

Dove:
- \( L_s \) è la perdita di sistema.
- \( L_a \) è la perdita atmosferica.

### Implementazione in Go

Ecco una funzione Go per ciascuna delle varianti descritte sopra.

```go
package main

import (
	"fmt"
	"math"
)

// RadarEquationMonostatic calcola la potenza ricevuta per un radar monostatico.
func RadarEquationMonostatic(Pt, Gt, Gr, lambda, sigma, R float64) float64 {
	constant := math.Pow(4*math.Pi, 3)
	Pr := (Pt * Gt * Gr * math.Pow(lambda, 2) * sigma) / (constant * math.Pow(R, 4))
	return Pr
}

// RadarEquationBistatic calcola la potenza ricevuta per un radar bistatico.
func RadarEquationBistatic(Pt, Gt, Gr, lambda, sigma, Rt, Rr float64) float64 {
	constant := math.Pow(4*math.Pi, 3)
	Pr := (Pt * Gt * Gr * math.Pow(lambda, 2) * sigma) / (constant * math.Pow(Rt, 2) * math.Pow(Rr, 2))
	return Pr
}

// RadarEquationWithLosses calcola la potenza ricevuta considerando le perdite di sistema e atmosferiche.
func RadarEquationWithLosses(Pt, Gt, Gr, lambda, sigma, R, Ls, La float64) float64 {
	constant := math.Pow(4*math.Pi, 3)
	Pr := (Pt * Gt * Gr * math.Pow(lambda, 2) * sigma) / (constant * math.Pow(R, 4) * Ls * La)
	return Pr
}

func main() {
	Pt := 1.0         // Potenza trasmessa in Watt
	Gt := 30.0        // Guadagno dell'antenna trasmittente (dimensionless)
	Gr := 30.0        // Guadagno dell'antenna ricevente (dimensionless)
	lambda := 0.03    // Lunghezza d'onda in metri (es. 10 GHz)
	sigma := 1.0      // Sezione radar equivalente dell'obiettivo in metri quadri
	R := 10000.0      // Distanza in metri
	Rt := 5000.0      // Distanza trasmettitore-target in metri
	Rr := 5000.0      // Distanza target-ricevitore in metri
	Ls := 2.0         // Perdite di sistema (dimensionless)
	La := 1.5         // Perdite atmosferiche (dimensionless)

	PrMonostatic := RadarEquationMonostatic(Pt, Gt, Gr, lambda, sigma, R)
	PrBistatic := RadarEquationBistatic(Pt, Gt, Gr, lambda, sigma, Rt, Rr)
	PrWithLosses := RadarEquationWithLosses(Pt, Gt, Gr, lambda, sigma, R, Ls, La)

	fmt.Printf("Potenza ricevuta (monostatico): %e Watt\n", PrMonostatic)
	fmt.Printf("Potenza ricevuta (bistatico): %e Watt\n", PrBistatic)
	fmt.Printf("Potenza ricevuta (con perdite): %e Watt\n", PrWithLosses)
}
```

In questo codice:
- `RadarEquationMonostatic` calcola la potenza ricevuta per un radar monostatico.
- `RadarEquationBistatic` calcola la potenza ricevuta per un radar bistatico.
- `RadarEquationWithLosses` calcola la potenza ricevuta considerando le perdite di sistema e atmosferiche.

Puoi modificare i valori degli input nel `main` per testare diversi scenari e ottenere i risultati desiderati.