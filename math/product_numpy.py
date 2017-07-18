import numpy

a = numpy.mat([1, 2])
b = numpy.mat([[10], [20]])

print("a*b=", a * b)
print("a.T * b.T", a.T * b.T)

a = numpy.mat([[1, 2], [3, 4]])
b = numpy.mat([[10, 20], [30, 40]])
print("a*b=", a * b)
