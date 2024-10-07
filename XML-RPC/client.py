import xmlrpc.client

proxy = xmlrpc.client.ServerProxy("http://localhost:9000/")

def switch_operation(proxy):
    flag = True
    while flag:
        try: 
            number_a = float(input("Give the first number (a): "))
            number_b = float(input("Give the second number (b): "))
            print("1: Add   2: Subtract   3: Multiply   4: Division")
            print("Just select one. Write the number.")
            operation = int(input("Which operation do you want to do? "))

            if operation == 1:
                result = proxy.add(number_a, number_b)
                print(f"Result of {number_a} + {number_b} is: {result}")
            elif operation == 2:
                result = proxy.rest(number_a, number_b)  
                print(f"Result of {number_a} - {number_b} is: {result}")
            elif operation == 3:
                result = proxy.mul(number_a, number_b)
                print(f"Result of {number_a} x {number_b} is: {result}")
            elif operation == 4:
                if number_b == 0:
                    print("Error: Division by zero is not allowed.")
                else:
                    result = proxy.div(number_a, number_b)
                    print(f"Result of {number_a} / {number_b} is: {result}")
            else:
                print("You didn't choose a valid option. This isn't the way.") #GOD BLESS STAR WARS

            print("Do you want to make another operation?")

            flag_keep = True
            while flag_keep:

                keep = input("y/n? ").lower()  
                if keep == "y":
                    flag = True
                    flag_keep = False  
                elif keep == "n":
                    flag = False
                    flag_keep = False 
                else:
                    print("You did not choose a valid option. This isn't the way.") #HELLO THERE
                    flag_keep = True  

        except ValueError:
            print("Please, write correct numbers.")
    
    print("Goodbye!")


def main():
    proxy = xmlrpc.client.ServerProxy("http://localhost:9000/")
    switch_operation(proxy)

if __name__ == "__main__":
    main()

