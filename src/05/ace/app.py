import sys
import webbrowser
import glob
import signal

from threading import Timer
from flask import Flask 
from flask import render_template
from flask import request
from flask import url_for
from flask import redirect
from flask import jsonify

app = Flask(__name__)

supported_formats = {
	"py": "python",
	"go": "golang",
	"c": "c",
	"js": "javascript"
}


def signal_handler(sig, frame):
	print('exiting...')
	sys.exit(0)


def examine_outputs(format: str):
	if format in supported_formats:
		return supported_formats[format]
	return None

def write_file(content: str):
	fd = open('TESTING', 'w')
	fd.write(content)
	fd.close()

def open_browser():
	webbrowser.open_new('http://127.0.0.1:5000/')

def get_name():
	filename = glob.glob("..\\session\\*")

	return filename[0]


@app.route('/save', methods=['POST'])
def save():
	content = request.get_json()['data']
	print(content)
	write_file(content)
	return jsonify(content)

@app.route('/')
def index():
	filename = get_name()
	content = open(filename).read()
	extension = filename.split(".")
	ext = ''
	if len(extension) > 1:
		ext = examine_outputs(extension[-1])
	print(ext)
	return render_template("editor.html", source_code=content, ext=ext)



if __name__ == '__main__':
	signal.signal(signal.SIGINT, signal_handler)
	Timer(0.5, open_browser).start();
	app.run()