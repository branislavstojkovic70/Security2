# Opis projekta

#### Web aplikacija za upravljanje sertifikatima. Aplikacija omogućava registraciju i autentifikaciju korisnika, izdavanje, povlačenje i validaciju sertifikata, uz mogućnost upravljanja lozinkama i sertifikatima. Neophodno je omogućiti sigurnu komunikaciju i čuvanje podataka, uz implementaciju funkcionalnosti poput dvofaktorske autentifikacije i zaštite formi pomoću ReCAPTCHA. Rad se obavlja u timovima i mora biti dostupan na platformi za kontrolu verzija.

# Lista članova razvojnog tima

- Anja Petković SV22/2020
- Branislav Stojković SV64/2020
- Aleksandr Misutkin SV81/2021

# Opis pronadjenih defekata

## Kredencijali/Tajne u Samom Kodu (Hardcoded)

Jedan od značajnih bezbednosnih rizika u razvoju softvera je uključivanje osetljivih podataka direktno u izvorni kod ili konfiguracijske datoteke koje su deo aplikacije. Ova praksa, poznata kao "hardcoding", uključuje direktno unošenje kredencijala kao što su korisnička imena, lozinke, tajni ključevi i drugi osetljivi podaci unutar kodova koje aplikacija koristi za pristup resursima poput baza podataka, email servera, eksternih API-ja i drugih sličnih servisa.

### Primeri Osetljivih Podataka u Vašem Konfiguracijskom Fajlu

U konfiguracijskom fajlu, sledeći unosi sadrže osetljive podatke koji ne bi trebali biti hardcoded:

- **Kredencijali baze podataka**:
  ```properties
  spring.datasource.username=root
  spring.datasource.password=rootroot
  ```
- **OAuth2 konfiguracija**:
  ```properties
  spring.security.oauth2.client.registration.auth0.client-id=okCV1iX8I6mb9BqQZ6YnNF1hDcx3fv5n
  spring.security.oauth2.client.registration.auth0.client-secret=OMfcJyR-DSJRRdd-L0AQQnxNeohhNfIv3_bOUTZgfVGQjKLFlHlKjh6k58hJnHmL
  ```
- **Email server kredencijali**:
  ```properties
  spring.mail.username=uberapptim18@gmail.com
  spring.mail.password=gckrtmiwtawtnyww
  ```

### Imenovanje Paketa:

- **Prilagodjenje imena paketa**: Imena paketa treba da prate konvenciju malih slova sa povremenim podvlakama, što olakšava razumevanje strukture i smanjuje rizik od konflikata imena. Ovo je posebno važno u velikim projektima gde struktura paketa može biti kompleksna.

### Java Klase i Kontroleri:

- **Uklanjanje nepotrebnih uvoza**: Nepotrebnih uvoza dodaju nepotrebnu zavisnost i mogu usporiti kompilaciju. Redovno proveravanje i čišćenje uvoza može povećati čitljivost koda i olakšati održavanje.

- **Zamena injekcija polja sa konstruktorskim injekcijama**: Ovo poboljšava i olakšava testiranje codea i redukuje zavisnost komponenata na specifične okvire, čineći kod čistijim i lakšim za upravljanje.

- **Uklanjanje praznih naredbi i beskorisnih dodela**: Ovo pomaže u održavanju čistog i efikasnog koda, oslobađajući resurse i poboljšavajući performanse.

- **Zamena `System.out` sa loggerom**: Loggeri pružaju bolju kontrolu nad izlaznim porukama i omogućavaju lakše praćenje problema u produkciji.

- **Obrada mogućih `NullPointerExceptions`**: Provera null vrednosti pre njihove upotrebe može spriječiti česte izuzetke koji uzrokuju prekide u izvršavanju programa.

- **Refaktorisanje konstruktora sa previše parametara**: Smanjenje broja parametara u konstruktorima povećava čitljivost i olakšava upravljanje zavisnostima.

### Rukovanje Izuzecima:

- **Definisanje i bacanje specifičnih izuzetaka**: Korišćenje specifičnih izuzetaka umesto generičkih pomaže u boljem razumevanju i obradi grešaka u aplikaciji.

### Refaktorisanje i Jasnoća Koda:

- **Smanjenje kognitivne složenosti metoda**: Održavanje metoda jednostavnim i jasnim povećava razumljivost i smanjuje greške.
- **Korišćenje eksplicitnih zagrada za sprečavanje grešaka vezanih za else klauzulu**: Jasno definisanje blokova koda pomaže u izbegavanju grešaka u logici.
- **Zamena `instanceof` provera sa modernijom Java sintaksom**: Povećava sigurnost tipa i čitljivost koda.

### Test Slučajevi:

- **Osiguranje da test slučajevi imaju bar jednu tvrdnju**: Validacija efektivnosti testova i osiguranje da oni adekvatno testiraju kod.

## CURENJE INFORMACIJA (INFORMATION LEAK)

### Cuvanje adminove lozinke

- **Konzolna informacija o lozinki**: Prilikom pokretanja rešenja, inicijalizuje se admin i njegovi kredencijali. Kredencijali se ispisuju u konzolnoj liniji što može biti štetno iz razloga:

- **Logovanje kredencijala**: Često se sesije konzole automatski loguju za potrebe praćenja i debagovanja. Ako su kredencijali ispisani u konzoli, oni mogu biti nehotično sačuvani u log fajlovima. Ovi logovi mogu biti dostupni drugim sistemskim administratorima ili mogu biti kompromitovani tokom sigurnosnih incidenata.
- **Izlaganje osetljivih informacija**: Kada se kredencijali ispisuju direktno u konzoli, oni postaju vidljivi svima koji imaju pristup toj konzoli ili terminalu. To uključuje i one koji mogu fizički videti ekran ili pristupiti sesiji konzole preko mreže.
- **Povećan rizik od unutrašnjih pretnji**: Ispisivanje kredencijala povećava rizik od unutrašnjih pretnji, jer zaposleni ili ugovorni radnici mogu zloupotrebiti te informacije. Čak i ako ne dođe do zloupotrebe, sama dostupnost tih informacija povećava rizik.

### NEDOSTATAK MEHANIZMA ZA ČUVANJE LOGOVA

Čuvanje logova bez adekvatnog filtriranja i zaštite može dovesti do brojnih sigurnosnih i operativnih problema. Evo nekoliko ključnih nedostataka ove prakse:

1. **Izlaganje osetljivih podataka**: Logovanje svih akcija bez filtera može uključivati nehotično snimanje osetljivih informacija poput lozinki, tokena za autentifikaciju i ličnih podataka. Ovi podaci mogu biti izloženi ako napadači dobiju pristup log fajlovima.

2. **Kršenje zakonskih regulativa**: U mnogim jurisdikcijama, postoje strogi propisi koji zahtevaju zaštitu ličnih podataka i drugih osetljivih informacija. Neadekvatno rukovanje ovim podacima u logovima može dovesti do pravnih problema, uključujući velike novčane kazne.

3. **Povećan rizik od unutrašnjih pretnji**: Logovi koji sadrže osetljive informacije mogu postati meta unutrašnjih pretnji. Zaposleni ili ugovornici sa pristupom ovim logovima mogu zloupotrebiti osetljive informacije.

4. **Upravljanje i skladištenje**: Logovi koji sadrže veliku količinu neprofiltriranih podataka mogu brzo rasti, što otežava upravljanje i povećava troškove skladištenja. Takođe, obrada i analiza ovakvih logova može biti otežana zbog prevelike količine nepotrebnih informacija.

# Preporučene prakse

### Preporučene Prakse za Upravljanje Kredencijalima

- **Korišćenje okruženjskih promenljivih**: Osetljivi podaci bi trebali biti postavljeni kao promenljive okruženja na serveru. Ovo omogućava da aplikacija čita ove podatke iz okruženja umesto da budu direktno u kodu.
- **Korišćenje servisa za upravljanje tajnama**: Alati kao što su HashiCorp Vault, AWS Secrets Manager, ili Azure Key Vault omogućavaju sigurno skladištenje i upravljanje pristupom osetljivim konfiguracijskim podacima.

- **Enkripcija**: Ako morate da čuvate osetljive podatke unutar projekta, koristite snažnu enkripciju za zaštitu tih podataka.

Primena ovih praksi ne samo da povećava bezbednost vaših aplikacija već i olakšava upravljanje i skaliranje u različitim okruženjima i fazama razvoja projekta.

### Preporučene Prakse za Čuvanje Logova

Da bi se poboljšala sigurnost i efikasnost logovanja, preporučuje se sledeće:

1. **Implementacija filtera za logovanje**: Postavite filtere koji će isključiti osetljive podatke iz logova. To uključuje automatsko maskiranje ili uklanjanje informacija kao što su lozinke, tokeni, brojevi kreditnih kartica i druge lične informacije.

2. **Korišćenje centralizovanih sistema za logovanje**: Implementirajte centralizovani sistem za logovanje koji podržava šifrovanje, sigurnu autentifikaciju i kontrolu pristupa. Ovo pomaže u smanjenju rizika od neautorizovanog pristupa logovima.

3. **Periodično čišćenje logova**: Automatizujte procese čišćenja logova kako bi se izbeglo prekomerno akumuliranje podataka. Periodično čišćenje pomaže u smanjenju količine skladištenih podataka i troškova.

4. **Revizija i monitoring pristupa logovima**: Postavite stroge politike za reviziju i monitoring pristupa log fajlovima. To uključuje praćenje i snimanje svih pristupa logovima, kao i alarme za sumnjive aktivnosti.

5. **Šifrovanje logova**: Osigurajte da su svi logovi šifrovani kako tokom prenosa, tako i tokom skladištenja. Ovo pomaže u zaštiti logova od neautorizovanog pristupa i čitanja.

6. **Obuka zaposlenih**: Obrazujte zaposlene o važnosti sigurnosti logova i potencijalnim rizicima povezanim sa neadekvatnim rukovanjem log fajlovima.

# Vreme koje je proveo svaki od članova tima (reviewer) pregledajući kod i broj defekata koji je identifikovao/la

- Branislav Stojkovic SV64/2020 : približno 1h za pokretanje alata za statičku reviziju koda SonarQube i oko 13h za ručno pregledanje koda
