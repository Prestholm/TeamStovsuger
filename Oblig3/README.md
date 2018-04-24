
# Obligatorisk oppgave 3: 

Medlemmer: 
- Henning Van Ta.
- Victor Prestholm.
- Arian Nuland. 

## Oppgave 1. 

### A. Sett opp en webserver som lytter til port 8080.
Den er satt opp med if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
}


### Oppgave B. Når klienter aksesserer webserveren på path lik "/" skal klienten få tilbake strengen "Hello, client". Strengen skal vises i nettleseren.

Ved localhost:8080 får vi "Hello, client" strengen fram.



## Oppgave 2. 

### A. Presenter data på webserveren slik at den er leselig for mennesker når klienter aksesserer stiene /1, /2, /3, /4, /5.

Dataen er hentet fra 5 ulike apier, og er leselig for mennesker ved bruk av html.

### B. Samtlige stier i oppgave 2 skal rendres til klienter ved bruk av Go templates.

Ved bruk av html lager vi format slik at det kan rendres til klienter ved bruk av html-teplatesene.


## Oppgave 3. 

### A. Implementer et serverprogram i henhold til RFC 865. Serveren skal svare både på UDP og TCP.

Server programmet som er implementert svarer på RFC 865, med quoten: "Good, better, best. Never let it rest. 'Til your good is better and your better is best". Ved en visualisering av resultatet under så kan man se at serveren svarer på både TCP og UDP.

Ved å ta i bruk programmer som ligner på nmap(ncat) så får man muligheten til å scanne porten på TCP óg UDP. I en standard windows så får man bare mulighet til å scanne etter TCP porter ved bruk av telnet.

Kommandoer brukt for å få serveren til å responde:

Kommandoer
- Kommando i netcat(CMD) for tcp: netcat -z -v ip port
- Kommando i netcat(CMD) for udp: netcat -z -v -u ip port
- Kommando i CMD for tcp: telnet ip port 

Visualisert i Linux med innebygd nmap(ncat).

![Linux](https://github.com/Prestholm/TeamStovsuger/blob/master/Oblig3/Bilder/6ca6423f65a32074e60b61d4bfd0e9cc.png)

Visualisert i Windows med nedlastet nmap(ncat).

![Windows](https://github.com/Prestholm/TeamStovsuger/blob/master/Oblig3/Bilder/b0702e286232c1515c9a43f6947e5040.png)

(Lokal IPen er annerledes fordi jeg har byttet internett imellom bildene).

