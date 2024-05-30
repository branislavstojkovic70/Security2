import shutil

from flask import Flask, request, render_template_string
import os
import tarfile

app = Flask(__name__)
UPLOAD_FOLDER = 'uploads'
os.makedirs(UPLOAD_FOLDER, exist_ok=True)


@app.route('/')
def index():
    return render_template_string('''
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>Upload File</title>
            <style>
                body { font-family: Arial, sans-serif; margin: 50px; }
                form { display: flex; flex-direction: column; width: 300px; }
                input { margin-bottom: 10px; }
            </style>
        </head>
        <body>
            <h1>Upload a Tar File</h1>
            <form action="/upload" method="post" enctype="multipart/form-data">
                <input type="file" name="file" required>
                <input type="submit" value="Upload">
            </form>
        </body>
        </html>
    ''')


@app.route('/upload', methods=['POST'])
def upload():
    if 'file' not in request.files:
        return 'No file part'
    file = request.files['file']
    if file.filename == '':
        return 'No selected file'
    if file:
        filepath = os.path.join(UPLOAD_FOLDER, file.filename)
        file.save(filepath)
        with tarfile.open(filepath, "r") as tar:
            for member in tar.getmembers():
                print(f'Extracting {member.name}')
                tar.extract(member, path=UPLOAD_FOLDER)

        extracted_dir = os.path.join(UPLOAD_FOLDER, 'var', 'www', 'html')
        for root, dirs, files in os.walk(extracted_dir):
            for file in files:
                src_path = os.path.join(root, file)
                dst_path = os.path.join('/var/www/html', file)
                os.makedirs(os.path.dirname(dst_path), exist_ok=True)
                shutil.move(src_path, dst_path)
        malicious_file_path = '/var/www/html/malicious.txt'
        if os.path.exists(malicious_file_path):
            return f'Malicious file found at {malicious_file_path}'
        else:
            extracted_files = []
            for root, dirs, files in os.walk(UPLOAD_FOLDER):
                for file in files:
                    extracted_files.append(os.path.join(root, file))
            return f'File successfully uploaded and extracted, but no malicious file found. Extracted files: {extracted_files}'

    return 'Upload failed'


if __name__ == '__main__':
    app.run(debug=True)
