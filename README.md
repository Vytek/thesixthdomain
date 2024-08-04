# thesixthdomain
Simulation War Games

![thesixthdomain](https://github.com/Vytek/thesixthdomain/blob/main/doc/Logo.jpeg?raw=true)

ITA:

## Cosa vorrei realizzare

I Board War Games 2D e 3D usano spesso semplici sistemi statistici per modelizzare a vario livello battaglie od intere campagne considerando il fattore casuale in vario modo. Ma se come dice Galileo "la natura è scritta in linguaggio matematico" allora anche la Gguerra è scritta in linguaggio matematico: sarà dunque necessario una "Fisica della guerra" che potrà dunque essere convertita in una vera e propria simulazione matematica grazie all'enorme potenza computazionale dei computer. Un campo di battaglia virtuale in cui diversi attori potranno affrontarsi e simulare un vero e proprio War Games simulativo basato su modelli matematici molto dettagliati.

![Tutto è equazioni](https://github.com/Vytek/thesixthdomain/blob/main/doc/all_i_see_are_equations.png?raw=true)

Questo "serius game" vuole invece modelizzare e simulare un campo di battaglia virtuale utilizzando modelli matematici reali applicabili in un conflitto. Il primo obiettivo è quello di descrivere in dettaglio la posizione ed il movimento di mezzi militari (come primo dominio quello navale) usando il sistema di coordinate previsto dal protocollo/standard DIS: https://open-dis.github.io/dis-tutorial/

Ogni mezzo è delineabile secondo questi moduli base (e per il momento ci occuperemo solo del primo): 

1. Sistemi di movimento
2. Sistemi di armamento
3. Sistemi di rilevamento attivi/passivi

## Achitettura software

*thesixthdomain* è un sistema multi utente/multi giocatore basato su architettura client/server classica a stella.

Using:

- https://github.com/Vytek/B1NukeBomber
- https://github.com/Vytek/radar2tacview
- https://github.com/Vytek/SeaWarfareSimulatorServer
- https://github.com/Vytek/SeaWarfareSimulatorClient
- https://github.com/Vytek/MayaVerseGolang

Services:

- https://openweathermap.org/api
- https://www.opentopodata.org/datasets/gebco2020/
- https://isitwater.com/
- https://github.com/open-meteo/open-meteo

Library:

- https://github.com/go-resty/resty
- https://github.com/nathan-osman/go-sunrise
- https://github.com/HectorMalot/omgo
- https://github.com/Vytek/tsd_passivesonar
