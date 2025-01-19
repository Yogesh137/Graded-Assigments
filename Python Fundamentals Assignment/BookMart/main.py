from book_management import add_book, view_books, search_book
from customer_management import add_customer, view_customers
from sales_management import sell_book, view_sales

def main_menu():
    while True:
        print("\nWelcome to BookMart!")
        print("1. Book Management")
        print("2. Customer Management")
        print("3. Sales Management")
        print("4. Exit")
        choice = input("Enter your choice: ")

        if choice == "1":
            print("\nBook Management")
            print("1. Add Book")
            print("2. View Books")
            print("3. Search Book")
            book_choice = input("Enter your choice: ")

            if book_choice == "1":
                title = input("Enter title: ")
                author = input("Enter author: ")
                try:
                    price = float(input("Enter price: "))
                    quantity = int(input("Enter quantity: "))
                    print(add_book(title, author, price, quantity))
                except ValueError:
                    print("Invalid input! Price and quantity must be numbers.")

            elif book_choice == "2":
                books = view_books()
                if isinstance(books, str):
                    print(books)
                else:
                    for book in books:
                        print(book)

            elif book_choice == "3":
                query = input("Enter title or author to search: ")
                results = search_book(query)
                if isinstance(results, str):
                    print(results)
                else:
                    for book in results:
                        print(book)

        elif choice == "2":
            print("\nCustomer Management")
            print("1. Add Customer")
            print("2. View Customers")
            customer_choice = input("Enter your choice: ")

            if customer_choice == "1":
                name = input("Enter name: ")
                email = input("Enter email: ")
                phone = input("Enter phone: ")
                print(add_customer(name, email, phone))

            elif customer_choice == "2":
                customers = view_customers()
                if isinstance(customers, str):
                    print(customers)
                else:
                    for customer in customers:
                        print(customer)

        elif choice == "3":
            print("\nSales Management")
            print("1. Sell Book")
            print("2. View Sales")
            sales_choice = input("Enter your choice: ")

            if sales_choice == "1":
                customer_name = input("Enter customer name: ")
                book_title = input("Enter book title: ")
                try:
                    quantity = int(input("Enter quantity: "))
                    print(sell_book(customer_name, book_title, quantity))
                except ValueError:
                    print("Invalid input! Quantity must be a number.")

            elif sales_choice == "2":
                sales = view_sales()
                if isinstance(sales, str):
                    print(sales)
                else:
                    for sale in sales:
                        print(sale)

        elif choice == "4":
            print("Exiting BookMart. Welcome Again!")
            break

        else:
            print("Invalid choice. Please try again.")

if __name__ == "__main__":
    main_menu()