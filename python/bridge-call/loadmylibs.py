import ctypes
import os


def load_dylib_from_single_path(src_path, dylib_path, is_cpp=False):
    assert os.path.isfile(src_path), "Source file {0} doesn't exist".format(src_path)

    if not os.path.isfile(dylib_path):
        compiler = 'gcc' if not is_cpp else 'g++'
        print 'Compiler=', compiler, ' src=', src_path
        os.system("{2} -o {0} -shared -fPIC {1}".format(dylib_path, src_path, compiler))

    return ctypes.cdll.LoadLibrary(dylib_path)