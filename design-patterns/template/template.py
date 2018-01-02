from abc import ABCMeta, abstractmethod


class SortingTemplate(object):
    __metaclass__ = ABCMeta

    @abstractmethod
    def compare(self, a, b):
        raise NotImplementedError

    def sort(self, l):
        return sorted(l, cmp=self.compare)


class AscendingSorting(SortingTemplate):

    def compare(self, a, b):
        return a - b


class DescendingSorting(SortingTemplate):

    def compare(self, a, b):
        return b - a


if __name__ == '__main__':
    l = [1, 7, 4, 8, 5, 10]

    asc = AscendingSorting()
    asc_list = asc.sort(l)
    print asc_list

    desc = DescendingSorting()
    desc_list = desc.sort(l)
    print desc_list
