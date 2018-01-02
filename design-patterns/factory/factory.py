from abc import ABCMeta, abstractmethod, abstractproperty
from random import random


class IHumanCreator(object):
    __metaclass__ = ABCMeta

    @abstractmethod
    def produce(self):
        raise NotImplementedError


class IHuman(object):
    __metaclass__ = ABCMeta

    @abstractmethod
    def talk(self):
        raise NotImplementedError

    @abstractproperty
    def race(self):
        raise NotImplementedError

    def set_sex(self, sex):
        raise NotImplementedError

    def get_sex(self):
        raise NotImplementedError

    sex = abstractproperty(get_sex, set_sex)


class Black(IHuman):
    __metaclass__ = ABCMeta

    def __init__(self):
        self._sex = 'man'

    def talk(self):
        print '%s %s talks' % (self.race, self.sex)

    @property
    def race(self):
        return 'Black'

    def set_sex(self, sex):
        assert sex in ['man', 'woman']
        self._sex = sex

    def get_sex(self):
        return self._sex

    sex = property(fget=get_sex, fset=set_sex, doc="Get or set sex of black")


class Yellow(IHuman):
    __metaclass__ = ABCMeta

    def __init__(self):
        self._sex = 'male'

    def talk(self):
        print '%s %s talks' % (self.race, self.sex)

    @property
    def race(self):
        return 'Yellow'

    def set_sex(self, sex):
        assert sex in ['male', 'female']
        self._sex = sex

    def get_sex(self):
        return self._sex

    sex = property(get_sex, set_sex)


class BlackCreator(IHumanCreator):
    __metaclass__ = ABCMeta

    def produce(self):
        r = random()
        obj = Black()
        if r < 0.2:
            obj.sex = 'woman'
        else:
            obj.sex = 'man'
        return obj


class YellowCreator(IHumanCreator):
    __metaclass__ = ABCMeta

    def produce(self):
        r = random()
        obj = Yellow()
        if r < 0.4:
            obj.sex = 'female'
        else:
            obj.sex = 'male'
        return obj


if __name__ == '__main__':
    black_creator = BlackCreator()
    yellow_creator = YellowCreator()

    def add_count(count_dict, human):
        key = human.race + ':' + human.sex
        if key not in count_dict:
            count_dict[key] = 1
        else:
            count_dict[key] += 1

    stat_dict = {}
    for i in range(0, 1000):
        b = black_creator.produce()
        b.talk()

        add_count(stat_dict, b)
        y = yellow_creator.produce()
        y.talk()
        add_count(stat_dict, y)
    print stat_dict
