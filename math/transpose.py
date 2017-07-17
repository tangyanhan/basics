def transpose(matrix):
    """
    Transpose matrix.
    :param matrix: 2d list containing numbers
    :return: transposed matrix
    """
    M = len(matrix)
    N = len(matrix[0])

    matrix_t = []
    for row in range(0, N):
        array = []
        for col in range(0, M):
            array.append(0)
        matrix_t.append(array)

    for m in range(0, M):
        for n in range(0):
            matrix_t[n][m] = matrix[m][n]

    return matrix_t


if __name__ == '__main__':
    m = [[1, 2], [3, 4], [5, 6]]
    print("Original matrix:\n", m)
    print("Transposed matrix:\n", transpose(m))
