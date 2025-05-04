fetch('/api/products')
  .then(res => res.json())
  .then(data => {
    const div = document.getElementById('products');
    data.forEach(p => {
      const el = document.createElement('div');
      el.innerText = p.name + ' - $' + p.price;
      div.appendChild(el);
    });
  });