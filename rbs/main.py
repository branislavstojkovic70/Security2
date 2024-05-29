from flask import Flask, request, render_template_string
import os
import tarfile

app = Flask(__name__)

html_form = '''
<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <title>Upload a file</title>
  </head>
  <body>
    <h1>Upload a tar file</h1>
    <form method="post" enctype="multipart/form-data">
      <input type="file" name="file" accept=".tar">
      <input type="submit" value="Upload">
    </form>
  </body>
</html>
'''

@app.route('/', methods=['GET', 'POST'])
def upload_file():
    if request.method == 'POST':
        if 'file' not in request.files:
            return "No file part"

        file = request.files['file']

        if file.filename == '':
            return "No selected file"

        if file:
            desktop_path = os.path.join(os.path.expanduser("~"), "Desktop")
            upload_path = os.path.join(desktop_path, file.filename)
            file.save(upload_path)

            if tarfile.is_tarfile(upload_path):
                with tarfile.open(upload_path) as tar:
                    tar.extractall(path=desktop_path)
                    extracted_files = [os.path.abspath(os.path.join(desktop_path, member.name)) for member in tar.getmembers()]

                # Opcionalno: obrisati tar fajl nakon ekstrakcije
                # os.remove(upload_path)

                return f"File uploaded and extracted to Desktop. Extracted files:<br>" + "<br>".join(extracted_files)
            else:
                return "Not a tar file"

    return render_template_string(html_form)

if __name__ == '__main__':
    app.run(debug=True)
