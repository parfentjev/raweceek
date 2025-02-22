from flask import Flask, render_template, send_from_directory
from requests import get, RequestException
from datetime import datetime, timezone
import os


app = Flask(__name__)


@app.route('/')
def index():
    session = get_next_session()

    if not session:
        return render_template('index.html')

    target_date = datetime.strptime(session['startTime'], '%Y-%m-%dT%H:%MZ')
    current_date = datetime.now(timezone.utc)
    race_week = target_date.isocalendar()[1] == current_date.isocalendar()[1]

    return render_template('index.html', data={'race_week': race_week, 'json': session})


@app.route('/favicon.ico')
def favicon():
    return send_from_directory(os.path.join(app.root_path, 'static'),
                               'favicon.ico', mimetype='image/vnd.microsoft.icon')


def get_next_session():
    try:
        response = get('https://raweceek.eu/api/sessions/next')
        if response.status_code == 200:
            return response.json()
        else:
            return None
    except RequestException:
        return None


app.run(host='0.0.0.0', debug=False)
