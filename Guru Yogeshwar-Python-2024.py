class Book:
    def __init__(self, book_id, title, author):
        self.book_id = book_id
        self.title = title
        self.author = author
        self.status = "available"  # available or borrowed
    
    def borrow(self):
        if self.status == "available":
            self.status = "borrowed"
        else:
            raise ValueError(f"The book '{self.title}' is already borrowed.")

    def return_book(self):
        if self.status == 'borrowed':
            self.status = 'available'
        else:
            raise ValueError(f"The book '{self.title}' is already available.")
    
class BorrowLimitExceededException(Exception):
    pass
class Member:
    def __init__(self, member_id, name, max_books):
        self.member_id = member_id
        self.name = name
        self.borrowed_books = []
        self.max_books = max_books

    def borrow_book(self, book):
        if len(self.borrowed_books) < self.max_books:
            if book.status == 'available':
                book.borrow()
                self.borrowed_books.append(book)
                print(f"'{self.name}' successfully borrowed '{book.title}'.")
            else:
                raise ValueError(f"The book '{self.title}' is already borrowed.")
        else:
            raise BorrowLimitExceededException(f"'{self.name}' has reached borrowing limit of {self.max_books} books.")

    def return_book(self, book):
        if book in self.borrowed_books:
            book.return_book()
            self.borrowed_books.remove(book)
            print(f"'{self.name}' successfully returned '{book.title}'.")
        else:
            raise ValueError(f"'{self.name}' did not borrow '{book.title}'.")

class RegularMember(Member):
    def __init__(self, member_id, name):
        super().__init__(member_id, name, max_books=3)

class PremiumMember(Member):
    def __init__(self, member_id, name):
        super().__init__(member_id, name, max_books=5)

class Library:
    def __init__(self):
        self.book_collection = []
        self.members = []

    def add_book(self, book):
        if book.title and book.author:
            self.book_collection.append(book)
            print(f"Book: '{book.title}' has been added successfully.")
        else:
            raise ValueError("Book must have a title and an author.")

    def register_member (self, member):
        if member.name:
            self.members.append(member)
            print(f"Member: '{member.name}' has been registered successfully.")
        else:
            raise ValueError("Invalid member data")

    def lend_book(self, member, book):
        if book in self.book_collection:
            member.borrow_book(book)
        elif book not in self.book_collection:
            raise ValueError("Book does not exist in the library.")
        elif member not in self.members:
            raise ValueError("Member does not exist.")
        

    def receive_return(self, member, book):
        if book in member.borrowed_books:
            member.return_book(book)
        elif book not in self.book_collection:
            raise ValueError("Book does not exist in the library.")
        elif member not in self.members:
            raise ValueError("Member does not exist.")
        
# Create a library instance
library = Library()

# Add books to the library
book1 = Book(book_id=1, title="Red Sea", author="Rutherford")
book2 = Book(book_id=2, title="American Things", author="Donald Pump")
book3 = Book(book_id=3, title="Ex President", author="Bin Laden")
book4 = Book(book_id=4, title="India ka Baap", author="Modiji")

library.add_book(book1)
library.add_book(book2)
library.add_book(book3)
library.add_book(book4)
print()

# Register members
regular_member = RegularMember(member_id=101, name="Ali")
premium_member = PremiumMember(member_id=102, name="Brahmi")

library.register_member(regular_member)
library.register_member(premium_member)
print()

# Lend books
try:
    library.lend_book(regular_member, book1)  
    library.lend_book(premium_member, book2) 
    print()
except ValueError as ve:
    print(ve)
    print()
except BorrowLimitExceededException as ble:
    print(ble)
    print()

# Return books
try:
    library.receive_return(regular_member, book1) 
    library.receive_return(premium_member, book2)
    print()
except ValueError as ve:
    print(ve)
    print()

# Test exceeding borrowing limit for regular members
try:
    library.lend_book(regular_member, book1)  # Ali borrows again
    library.lend_book(regular_member, book2)  # Ali borrows second book
    library.lend_book(regular_member, book3)  # Ali borrows third book
    library.lend_book(regular_member, book4)  # Ali tries to borrow fourth (should raise Exceptin cuz - limit = 3)
    print()
except ValueError as ve:
    print(ve)
    print()
except BorrowLimitExceededException as ble:
    print(ble)