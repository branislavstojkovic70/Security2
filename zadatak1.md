# Zadatak 2 - logovi

##  Log datoteke moraju pružiti informacije potrebne za razrešavanje problema;
#### Kada dođe do problema u sistemu, ukoliko je sistem dobro implementiran i ima log datoteke u koje se upisuju informacije pravilno, moguće je otkriti uzrok problema upravo iz ovih fajlova. Log fajlovi nam mogu sugerisati na grešku koja se dogodila u sistemu, izuzetak do kojeg je došlo ili maliciozne radnje korisnika koje su izvršene nad sistemom. Takođe se neke vrtse podataka iz log datoteka mogu vizualizovati i otkriti dodatne informacije o radu sistema ili akcijama korisnika. Svedoci smo sve češće upotrebe AI alata, te nam mogu i oni biti korisni prilikom identifikovanja problema u sistemu
### Koraci pri analizi logova: 
#### Pošto je više parametara moglo da doprinese grešci, prvi korak je utvrđivanje da li je greška u infrastrukturi, greška u praćenju ili greška u transakciji izazvala probleme u sistemu
#### Drugi korak bi bio analiza logova na granularnom nivou. Naime, ako smo otkrili koja je vrsta greške, sada treba dalje analizirati šta je uzrok greške i na kom mestu se ona u sistemu dogodila. Na primer, pretpostavimo da se veb lokacija pokvari. U tom slučaju, od vitalnog je značaja da se odmah utvrdi da li je razlog server aplikacija, server baze podataka ili problem sa korišćenjem CPU-a, memorije ili diska da biste precizno došli do osnovnog uzroka i pravilno identifikovali problem..
#### Treći korak bi eventualno bio da nakon otkrivanja greške napravimo strategiju za rešavanje problema i izračunamo cenu poravke.
### Struktura loga:
#### Najčešći sadržaj log zapisa: 
 - **Datum i vreme**
 - **Nivo ozbiljnosti (severity level)**
 - **ID procesa ili thread-a**
 - **Izvor događaja**
 - **Poruka**
 - **Kod greške ili izuzetka (Stack trace)**
 - **Korisnički identifikator**
 - **Dodatni kontekst**
### Najčešći formati log fajla: 
 - *Plain Text (Čist tekst)* 
 - *CSV (Comma-Separated Values)*
 - *JSON (JavaScript Object Notation)*
 - *XML (eXtensible Markup Language)*
 - *Syslog*
 - *Log4j*

## Svi događaji za koje su akteri bitni moraju biti zapisani, sa dovoljno informacija kako akteri ne bi mogli da poriču odgovornost (non-repudiation). Potrebno je obezbediti lako izdvajanje tih događaja;
### Neporecivost se može u kontekstu log fajlova posmatrati sa dva aspekta: 
#### **Neporecivost fajla**
Odnosi se na to da fajl sam po sebi nije menjan, tako da se originalni neobrađeni fajl evidencije može predstaviti kao dokaz, bez pitanja autentičnosti, na sudu
#### **Neporecivost zapisa u fajlu**
Može se sprovesti praćenjem sledećih bitnih koraka:
1. **Osiguranje integriteta i autentičnosti log zapisa** - Obežbeđuje se upotrebom digitalnih potpisa, prilikom potpisivanja se koristi privatni ključ, a provera se vrši odgovarajućim javnim ključem. Ovo nam dokazuje da log zapis u međuvremenu nije promenjen od strane trećeg lica. Integritet se dodatno može proveravati pomoću checksum-a ili hash vrednosti, kako ne bi došlo do neautorizovanih izmena.
1. **Enkripcija logova** - Šifrovanje će zaštititi poverljive informacije iz logova. Ukoliko dođe do toga da su logovi kompromitovani, bez odgovarajućeg algoritma za dekripciju, napadač ne može doći do njihovog sadržaja.
1. **Čuvanje logova na centralizovanom i sigurnom log serveru** - Ovaj postupak čuvanja sprečava maliciozne aktere da vrše manipulaciju logova i pristupaju log fajlovima ukoliko za to nemaju prava.
1. **Automatizacija prikupljanja i prenosa logova** - Ovo nas štiti od ljudskih grešaka i namernih manipulacija, ali i obezbeđuje sistem ukoliko dođe do otkaza. Osim ovoga, neophodno je sprovoditi i politike koje određuju koliko dugo logovi treba da se čuvaju, gde i kako.
1. **Korišćenje skladišta podataka koje se ne može menjati** - Ovo je dobra stvar jer kada se log jednom kada upiše u takvo skladište, on ne može biti izmenjen ni obrisan. Čak i pored ovakvih skladišta, pristup logovima treba omogućiti samo onim korisnicima kojima je to neophodno radi monitoringa ili verifikacije akcija.

## Stavke log datoteke ne smeju sadržati osetljive podatke
### Šta su osetljivi podaci: 
- **Identifikacioni podaci** (*puna imena, adrese, adresa e-pošte, broj vozačke dozvole i broj telefona*)
- **Finansijski podaci** (*podaci o kreditnoj kartici i drugi finansijski podaci*)
- **Podaci o zdravstvenoj zaštiti** (*istorija bolesti i kartoni*)
- **Lozinke** 
- **IP adresa** 
### Opasnosti evidentiranja osetljivih podataka
Prema zakonima o privatnosti kao što su GDPR u EU i CCPA u Kaliforniji korisnici imaju sledeća prava:
- Zatražite informacije o tome koji podaci o njima postoje
- Dobijte informacije o tome zašto se njihovi podaci čuvaju
- Zahtevajte brisanje ličnih podataka
Ispunjavanje bilo kog od ovih zahteva postaje izuzetno teško ako imate korisničke podatke duplirane po sistemima i raširene po vašim evidencijama i deponijama i rezervnim kopijama baze podataka.

Istorijski gledano, evidencije su često meta kršenja podataka ili izvor slučajnog curenja podataka. Čuvanje osetljivih podataka iz vaših evidencija je jednostavan način za rešavanje ovog problema. Napadi će se desiti, ali držanjem osetljivih podataka van evidencije, značajno smanjujete vrednost svih podataka koji budu ugroženi.

### Najbolje prakse za čuvanje osetljivih podataka iz evidencije
1. **Izolujte osetljive podatke**
1. **Zapišite tokene, a ne vrednosti**
1. **Šifrujte tokom tranzita i u mirovanju**
1. **Držite lične podatke van URL-ova**
1. **Redigujte i maskirajte podatke**
1. **Upravljajte podacima**

### Najbolje prakse za prevenciju grešaka
1. **Revidirajte kod često**
1. **Zapisujte isključivo strukturirane logove**
1. **Postavite automatska obaveštenja ukoliko dođe do opasnosti**

## Mehanizam za logovanje morabiti pouzdan, mora obezbediti dostupnost i integritet log datoteka;
### **Pouzdanost**
 - Sistem za logovaje u svakom trenutku mora biti sposoban da zabeleži svaki relevantan događaj u sistemu u odgovarajućem formatu, bez obzira na opterećenost i broj događaja koji se trenutno vrše u sistemu.
  - Sistem za logovanje mora biti otporan na greške, prekide u radu, ali i veliki broj zahteva za upisivanje u fajl.
  - Pouzdan sistem logovanja garantuje da će sve bitne informacije biti zabeležene u standardnom formatu i sa tačnim podacima.

### **Dostupnost**
 - Log datoteke moraju biti dostupne za pregled i analizu u svakom trenutku
 - Sistem za logovaje mora biti dizajniran tako da podržava visoku dostupnost, uključujući redundantnost i mehanizme za oporavak podataka u slučaju kvara. Kao što su replikacije podataka i sinhronizacija istih.
 - Dostupnost osigurava da se logovi mogu koristiti za brzo reagovanje na incidente ili probleme u realnom vremenu, kao i za dugoročne analize.

### **Integritet**
 - Integritet log datoteka je od suštinske važnosti za očuvanje njihove pouzdanosti kao dokaza u sigurnosnim istragama i analizama.
 - Sistem za logovaje mora biti dizajniran tako da uključuje mehanizme koji sprečavaju neautorizovanu izmenu, brisanje ili na bilo koji način kompromitovanje log datoteka.
 - Ovo obuhvata mehanizme digitalnih potpisa, enkripcije, i druge metoda zaštite podataka koji osiguravaju verifikaciju autentičnosti logova.

## Stavke log datoteke moraju precizno iskazati vreme nastanka;
Ova osobina log sistema je veoma važna za detekciju početka problema ili u distribuiranim sistemima, gde ne postoji globalni časovnik.
 - **Precizne vremenske oznake** omogućavaju da se hronološki prate zapisi. Ovo je ključno za kreiranje retrospektive događaja koji dovode do problema. Takođe nam je korisno u situacijama kada je potrebno analizirati uzroke problema ili sigurnosnih incidenata.
 - U okviru **distribuiranih sistema**, u kojima ne postoji globalni časovnik, a događaji se mogu dešavati na različitim mestima, od neizmerne je važnosti sinhronizacija log zapisa kako bismo utvrdili koja je uzročno-posledična veza između dva takva događaja.
 - Vremenske oznake su ključne i za **uspostavljanje odgovornosti** akcija korisnika unutar sistema. Tokom praćenja korisničkih aktivnosti ili forenžičke analize vremenske oznake nam omogućavaju povezivanje određenog korisnika sa tačnim događajem zapisanim u sistemu.
 - **Industrijski standardi i zakonske regulative** zahtevaju precizno logovanje vremena zbog *sigurnosti, privatnosti i revizije*. 

## Mehanizam za logovanje mora stremiti ka tome da su logovi uredni, da je “pretrpanost” minimalizovana;
Kako bismo minimizovali pretrpanost logova i izbegli neurednost neophodno je da primenimo sledeće strategije prilikom zapisivanja logova:
 1. Postaviti **nivo evidencije loga** (DEBUG, INFO, WARNING, ERROR, CRITICAL) radi kontrole opširnosti evidencije.
 1. Implementirati **rotaciju logova** radi automatskog arhiviranja starih datoteke evidencije i kreiranja novih nakon što evidencije dostignu određenu veličinu ili starost. Ovo nam je korisno za izbegavanje predugačkih log datoteka.
 1. Potrebno je često vršiti **uzorkovanje log fajla**. U velikim sistemima sa ogromnim protokom informacija nije realno čitati svaki log zapis, ali s vremena na vreme je potrebno uzorkovati neku kolčinu zapisa radi analize rada sistema.
 1. Kao što smo i do sada navodili, neophodno je log zapise u sistem unositi strukturirano i u određenom formatu. Radi lakše pretrage i filtriranja evidencije na osnovu specifičnih polja.
 1. Takođe, evidencija log zapisa se mora centralizovati. **Centralizovani sistem logovanja** objedinjuje evidenciju iz više izvora. Ovakvi sistemi često dolaze sa funkcijama za filtriranje, pretraživanje i upozorenja, a ovo nam omogućava da se fokusiramo samo na najvažnije informacije.
 1. Neophodno je implementirati mehanizme **filtriranja log fajla**, te pomoću njih isključiti pretragu zapisa koji nisu korisni ili relevantni iz skladišta. Filtriranje se može vršiti na izvornom nivou ili na nivou agregatora. 
 1. Nekada nije dovoljno samo voditi monitoring fajlova, potrebno je implementirati **mehanizme upozorenja i metrike** adekvatne za to.
 1. Uvek je potrebno razmotriti dužinu čuvanja log zapisa. Ukoliko se ovakvi fajlovi skladište najbolje prakse su servisi za **skladištenje** koji podržavaju automatsku **kompresiju** i upravljanje životnim ciklusom.
 1. Još jedna od dobrih praksi je i edukacija ostalih developera o struj+krturi log zapisa, jer se time može na početku izbeći pretrpanost fajlova.

# Zadatak 3 - Izveštaj o Bezbednosnoj Implementaciji i Poboljšanjima

Analiza i ocena bezbednosnih aspekata implementiranih u projektu za upravljanje digitalnim sertifikatima. Projekat omogućava korisnicima da registruju naloge, autentifikuju se, upravljaju digitalnim sertifikatima, i sprovode bezbednosne operacije kao što su izdavanje, povlačenje, i validacija sertifikata. Bezbednost informacija i zaštita privatnosti korisnika su od suštinske važnosti, zbog čega se posebna pažnja posvećuje implementaciji robustnih bezbednosnih kontrola.

## Korišćeni Algoritmi i Bezbednosni Aspekti

### **Korišćen heš algoritam **
    Projekat koristi za autorizaciju autora koristili smo autorizator **Auth0** koji implementira autentikaciju **Auth - N**. Za heširanje koristi **bcryptJS biblioteku** kroz Auth0 za hešovanje lozinki. BcryptJS koristi **bcrypt algoritam** za heširanje lozinki. Bcrypt algoritam se široko koristi za heširanje lozinki zbog svojih robustnih sigurnosnih karakteristika. Implementira i tehniku salting, što sprečava efikasno napade korišćenjem prethodno izračinatih heš tablica, poznatih kao dugine tabele (rainbow tables). Dodatno, bcrypt omogućava podešavanje broja prolaza kroz heširanje, što znatno otežava brute-force napade, jer svako povećanje broja prolaza eksponencijalno povećava vreme potrebno za izračunavanje heša. U projektu algoritam prolazi kroz heširanje 10 puta. Bcrypt algoritam još uvek nema otkrivene ranjivosti.

#### **Poboljšanje heš algoritma**
    Iako bcrypt nema otkrivenih ranjivosti, postoji noviji algoritam **Argon2** koji je bio pobednk na Password Hashing Competition 2015. godine. Dizajniran je da bude otporan na različite vrste napada, uključujući napade pomoću specijalizovanog hardvera kao što su ASIC (Application-Specific Integrated Circuit) i FPGA (Field-Programmable Gate Array) uređaji. Postoje tri varijante Argon2: Argon2d, Argon2i, i Argon2id, gde svaka služi različitim sigurnosnim potrebama.

####**Argon2d**
    **Argon2d** je optimizovan za aplikacije koje zahtevaju najveću otpornost na napade pomoću GPU-a (Grafičkih Procesorskih Jedinica). Koristi pristup zasnovan na pristupu podacima koji je zavistan od lozinke, što ga čini veoma otpornim na napade pomoću specijalizovanih uređaja, ali može biti podložniji napadima preko bočnih kanala kada se izvršava na uređajima gde napadač može da posmatra pristup memoriji.

####**Argon2i**
    **Argon2i** je dizajniran za aplikacije koje zahtevaju otpornost na napade preko bočnih kanala, koristeći pristup pristupu podacima koji nije zavistan od lozinke. Ovo ga čini idealnim za upotrebu u situacijama gde je integritet memorije potencijalno kompromitovan, kao što su sistemi sa više korisnika ili virtualizovani okruženja, nudeći bolju bezbednost u takvim kontekstima.

####**Argon2id**
    **Argon2id** kombinuje pristupe Argon2i i Argon2d, nudeći balans između otpornosti na napade preko bočnih kanala i  napade koji koriste specijalizovane uređaje. Ova hibridna verzija pruža dobru otpornost u širokom spektru aplikacija, čineći ga fleksibilnim izborom za opštu upotrebu u različitim scenarijima heširanja lozinki.

### **Korišćenje tehnika za logove**
    Projekat koristi tehnike aspekto-orijentisanog programiranja za praćenje logova. Kada se izvrši bilo koja akcija na serveru funckija je upisivala logove u tekstualni fajl koji se čuvao na disku. Log je sadržao vreme izvršavanja akcije (timestamp), vrstu akcije, kao i tekst o samoj akciji.

#### **Unapredjenje tehnike za upisivanje logova**
    Tehnika koja se primenjuje u projektu je krajnje jednostavna i može se poboljšati. Da bi se logovi koristili efektivno, ključno je usvojiti pristupe koji omogućavaju jasnoću, preciznost i pouzdanost zabeleženih informacija, dok istovremeno minimizuju pretrpanost. Idealna praksa podrazumeva selektivno logovanje, gde se koristi hijerarhija nivoa važnosti (DEBUG, INFO, WARNING, ERROR, CRITICAL) za filtriranje informacija na osnovu njihove važnosti i relevantnosti, čime postižemo očuvanje čistoće log fajlova. Rotacija logova, uzorkovanje, i centralizacija takođe igraju ulogu u održavanju urednosti logova, omogućavajući sistematično arhiviranje starih zapisa i olakšavajući analizu podataka kroz centralizovane platforme za logovanje. Pored toga, primena strukturiranih logova (npr. u JSON ili XML formatu) značajno poboljšava čitljivost i olakšava automatizovanu obradu i analizu logova. Korišćenje naprednih tehnika, poput digitalnog potpisivanja i enkripcije logova, dodatno osigurava integritet i pouzdanost zabeleženih informacija, dok prakse kao što su filtriranje i upravljanje metrikama omogućavaju fokusiranje na informacije koje su bitne.

### **Korišćenje reCAPTCHA**
    **Googleova reCAPTCHA** predstavlja široko prihvaćen mehanizam za zaštitu web aplikacija od zloupotrebe i automatskih napada, kao što su spam i botovi koji automatski popunjavaju forme. Koristeći napredne tehnike analize ponašanja korisnika i izazove zasnovane na slikama, reCAPTCHA uspešno razlikuje ljude od mašinskih agenata, čime doprinosi sigurnosti i integritetu web sajtova. Osim toga, reCAPTCHA v3 unapređuje korisničko iskustvo eliminisanjem potrebe za direktnim interakcijama korisnika, kao što su selektovanje slika, pružajući nevidljivu verifikaciju koja kontinuirano ocenjuje rizik koristeći model ocenjivanja za detekciju potencijalno zlonamernih aktivnosti.

#### **Unapredjenje privatnosti**
    Kao alternativa, projekt može razmotriti implementaciju **hCAPTCHA**, koji služi sličnoj svrsi kao Googleova reCAPTCHA ali se fokusira na privatnost korisnika i transparentnost. hCAPTCHA ne samo da pruža efikasnu zaštitu protiv botova i automatizovanih napada, već takođe nudi kompenzaciju vlasnicima web sajtova za njihovu upotrebu putem sistema koji deli prihode od analize i obrade podataka. Osim toga, hCAPTCHA je dizajniran da bude jednostavan za integraciju i konfiguraciju, čineći ga privlačnim rešenjem za organizacije koje žele da poboljšaju svoje bezbednosne postavke bez kompromitovanja korisničkog iskustva ili privatno
