#!/bin/python3

if __name__ == "__main__":
    
    birthdays = {
        "Alice": "Apr 1",
        "Bob": "Dec 12",
        "Carol": "Mar 4"
     }

    while True:
        print("Bob, Alice or Carol: (blank to quit)")
        name = input()

        if name == "":
            break

        if name in birthdays:
            print(birthdays[name] + " is the birthday of " + name)
            break
        else:
            print("I don't have birthday information for " + name)
            print("What is their birthday?")
            new_birthday = input()
            birthdays[name] = new_birthday
            print("Birthday Database Updated.")

