import loadmylibs


lib = loadmylibs.load_dylib_from_single_path('./myclib.c', './myclib.so')

result = lib.sum(1, 3)

print 'Load c lib, result =', result

cpplib = loadmylibs.load_dylib_from_single_path('./mycpplib.cc', './mycpplib.so', is_cpp=True)
result = cpplib.ext_sum(1, 3)
print 'Load cpp lib, result =', result

