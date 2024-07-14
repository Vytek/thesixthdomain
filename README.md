# thesixthdomain
Simulation War Games

![thesixthdomain](https://github.com/Vytek/thesixthdomain/blob/main/doc/Logo.jpeg?raw=true)

ITA:

## Cosa vorrei realizzare

I Board War Games 2D e 3D usano spesso semplici sistemi statistici per modelizzare a vario livello battaglie od intere campagne considerando il fattore casuale in vario modo. 
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

Library:

- https://github.com/go-resty/resty
