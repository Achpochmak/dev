import requests

def test_products():
    response = requests.get("http://localhost:5000/api/products")
    assert response.status_code == 200
