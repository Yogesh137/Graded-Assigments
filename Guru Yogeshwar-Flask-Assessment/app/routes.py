from flask import Blueprint, render_template

main_routes = Blueprint('main_routes', __name__)

# This is the route that will be displayed when the user visits the root URL of the site.


@main_routes.route('/')
def home():
    return render_template('home.html', title='Home', message='This is the home page.')


@main_routes.route('/about')
def about():
    return render_template('about.html', title='About', message='This is the about page.')