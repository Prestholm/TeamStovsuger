
# Obligatorisk oppgave 4: 

Medlemmer: 
- Henning Van Ta.
- Victor Prestholm.
- Arian Nuland. 


### Systemspesifikasjon

Applikasjonen er en tjeneste som gjør det enkelt for brukeren å følge med på verdien av ulike typer kryptovaluta. Den tar i bruk live data fra coinmarketcap.com. Her får brukeren blant annet informasjon om forskjellige kryptovaluta, forkortelser, rangering, pris, antall i omløp og prosentvis endring over en gitt tidsperiode. 

På startsiden skriver brukeren inn ønsket valuta, eller trykker på knappen for å få opp en ny side med oversikt i tabellform. Tabellen i programmet viser 100 forskjellige valuta i rangert rekkefølge. Når brukeren skriver ønsket valuta får han opp en side med mer informasjon om valutaen. 

For å gjøre applikasjonen mer dynamisk har vi implementert en converter som konverterer verdien av forskjellig valuta, for eksempel hvor mye 1 Bitcoin er verdt i USD eller motsatt. Vi har også lagt til en ekstra API som henter logoene til forskjellig kryptovaluta. Applikasjonen oppdaterer ikke currency-data automatisk. Dette kunne vi fått til ved å implementere en websocket og jQuery AJAX. 


### Systemarkitektur

Nodene som er involvert i systemet er klient, webserver og databaseserver. Applikasjonen oppretter en webserver på port ":8080". Klienten sender så en forespørsel til webserveren som videre henter live-data fra databaseserveren til coinmarketcap. Deretter blir denne returnert til klienten. 


![Mac](https://github.com/Prestholm/TeamStovsuger/blob/master/Oblig4/Bilder/32323871_1815296048513457_1596700920316428288_n.png)


### Hvordan ta i bruk applikasjonen

1. Skriv "go run server.go" i terminalen.
2. Gå til "localhost:8080" i nettleseren.
3. Skriv inn kryptovaluta du ønsker informasjon om, eller trykk på knappen for å få oversikt over live-rangeringen i tabellform. 
4. For å få opp converteren må du søke etter en valuta. 


### Testing

Vi har utviklet enhetstester (se "server_test.go") for å sikre at kildekoden oppfører seg som den skal. Her tester vi blant annet om filene i mappen "Templates" er tilstede og om URL-linken vi bruker inneholder JSON. 







