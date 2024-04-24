
# SECURITY CODE REVIEW AND ANALYSIS

# Opis projekta 
Ideja projekta je implementacija Web aplikacije za zakazivanje vožnje nalik 
postojećem *Uber-u*. Svaki korisnik može da se registruje i da zakaže vožnju pri čemu može da bira kakvo vozilo želi, za koliko putnika itd. U sitemu postoje vozači, putnici i administrator. Neregistrovani korisnici mogu da se informišu o ceni vožnje, dok registrovani mogu da zakažu vožnju. Od tehnologija su korišćene Java + *Spring Boot* okruženje za Backend, i *Angular* za Frontend. Ovde će biti analiziran samo Backend deo aplikacije. 

# Članovi razvojnog tima
- Veljko Bubnjević SV 51-2020
- Nebojša Vuga SV 53-2020
- Jelena Miković SV 61-2020

# Opis pronađenih defekata

### 1. Login

```
    @PostMapping("/login")
    public ResponseEntity<TokenDTO> logIn(@Valid @RequestBody LoginDTO login) {
        try {

            TokenDTO token = new TokenDTO();
            SecurityUser userDetails = (SecurityUser) this.userService.findByUsername(login.getEmail());

            boolean isEmailConfirmed = this.passengerService.getIsEmailConfirmed(login.getEmail());


            String tokenValue = this.jwtTokenUtil.generateToken(userDetails);
            token.setToken(tokenValue);
            token.setRefreshToken(this.jwtTokenUtil.generateRefreshToken(userDetails));
            Authentication authentication =
                    this.authenticationManager.authenticate(
                            new UsernamePasswordAuthenticationToken(login.getEmail(),
                                    login.getPassword()));

            SecurityContextHolder.getContext().setAuthentication(authentication);

            return new ResponseEntity<>(token, HttpStatus.OK);
        } catch (Exception e) {
            return new ResponseEntity(new ErrorResponseMessage(
                    this.messageSource.getMessage("user.badCredentials", null, Locale.getDefault())
            ), HttpStatus.BAD_REQUEST);
        }

    }
```

U kodu za **@PostMapping("/login")** postoji nekoliko potencijalnih sigurnosnih nedostataka ili potencijalnih problema:

**Autentifikacija bez provere potvrđenog email-a**: Kod pristupačan omogućava autentifikaciju korisnika bez provere da li je njihov email potvrđen. To može dovesti do situacije gde korisnici sa neaktivnim ili nevalidnim email adresama mogu pristupiti aplikaciji. Konkretno problem se javlja jer flag isEmailConfirmed u kodu se nigde ne proverava.

**Lozinke u plain textu**: Lozinke se šalju u plain text formatu putem HTTP zahteva. Ovo može biti problematično jer osobe sa pristupom mreži mogu da uhvate ove zahteve i otkriju lozinke.

**Nema ograničenja broja pokušaja logovanja**: Nema ograničenja na broj pokušaja prijavljivanja. Ovo otvara mogućnost za *brute-force* napade gde napadači mogu pokušati automatski probati različite kombinacije korisničkih imena i lozinki dok ne pronađu validne kredencijale.

**Greška otkrivanja**: U slučaju greške pri autentifikaciji, specifično kada su kredencijali nevalidni, server vraća detaljnu poruku o grešci (**user.badCredentials**). Ovo otkriva suviše informacija o unutrašnjem stanju aplikacije i može biti iskorišćeno od strane napadača za pripremu napada.



### 2. Izmena lozinke

    @GetMapping("/{id}/resetPassword")
    public ResponseEntity<String> resetPassword(@PathVariable Long id) {

        User user = this.userService.findOne(id);
        if (user == null)
            return new ResponseEntity<>("User does not exist!", HttpStatus.NOT_FOUND);

        ResetPasswordToken resetPasswordToken = new ResetPasswordToken(id);
        this.userService.saveResetPasswordToken(resetPasswordToken);
        System.out.println("\n\n" + resetPasswordToken.toString() + "\n\n");
        System.out.println("\n\n" + resetPasswordToken.getCode() + "\n\n");
        return new ResponseEntity<>(resetPasswordToken.getCode(), HttpStatus.OK);
    }

**Lozinke u plain textu**: Slanje i stare i nove lozinke obavlja se kroz otvoren tekst bez heširanja što daje mogućnost napadaču da prepozna lozinke. Ono što je dobro je što za izmenu lozinke klasa *ChangePasswordDTO* podržava *regex* za proveru valjanosti formata lozinki.

**Odgovor na grešku**: Kada dođe do greške prilikom promjene lozinke, odgovor servera otkriva detalje o grešci, što može biti korisno za potencijalnog napadača. Bolje je koristiti generičke poruke o grešci bez detalja kako bi se smanjio rizik od otkrivanja informacija.

### 3. Dobavljanje informacije o korisniku na osnovu *email-a*

    @GetMapping("/email")
    public ResponseEntity<UserFullDTO> getUserByEmail(@RequestParam("email") String email) {
        return new ResponseEntity<>( new UserFullDTO(this.userService.findByEmail(email).get()), HttpStatus.OK);
    }

**Odgovor bez autorizacije:** Ova metoda omogućava pristup informacijama o korisniku samo na osnovu email adrese koja se prosleđuje kao parametar. Međutim, ne postoji provera autentifikacije ili autorizacije koja bi ograničila pristup samo autorizovanim korisnicima. To može dovesti do curenja osetljivih informacija, kao što su podaci o korisnicima, ako se metoda zloupotrebi.

**Otkrivanje informacija o korisnicima:** Vraćanje potpunog korisničkog objekta, uključujući sve informacije o korisniku, na osnovu samo *email* adrese može biti rizično. To može olakšati napadačima da prikupe informacije o korisnicima, što može biti korisno za potencijalne napade poput ribarenja (phishing) ili drugih oblika socijalnog inženjeringa.

**Nepotpuna validacija parametara:** Parametar email koji se prosleđuje putem zahteva nije validiran u smislu ispravnosti formata email adrese ili njegove autentičnosti. Ovo može dovesti do problema sa sigurnošću ili performansama, kao što su SQL injection ili XSS (*Cross-Site Scripting*) napadi.

### 4. Izmena podataka o putnicima

    @PutMapping(value = "/{id}")
    @PreAuthorize("hasAnyRole('ADMIN', 'PASSENGER')")
    public ResponseEntity<PassengerDTO> updatePassenger(@RequestBody @Valid PassengerDTO passengerDTO,
                                                        @PathVariable Long id,
                                                        @RequestHeader Map<String, String> headers){
        String role = this.userRequestValidation.getRoleFromToken(headers);
        if(role.equalsIgnoreCase("passenger")){
            boolean areIdsEqual = this.userRequestValidation.areIdsEqual(headers, id);
            if(!areIdsEqual) return new ResponseEntity("Passenger does not exist!", HttpStatus.NOT_FOUND);
        }
        if(this.passengerService.findOne(id).isEmpty())
            return new ResponseEntity("Passenger does not exist!", HttpStatus.NOT_FOUND);
        Passenger passenger = this.passengerService.findByEmail(passengerDTO.getEmail());
        if(passenger != null && !passenger.getId().toString().equals(id.toString())){
            return new ResponseEntity("Invalid data. For example Bad email format.", HttpStatus.BAD_REQUEST);
        }
        passenger = this.passengerService.getPassengerFromPassengerDTO(id, passengerDTO);
        passenger = this.passengerService.save(passenger);
        PassengerDTO updatedPassenger = new PassengerDTO(passenger);
        return new ResponseEntity<>(updatedPassenger, HttpStatus.OK);
    }

**Zastita od SQL injekcija:** Podatak iz putanje *id* se koristi 	za direktan upit u bazu podataka. Trebalo bi koristiti parametrizovane upite ili druge metode kako bi se izbegle SQL injekcije.

**CSRF:** Nije implementiran CSRF token, što bi značilo da bi napadač mogao da iskoristi ranjivost kako bi izvršio neželjene akcije na ime korisnika bez njihovog pristanka.

Provera identiteta putnika se vrši na osnovu podataka iz JWT tokena. Ovo može biti ranjivo na manipulaciju ili neispravne informacije iz zaglavlja zahteva. Potrebno je koristiti pouzdane mehanizme za autentifikaciju i autorizaciju kako bi se osiguralo da su informacije o identitetu korisnika pouzdane i integritetne.

# Preporuke za poboljšanje koda

1. **Implementacija CSRF zaštite**

   Trebalo bi uključiti CSRF zaštitu u Spring Boot konfiguraciju kako bi se aplikacija zaštitila od CSRF napada. 

1. **Heširanje i soljenje lozinki**

   Potrebno je da se lozinke primaju u hash formatu a ne u izvornom obliku. Pre čuvanja u bazu podataka one se hešuju ali je potrebno dodati *salt* pre procesa samog hešovanja kako bi čak i da dođe do kompromitovanja baze podataka, lozinka neće biti lako otkrivena.

1. **Odgovori na greške**

   Potrebno je osigurati da poruke o greškama ne otkrivaju previše detalja, kako ti detalji ne bi bili korisni napadačima. 

1. **Princip najmanjih privilegija**

   Korisnicima treba dati samo one privilegije koje su im neophodne za obavljanje njihovih zadataka. To bi trebalo da smanji površinu napada i posledice u slučaju kompromitovanja naloga. 

# Statička analiza koda upotrebom alata Qodana

Koristeći alat Qodana integrisan u *IntelliJ* IDE izvršena je statička analiza koda i dobijeni su sledeći rezultati:

<div style="text-align:center;">
    <img src="analiza.png" alt="Alt tekst" width="300">
</div>

U kontekstu bezbednosti konkretno je pronađeno 18 problema a oni se tiču korišćenja odgovarajućih biblioteka za  koje postoje određene ranjivosti. Ranjivosti su kategorisane prema CVE (*Common Vulnerabilities and Exposures*) identifikatorima, koji su standardizovani identifikatori za poznate sigurnosne probleme.

Jedna od ranjivosti je sledeća:

1. **CVE-2022-25857**: Ova ranjivost se odnosi na nekontrolisanu potrošnju resursa i ima visok stepen ozbiljnosti (*High severity*).
1. **CVE-2022-38751, CVE-2022-38752, CVE-2022-38749, CVE-2022-38750, CVE-2022-41854**: Ove ranjivosti se odnose na pisanje van granica (*out-of-bounds write*) i imaju srednji stepen ozbiljnosti (*Medium severity*).
1. **CVE-2022-1471**: Ova ranjivost se odnosi na deserijalizaciju nepoverljivih podataka i ima visok stepen ozbiljnosti (*High severity*).

Da bi se rešili ovakve ranjivosti, trebalo bi ažurirati novije verzije biblioteka kako bi prikazane ranjivosti bile uklonjene. 

Ostali problemi koji su detektovani tiču se potencijalnih sintaksnih i semantičkih grešaka.

# Vreme potrebno za analizu koda
Sat vremena za ručnu analizu i pola sata koristeći Qodan-u. 

