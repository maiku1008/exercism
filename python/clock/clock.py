class Clock:
    def __init__(self, hour, minute):
        self.hours, self.minutes = divmod((hour * 60 + minute) % (24 * 60), 60)

    def __repr__(self):
        return "{:02d}:{:02d}".format(self.hours, self.minutes)

    def __eq__(self, other):
        return str(other) == str(self)

    def __add__(self, minutes):
        return Clock(self.hours, self.minutes + minutes)

    def __sub__(self, minutes):
        return Clock(self.hours, self.minutes - minutes)
