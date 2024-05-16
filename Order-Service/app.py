from flask import Flask
import logging
app = Flask(__name__)

logging.basicConfig(filename='/var/log/order.log', level=logging.INFO)
logging.info('Hello from Order-Service')


@app.route('/')
def home():
    logging.info('Hello from Order-Service')
    return 'Hello from Order-Service'

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)

