def bubble_sort(l):
    for j in range(0, len(l)):
        has_swap = False
        start = 0
        end = len(l) - start - 1
        while start < end:
            if l[start] > l[start+1]:
                has_swap = True
                l[start], l[start+1] = l[start+1], l[start]
            start += 1
        if not has_swap:
            break


if __name__ == '__main__':
    l = [1, 7, 4, 6, 10, 13, 2, 3]
    bubble_sort(l)
    print l
