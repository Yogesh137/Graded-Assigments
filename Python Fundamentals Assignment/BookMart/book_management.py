class Book:
    def __init__(self, title, author, price, quantity):
        self.title = title
        self.author = author
        self.price = price
        self.quantity = quantity

    def display_details(self):
        return f"Title: {self.title}, Author: {self.author}, Price: {self.price}, Quantity: {self.quantity}"

books = []

def add_book(title, author, price, quantity):
    try:
        if price <= 0 or quantity <= 0:
            raise ValueError("Price and quantity must be positive numbers.")
        book = Book(title, author, price, quantity)
        books.append(book)
        return "Book added successfully!"
    except ValueError as e:
        return str(e)

def view_books():
    if not books:
        return "No books available."
    return [book.display_details() for book in books]

def search_book(query):
    results = [book for book in books if query.lower() in book.title.lower() or query.lower() in book.author.lower()]
    if not results:
        return "No books found."
    return [book.display_details() for book in results]