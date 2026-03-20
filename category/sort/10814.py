import sys

class Person:
    age: int
    name: str
    reg: int

def main():
    members :list[Person] = []
    n = int(sys.stdin.readline())

    for i in range(0, n):        
        age, name = sys.stdin.readline().split()
        mm = Person()
        mm.age = int(age)
        mm.name = name
        mm.reg = i
        members.append(mm)

    members.sort(key=lambda x: (x.age, x.reg))

    for mm in members:
        print(f"{mm.age} {mm.name}")

if __name__ == "__main__":
    main()