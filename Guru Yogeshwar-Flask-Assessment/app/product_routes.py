from flask import Blueprint, render_template, request, redirect

product_routes = Blueprint('product_routes', __name__)

products = [] 


@product_routes.route('/products')
def index():
    return render_template('products/index.html', products=products)  

@product_routes.route('/products/create', methods=['GET', 'POST'])
def create():
    if request.method == 'POST':
        new_product_id = int(request.form['id'])

        
        new_product = {
            'id': new_product_id,  
            'name': request.form['name'],  
            'price': request.form['price'] 
        }

        products.append(new_product) 
        print(products)  

        return redirect('/products')  

    return render_template('products/create.html')

@product_routes.route('/products/update/<int:product_id>', methods=['GET', 'POST'])
def update(product_id):
    
    product = products[product_id - 1] 
    
    if request.method == 'POST':
        
        updated_product = {
            'id': product_id,
            'name': request.form['name'],
            'price': request.form['price']
        }
        
        
        products[product_id - 1] = updated_product  
        return redirect(('/products'))    
        
        
        

    
    return render_template('products/update.html', product=product, product_id=product_id)
