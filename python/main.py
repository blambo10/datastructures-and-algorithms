# This is a sample Python script.

# Press ⌃R to execute it or replace it with your code.
# Press Double ⇧ to search everywhere for classes, files, tool windows, actions, and settings.

def main():
    numbers = [5, 2, 4, 6, 1, 3]
    insertion_sort(numbers, 6)
    print(numbers)

def insertion_sort(A, n):
    ind = 0

    while ind < len(A) - 1:
        if ind >= n:
            return
        j = ind
        while j > 0:
            if A[j-1] > A[j]:
                A[j-1], A[j] = A[j], A[j-1]

            j = j - 1
        ind = ind+1




# Press the green button in the gutter to run the script.
if __name__ == '__main__':
    main()

# See PyCharm help at https://www.jetbrains.com/help/pycharm/
