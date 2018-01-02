import threading


class Singleton(object):
    """
    A thread safe singleton base class.
    It can be inherited to create new classes with get_instance
    """
    __instance = None
    __instance_write_lock = threading.Lock()

    @classmethod
    def get_instance(cls):
        if not cls.__instance:
            with cls.__instance_write_lock:
                if not cls.__instance:
                    cls.__instance = cls()
        return cls.__instance


class A(Singleton):
    def __init__(self):
        import datetime
        self._name = 'A' + str(datetime.datetime.utcnow())

    @property
    def name(self):
        return self._name


if __name__ == '__main__':
    import thread

    # Some function to create a
    def create_a():
        try:
            a = A.get_instance()
            print a.name
        except Exception as e:
            print "Error in thread %d: %s" % (thread.get_ident(), str(e))

    try:
        thread_list = list()
        for i in range(0, 10):
            t = threading.Thread(target=create_a)
            thread_list.append(t)
            t.start()

        for t in thread_list:
            t.join()
    except Exception as e:
        print "Exceptions: %s" % str(e)
    finally:
        print "Jobs complete"
