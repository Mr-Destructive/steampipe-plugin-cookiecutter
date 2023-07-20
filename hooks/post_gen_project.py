import os

SUCCESS = "\x1b[1;32m [SUCCESS]: "

def go_get_dependencies():
    os.system("go get ./...")

def main():
    go_get_dependencies()
    print(SUCCESS + "Project initialized, keep up the good work!")

if __name__ == "__main__":
    main()
