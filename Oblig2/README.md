
# Obligatorisk oppgave 2: 

Medlemmer: 
- Henning Van Ta.
- Victor Prestholm.
- Arian Nuland. 

## Oppgave 1. 

### A. Skriv et Golang program med navn fileinfo.go. 
Se "fileinfo.go" i mappen "Oblig2/src/oppgave1/". Programmet returnerer detaljert informasjon om en fil. Den skriver ut en error om det ikke er valgt noen fil. 

![Mac](https://github.com/Prestholm/TeamStovsuger/blob/master/Oblig2/src/Bilder/fileinfo.png)

### Oppgave B. Lag et kjørbart program (build) av fileinfo.go som kan utføres fra kommandolinje.
I denne oppgaven bruker vi command-line flag parsing. Alt man trenger å gjøre er å builde og så kjøre den filen du vil ha info om ved "fileinfo -f [filnavn]", eller ved "go run fileinfo.go -f [filnavn]" (men dette lager ikke en executable). Slik henter programmet info om filen, dens størrelse og hva slags type fil etc.

Oppgave 1b er gjort litt anderledes enn de andre oppgavene. Her bruker vi Flags for å finne filnavn. Det er ikke nødvendig å bruke flags her, oppgaven kan også løses ved å bruke vanlig os.args[1] som i oppgave 2. 


## Oppgave 2. 

### A. Skriv et program med navn filecount.go som skal lese en tekst-fil, skrive ut totalt antall linjer og de fem “runes” (bokstaver/tegn) som forekommer mest i filen samt hvor mange ganger de telles i den medfølgende filen text.txt. 

Se "filecount.go" i mappen "Oblig2/src/oppgave2/". Programmet leser tekstfiler og skriver ut antall linjer og runes som forekommer mest. Vi har forsøkt med flere typer tekstfiler enn "text.txt" for å være sikre på at det fungerer. I bildet under kan du se hvilken format som blir skrevet ut når programmet kjører. Oppgave 2 bruker os.args[1] istedet for flag parsing slik som oppgave 1 bruker.

![Mac](https://github.com/Prestholm/TeamStovsuger/blob/master/Oblig2/src/Bilder/filecount.png)


## Oppgave 3. 

### A. Skriv et program med navn addup.go bestående av to go funksjoner og kommunisere med hverandre gjennom channels
Se "(filnavn).go" i mappen "Oblig2/src/oppgave3/".

!(https://github.com/Prestholm/TeamStovsuger/blob/master/Oblig2/src/Bilder/D1.png)

### B. Skriv to programmer (addtofile.go og sumfromfile.go) som kommuniserer med hverandre gjennom en fil.
Se "addtofile.go" og "sumfromfile.go" i mappen "Oblig2/src/oppgave3/".

!(https://github.com/Prestholm/TeamStovsuger/blob/master/Oblig2/src/Bilder/D2.png)

### C. Beskriv og implementer feilhåndtering på all I/O i oppgave a) og b).
Se "addup.go", "addtofile.go" og "sumfromfile.go" i mappen "Oblig2/src/oppgave3/".

Feilhåndteringen sørger for at tall 0 er den eneste input verdien man kan sette. Dersom man setter inn en verdi av type string eller en verdi under 0 vil man få opp feilmeldingen og programmet avsluttes. Dette vil man få opp både for nummer1 og nummer2. Denne feilhåndteringen representeres både i "addup.go" og "addtofile.go".

!(https://github.com/Prestholm/TeamStovsuger/blob/master/Oblig2/src/Bilder/D3.png)
!(https://github.com/Prestholm/TeamStovsuger/blob/master/Oblig2/src/Bilder/D4.png)

### D. Implementer håndtering av SIGINT i programmene i a) og b); programmene skal skrive ut en avslutningsmelding dersom de mottar SIGINT før de fullfører naturlig.
Se "addup.go" og "addtofile.go" i mappen "Oblig2/src/oppgave3/".

Ved å ta i bruk SIGINT har brukeren mulighet til å stoppe programmet, dette fungerer bra om man kjører programmets executable fordi man har da mulighet til å se resultatet, uten at programmet lukker seg før du har sett det.
Begge programmen stoppes ved bruk av tastene "ctrl+c". Og begge programmene sender ut den samme avslutningsmeldingen.

!(https://github.com/Prestholm/TeamStovsuger/blob/master/Oblig2/src/Bilder/D6.png)
!(https://github.com/Prestholm/TeamStovsuger/blob/master/Oblig2/src/Bilder/D5.png)

### E. Bygg filene som i oppgave 1 og 2 til kjørbare filer på ditt operativsystem og legg dem ved i /bin mappen.

Se i mappen "Oblig2/src/bin/".
