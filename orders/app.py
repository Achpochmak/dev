from flask import Flask, request
import json

app = Flask(__name__)

@app.route('/orders/', methods=['POST'])
def order():
    data = request.get_json()
    with open('orders.txt', 'a') as f:
        f.write(json.dumps(data) + '\n')
    return 'Order received', 200

@app.route('/orders/', methods=['GET'])
def health_check():
    return 'OK', 200

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5001)
