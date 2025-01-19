from book_management import books
from customer_management import customers

class Transaction:
    def __init__(self, customer_name, book_title, quantity_sold):
        self.customer_name = customer_name
        self.book_title = book_title
        self.quantity_sold = quantity_sold

    def display_details(self):
        return f"Customer: {self.customer_name}, Book: {self.book_title}, Quantity: {self.quantity_sold}"

sales = []

def sell_book(customer_name, book_title, quantity):
    try:
        book = next((book for book in books if book.title.lower() == book_title.lower()), None)
        if not book:
            return "Book not found."
        if book.quantity < quantity:
            return f"Error: Only {book.quantity} copies available. Sale cannot be completed."

        book.quantity -= quantity
        transaction = Transaction(customer_name, book.title, quantity)
        sales.append(transaction)
        return f"Sale successful! Remaining quantity: {book.quantity}"
    except Exception as e:
        return str(e)

def view_sales():
    if not sales:
        return "No sales records available."
    return [sale.display_details() for sale in sales]