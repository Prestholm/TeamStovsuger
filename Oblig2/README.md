
# Obligatorisk oppgave 2: 

Medlemmer: 
- Henning Van Ta.
- Victor Prestholm.
- Arian Nuland. 

## 1. 

### Oppgave A. Skriv et Golang program med navn fileinfo.go. 
Returnerer fil-info. Returnerer en error hvis det ikke er noen fil. Finner filen i oppgave1 @ fileinfo.go. 

### Oppgave B. Lag en kjørbart program (build) av fileinfo.go som kan utføres fra kommandolinje.
Oppgave b, bruker command-line flag parsing. Alt man trenger å gjøre er å builde og så kjøre den filen du vil ha info om ved "fileinfo -f [filnavn]", eller ved "go run fileinfo.go -f [filnavn]" men dette lager ikke en executable. Slik drar programmet info fram om filen, størrelse og hva slags type fil etc.
Oppgave 1b er gjort litt anderledes enn de andre oppgavene. Her bruker man Flags for å finne fil-navn. Det er ikke helt nødvendig å bruke flags her, oppgaven kan også løses ved å bruke vanlig os.args[1]. 
|![oppg1](https://github.com/Prestholm/TeamStovsuger/blob/master/Oblig2/src/Bilder/terminal.png)|

## 2. 

### Oppgave A. Skriv programmer med navn filecount.go som skal lese en tekst-fil, skrive ut totalt antall linjer og de fem “runes” (bokstaver/tegn) som forekommer mest i filen samt hvor mange ganger de telles i den medfølgende filen text.txt. 

Se "filecount.go" i mappen "Oblig2/src/oppgave2/". Programmet leser tekstfiler og skriver ut antall linjer og runes som forekommer mest. Vi har forsøkt med flere typer tekstfiler enn "text.txt" for å være sikre på at det fungerer. I bildet under kan du se hvilken format som blir skrevet ut når programmet kjører. Oppgave 2 bruker os.args[1] istedet for flag parsing slik som oppgave 1 bruker.

|![Mac](https://github.com/Prestholm/TeamStovsuger/blob/master/Oblig2/src/Bilder/filecount.png)|


## 3. 

### Oppgave A. Skriv et program med navn addup.go bestående av to go funksjoner og kommunisere med hverandre gjennom channels
Filen addup.go ligger i mappen oppgave 3. Filen har to funksjoner "main" og "b" som kommuniserer med hverandre og legger sammen to tall som blir angitt i terminalen.

### Oppgave B. Skriv to programmer (addtofile.go og sumfromfile.go) som kommuniserer med hverandre gjennom en fil.



### Oppgave C. Beskriv og implementer feilhåndtering på all I/O i oppgave a) og b).



### Oppgave D. Implementer håndtering av SIGINT i programmene i a) og b); programmene skal skrive ut en avslutningsmelding dersom de mottar SIGINT før de fullfører naturlig.



### Oppgave E. Bygg filene som i oppgave 1 og 2 til kjørbare filer på ditt operativsystem og legg dem ved i /bin mappen.


