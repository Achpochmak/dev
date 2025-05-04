import requests

def test_products():
    response = requests.get("http://localhost/api/products")
    assert response.status_code == 200
