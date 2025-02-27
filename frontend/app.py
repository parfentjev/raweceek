from flask import Flask, render_template, send_from_directory
from requests import get, RequestException
from datetime import datetime, timezone
import os


app = Flask(__name__)


@app.route('/')
def index():
    countdown = get_countdown()

    if not countdown:
        return render_template('index.html')

    return render_template('index.html', data=countdown)


@app.route('/favicon.ico')
def favicon():
    return send_from_directory(os.path.join(app.root_path, 'static'),
                               'favicon.ico', mimetype='image/vnd.microsoft.icon')


def get_countdown():
    try:
        response = get('https://raweceek.eu/api/sessions/countdown')
        if response.status_code == 200:
            return response.json()
        else:
            return None
    except RequestException:
        return None


app.run(host='0.0.0.0', debug=False)
