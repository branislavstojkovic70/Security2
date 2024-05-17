# Izveštaj o Očvršćavanju Linux Servera

## 1. Priprema

### Preporučeno za pročitati
- The practical Linux hardening guide
- 40 Linux Server Hardening Security Tips

## 2. Uvod

Ovaj izveštaj se bavi postupkom pregleda i očvršćavanja (hardening) Linux sistema. Opisane su ključne korake koje treba preduzeti kako bi se poboljšala bezbednost sistema, uz primere uobičajenih problema koji se mogu pronaći na Linux serverima.

## 3. Zadaci

### 1. Postavljanje serverske mašine
Postavila sam serversku mašinu sa jednostavnim HTTP serverom koristeći LTS Ubuntu 22.04.

### 2. Pokretanje fiktivnog projekta – aplikacije
Pokrenula sam fiktivni projekat na podešenom serveru kao jednostavan HTTP server.

### 3. Dozvoljavanje pristupa aplikaciji iz mreže
Podesila sam firewall pravila da dozvoljavaju pristup aplikaciji iz mreže, omogućavajući dostupnost aplikacije fizičkoj host mašini.

### 4. Dozvoljavanje SSH saobraćaja
Omogućila sam SSH saobraćaj ka i od serverske mašine koristeći ključeve za autentifikaciju:

```bash
sudo apt install openssh-server -y
sudo systemctl enable ssh
sudo systemctl start ssh
ssh-keygen -t rsa -b 4096
ssh-copy-id korisnik@server_ip_address
```

### 5. Izvođenje pregleda sistema
Pokrenula sam pregled sistema koristeći Lynis alat:

```bash
sudo apt install lynis -y
sudo lynis audit system
```

Rezultati pregleda su sačuvani u odgovarajuće fajlove za kasniju analizu.

### 6. Kreiranje heš fajla za John-The-Ripper
Izvukla sam heš lozinke iz `/etc/shadow` fajla:

```bash
sudo grep 'korisnik' /etc/shadow > hash.txt
```

Pokrenula sam John-The-Ripper sa rockyou.txt wordlist-om:

```bash
sudo john --wordlist=/usr/share/wordlists/rockyou.txt hash.txt
```

## 4. Primer pregleda sistema

### 1. Operativni sistem
Informacije o operativnom sistemu:

```bash
lsb_release -a > /tmp/audit/lsb_release.txt
uname -a > /tmp/audit/uname.txt
```

### 2. Kernel
Verzija kernela:

```bash
uname -r > /tmp/audit/kernel_version.txt
uptime > /tmp/audit/uptime.txt
```

### 3. Upravljanje vremenom
Proverila sam da je sistem sinhronizovan sa NTP serverom:

```bash
cat /etc/timezone > /tmp/audit/timezone.txt
ps -edf | grep ntp > /tmp/audit/ntp_process.txt
ntpq -p -n > /tmp/audit/ntpq.txt
```

### 4. Instalirani paketi
Pregled instaliranih paketa:

```bash
dpkg -l > /tmp/audit/dpkg_list.txt
```

### 5. Logovanje
Konfiguracija logovanja:

```bash
ps -edf | grep syslog > /tmp/audit/syslog_process.txt
cat /etc/rsyslog.conf > /tmp/audit/rsyslog.conf.txt
```

### 6. Pregled mreže
Informacije o mrežnim interfejsima:

```bash
ifconfig -a > /tmp/audit/ifconfig.txt
route -n > /tmp/audit/route.txt
cat /etc/resolv.conf > /tmp/audit/resolv.conf.txt
cat /etc/hosts > /tmp/audit/hosts.txt
```

Firewall pravila:

```bash
iptables -L -v > /tmp/audit/iptables.txt
```

### 7. Pregled filesystem-a
Pregled mount-ovanih particija i permisija:

```bash
cat /etc/fstab > /tmp/audit/fstab.txt
find / -perm -4000 -ls > /tmp/audit/setuid_files.txt
find / -type f -perm -006 2>/dev/null | grep -v /proc > /tmp/audit/readable_files.txt
find / -type f -perm -002 2>/dev/null | grep -v /proc > /tmp/audit/writeable_files.txt
```

### 8. Pregled korisnika
Pregled passwd i shadow fajlova:

```bash
cat /etc/passwd > /tmp/audit/passwd.txt
cat /etc/shadow > /tmp/audit/shadow.txt
```

### 9. Konfiguracija sudo
Konfiguracija sudoers fajla:

```bash
sudo cat /etc/sudoers > /tmp/audit/sudoers.txt
```

### 10. Pregled servisa
Pregled pokrenutih servisa:

```bash
ps -edf > /tmp/audit/ps_ef.txt
lsof -i UDP -n -P > /tmp/audit/lsof_udp.txt
lsof -i TCP -n -P > /tmp/audit/lsof_tcp.txt
```

### 11. OpenSSH konfiguracija
Konfiguracija SSH:

```bash
cat /etc/ssh/sshd_config > /tmp/audit/sshd_config.txt
```

### 12. MySQL konfiguracija
Konfiguracija MySQL:

```bash
cat /etc/mysql/my.cnf > /tmp/audit/mysql_my.cnf.txt
mysql -u root -e "SELECT Host, User, Password FROM mysql.user;" > /tmp/audit/mysql_users.txt
```

### 13. Apache konfiguracija
Konfiguracija Apache-a:

```bash
cat /etc/apache2/apache2.conf > /tmp/audit/apache2.conf.txt
ls -l /var/www/html > /tmp/audit/www_permissions.txt
```

### 14. PHP konfiguracija
Konfiguracija PHP-a:

```bash
cat /etc/php/7.4/apache2/php.ini > /tmp/audit/php.ini.txt
```

### 15. Crontab
Pregled crontab zadataka:

```bash
crontab -u root -l > /tmp/audit/root_crontab.txt
ls -l /root/backup.sh > /tmp/audit/backup_permissions.txt
```

## 5. Zaključak

Na osnovu izvršenih pregleda i analize, sistem je podešen prema najboljim bezbednosnim praksama. Dokumentacija i screenshot-ovi su priloženi kao dokaz o preduzetim koracima za očvršćavanje sistema.
