from flask import Flask, jsonify
app = Flask(__name__)

@app.route('/api/products')
def products():
    return jsonify([
        {"id": 1, "name": "T-shirt", "price": 19.99},
        {"id": 2, "name": "Mug", "price": 9.99}
    ])

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)