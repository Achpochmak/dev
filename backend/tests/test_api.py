import requests

def test_products():
    response = requests.get("http://localhost/orders/")
    assert response.status_code == 200
