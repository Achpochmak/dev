import requests

def test_products():
    response = requests.get("http://backend:5000/api/products")
    assert response.status_code == 200
