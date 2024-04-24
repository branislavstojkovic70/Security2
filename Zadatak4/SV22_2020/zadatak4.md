# Opis projekta
#### Klijentska aplikacija za sistem zakazivanja vožnje. Projekat je implementiran u Angular framework-u, koristeći TypeScript programski jezik. Specifikacijom je predviđeno poručivanje vožnje, ocenjivanje vožnje i vozača, izbor vozila, dodavanje putnika, vozača i vozila.

# Lista članova razvojnog tima
- Anja Petković SV22/2020
- Branislav Stojković SV64/2020
- Andrijin kristina SV29/2020

# Pronađeni defekti

## KREDENCIJALI/TAJNE U SAMOM KODU (HARDCODED) 
 - Što se tiče kredencijala u kodu, nema ih direktno jer se svi nalaze u environments/environment.ts
 - Kredencijala ima indirektno kao što je ovaj primer za ključ u sessionStorage

        const USER_KEY = 'auth-user';

        @Injectable({
        providedIn: 'root'
        })
        export class StorageService {
        constructor(private http: HttpClient) {}
 - Još jedan od primera: 

        @Injectable({
        providedIn: 'root'
        })
        export class HttpRequestInterceptor implements HttpInterceptor {
        constructor(private storageService: StorageService) {}
        intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
            const jwt = localStorage.getItem('jwt');
 - Ovo bi napadaču omogućilo da preko sessionStorage item-a ukrade token za autentifikaciju

 ## CURENJE INFORMACIJA (INFORMATION LEAK) 
 - Curenje informacija se može desiti na mnogo mesta u ovom kodu:
    ### Nebezbedno čuvanje tokena u localStorage ili sessioStorage

        public saveUser(user: any): void {
            window.sessionStorage.removeItem(USER_KEY);
            window.sessionStorage.setItem(USER_KEY, JSON.stringify(user));
        }
     - Ovo može omogućiti napadačima da ukradu token i koriste ga za dobijanje neovlašćenog pristupa
    
    ### Loša implementacija sigurnosnih zaglavlja u http zahtevima

        intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
        const jwt = localStorage.getItem('jwt');
            // console.log(jwt);
            //console.log(this.storageService.getUser().id);
            const jwtSpring = this.storageService.getUser().accessToken;
            //console.log(jwtSpring);
            if (jwtSpring) {
            const cloned = req.clone({
                setHeaders: {
                Authorization: `Bearer ${jwtSpring}`
                }
            });
            return next.handle(cloned);
            } else {
            return next.handle(req);
            }
        }
     - Aplikacija može postati ranjiva na različite vrste napada kao što su XSS (Cross-Site Scripting) i CSRF (Cross-Site Request Forgery) koji mogu dovesti do curenja informacija

    ### Neadekvatno filtritriranje unetih podataka

        <input type="text" name="departure" formControlName="departure" placeholder="Departure"/>

    - Ovo može dovesti do XSS ili injection upada jer input forma nije adekvatno validirana ili escape-ovana

    ### Neimplementirana CSP pravila
    - Ovo takođe može dovesti do XSS napada

## NEDOSTATAK MEHANIZMA ZA HEŠOVANJE LOZINKI

 - Mehanizam za hash lozinke nije uopšte urađen na klijentskoj strani, već se lozinka na serversku stranu prenosi kao *plain text*

        let loginVal = {
        username: this.loginForm.value.email?.toString(),
        password: this.loginForm.value.password?.toString(),
        };


        if (this.loginForm.valid) {
        this.valid = true;
        let email = loginVal?.username;
        let password = loginVal?.password;

        // @ts-ignore
        this.authService.login(email, password).subscribe({
        next: data => {
            this.storageService.saveUser(data);

            console.log("successful login");

    
        export class AuthService {

        constructor(private http: HttpClient) { }

        login(email: string, password: string): Observable<any> {
            return this.http.post(
            AUTH_API + 'login',
            {
                email,
                password,
            },
            httpOptions
            )
            ;
        }
## NEDOSTAJUĆA CSRF ZAŠTITA 
 - Ni jedna forma nije zaštićena CSRF tokenom, niti je CRSF igde primenjen

## IZLISTAVANJE DIREKTORIJUMA 
 - Nije urađena kontrola pristupa za direktorijume niti je implementiran neki servis koji bi proveravao da li je korisnik ulogovan i na osnovu tokena vodio računa o pristupu resursima.

## KRIPTOGRAFSKI PROBLEM
 - Lozinka nije hash-ovana kao što je već navedeno, niti je bilo koji kriptografski algoritam igde primenjen

## ZAOBILAŽENJE POTPISA i ZAOBILAŽENJE AUTORIZACIJE 
 - Zaobilaženje potpisa i zaobilaženje autorizacije ipak nije moguće jer je implementiran interceptor za slanje zahteva.

## REMOTE CODE EXECUTION 
 - U kodu nije pronađeno mesto gde bi se mogao vršiti napad REMOTE CODE EXECUTION 

# Sažetak preporuka
#### 1. Dinamičko dobavljanje ključeva za local i session storage
#### 2. Upotreba memorije ili third party biblioteke za čuvanje JWT tokena, a ne localStorage ili sessionStorage
#### 3. Implementirati zahteve sa CSP zaglavljima
#### 4. Vršiti adekvatnu validaciju i escaping polja za unos
#### 5. Obavezno implementirati mehanizam za hash lozinke i nikada je ne slati u plain text-u
#### 6. Implementirati CRSF zaštitu za forme
#### 7. Obavezno implementirati mehanizam za kontrolu pristupa direktorijumima, kao i autorizaciju za putanje koje posećuje
#### 8. Razmotriti upotrebu kriptografskih algoritama na mestima gde je neophodna zaštita podataka

# Rezultati SonarQube alata za statičku analizu Angular projekta:
 - Security (Sigurnost): 0 Open issues
 - Reliability (Pouzdanost): 122 Open issues
 - Maintainability (Održivost): 240 Open issues
 #### Preporuke koje su bile zastupljene u izveštaju
 1. Remove Commented Out Code

        <!--    <input type="file" class="file" name="avatar" formControlName="profilePicture" placeholder="Picture" [readonly]="true">-->
 2. Add Descriptions to Tables

        <table mat-table [dataSource]="dataSource" style="background-color: #EBE8D7">
 3. Associate Form Labels with Controls

        <div class="col">
            <label>ID: </label>
            <input type="text" [formControl]="form.controls.id" class="myInput2" readonly>
        </div>
 4. Remove Unused Imports

        import {Component, NgModule, ViewChild} from '@angular/core';
 5. Fix Unexpected Missing Generic Font Family

        h2{
            text-align: left;
            padding-left: 10px;
            font-family: Crete round;
        }

 6. Address Deprecated Features

        private getDrivers(request: { page?: string; size?: string; }) {
            this.driverService.getAll(request)
            .subscribe(data => {
                this.drivers = data['results'];
                this.totalElements = data['totalCount'];
                }
                , error => {
                console.log(error.error.message);
                }
            );
        }

    - '(next?: ((value: { totalCount: number; results: Driver[]; }) => void) | null | undefined, error?: ((error: any) => void) | null | undefined, complete?: (() => void) | null | undefined): Subscription' is deprecated.
 9. Implement Proper Error Handling
        
        error: (error) => {
          reject(error);
        },
 10. Address Accessibility Issues

    <a>
        <img src="../../assets/chart.png">
        <p>Stats</p>
    </a>

    - Add an "alt" attribute to this image.
 11. Correct CSS Errors

            margin-left: 0.7vh;
            margin-right: 0.7vh;
            color: #485162;
            margin: 3rem;

        - Unexpected shorthand "margin" after "margin-right"
        
 #### Severity levels
 1. High: 23
 2. Medium: 184
 3. Low: 155

# Vreme koje je proveo svaki od članova tima (reviewer) pregledajući kod i broj defekata koji je identifikovao/la
 - Anja Petković SV22/2020 : približno 1.5h za pokretanje alata za statičku reviziju koda i oko 3h za ručno pregledanje koda