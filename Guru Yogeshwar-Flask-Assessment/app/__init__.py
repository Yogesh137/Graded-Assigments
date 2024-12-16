from flask import Flask
from app.routes import main_routes
from app.product_routes import product_routes

def create_app():
    app = Flask(__name__)
    app.register_blueprint(main_routes)
    app.register_blueprint(product_routes)
    return app