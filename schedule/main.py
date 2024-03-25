# import faker

class Teacher:
    '''
    Класс, представляющий учителя.

    Атрибуты:
        id (int): Идентификатор учителя.
        name (str): Имя учителя.
        weekend_day (int): Список дней недели, когда учитель не работает (1 - понедельник, 7 - воскресенье).
        pref_slots (list[tuple[int, int]]): Список предпочитаемых временных слотов для занятий (пара (день недели, номер пары)).
        pref_aud (list[str]): Список предпочитаемых аудиторий для занятий.
        pairs (list[Pair]): Список пар, которые ведёт учитель.
    '''
    def __init__(self, id, name, weekend_day, pref_slots: [(int, int)], pref_aud: [str]):
        '''
        Инициализация экземпляра класса Teacher.
        '''
        self.id = id
        self.name = name
        self.pref_slots = pref_slots
        self.pref_aud = pref_aud
        self.weekend_days = [i for i in range(weekend_day, 190, 7)]
        self.pairs = []

    def __eq__(self, other_teacher):
         '''
        Сравнение двух экземпляров класса Teacher.

        Args:
            other_teacher (Teacher): Другой экземпляр класса Teacher.

        Returns:
            bool: True, если экземпляры эквивалентны, False - в противном случае
        '''
        return self.id == other_teacher.id and self.name == other_teacher.name


class Subject:
    '''
    Класс, представляющий урок.

    Атрибуты:
        id (int): Идентификатор урока.
        lessons (int): Количество пар по данному предмету, исходя из академических часов.
        teacher (Teacher): Преподаватель, который ведёт урок.
        type (string): Тип пары(практика/лекция).
        name (str): Название дисциплины.
    '''
    def __init__(self, id, hours, teacher: Teacher, type, name):
        '''
        Инициализация экземпляра класса Subject.
        '''
        self.id = id
        self.lessons = hours // 2
        self.teacher = teacher
        self.type = type
        self.name = name

    def __copy__(self):
        '''
        Создание копии объекта.
        
        Returns:
            Subject: Копия объекта.
        '''
        return Subject(self.id, self.lessons * 2, self.teacher, self.type, self.name)

    def __eq__(self, other_subject) -> bool:
         '''
        Сравнение двух объектов.

        Args:
            other_subject (Subject): Другой объект класса Subject.

        Returns:
            bool: True, если объекты эквивалентны, False - в противном случае.
        '''
        return (self.id == other_subject.id and
                self.name == other_subject.name and
                self.lessons == other_subject.lessons and
                self.teacher.id == other_subject.teacher.id and
                self.type == other_subject.type)


class Group:
    '''
    Класс, представляющий группу студентов.

    Атрибуты:
        id (int): Идентификатор группы.
        year (int): Курс группы.
        admission (int): Год поступления (11 - первый курс, 12 - второй и т.д.).
        direction (str): Направление обучения.
        subjects (dict[int, Subject]): Словарь, где ключ - идентификатор дисциплины, значение - объект класса Subject.
        pairs (list[Pair]): Список уроков, которые посещает группа.
    '''
    def __init__(self, id, year, admission, direction, subjects: {int: Subject}):
        '''
        Инициализация экземпляра класса Group.
        '''
        self.id = id
        self.year = year
        self.classes = admission
        if admission == 11:
            self.admission = 0
        else:
            self.admission = 1
        self.direction = direction
        self.subjects = subjects
        self.pairs = []

    def __eq__(self, other):
        '''
        Сравнение двух объектов.

        Args:
            other (Group): Другой объект класса Group.

        Returns:
            bool: True, если объекты эквивалентны, False - в противном случае.
        '''
        return self.id == other.id


class Room:
    '''
    Класс, представляющий аудиторию.

    Атрибуты:
        number (str): Номер аудитории.
        type (str): Тип аудитории (лекционная, компьютерный класс и т.д.).
        pairs (list[Pair]): Список уроков, которые проводятся в данной аудитории.
    '''
    def __init__(self, number, type):
        '''
        Инициализация экземпляра класса Room.
        '''
        self.number = number
        self.type = type
        self.pairs = []

    def __eq__(self, other):
        '''
        Сравнение двух объектов.

        Args:
            other (Room): Другой объект класса Room.

        Returns:
            bool: True, если объекты эквивалентны, False - в противном случае.
        '''
        return self.number == other.number


class Schedule:
    def __init__(self, groups: dict, rooms: dict, holidays, teachers: dict, days: int):
        self.groups = groups
        self.rooms = rooms
        self.holidays = holidays
        self.teachers = teachers
        self.schedule = list()
        self.days = days
        self.generate(with_pref=True)

    def generate(self, with_pref: bool):
        for pair in range(1, 7):
            for day in range(1, self.days):
                if day not in self.holidays:
                    self._add_pair(day, pair, with_pref)

    def _check_lecture_compatibility(self, group1, group2):
        return (self.groups[group1].admission + self.groups[group1].year == self.groups[group2].admission + self.groups[
            group2].year)

    def is_free(self, pair, teacher_id: int, room_id: int, group_id: int):
        for other_pair in self.teachers[teacher_id].pairs:
            if other_pair[0][0] == pair[0] and other_pair[0][1] == pair[1]:
                return False
        for other_pair in self.rooms[room_id].pairs:
            if other_pair[0][0] == pair[0] and other_pair[0][1] == pair[1]:
                return False
        for other_pair in self.groups[group_id].pairs:
            if other_pair[0][0] == pair[0] and other_pair[0][1] == pair[1]:
                return False
        pair_for_teacher = (pair[0] % 7, pair[1])
        if pair[0] in self.teachers[teacher_id].weekend_days or pair_for_teacher in self.teachers[teacher_id].pref_slots:
            return False
        return True

    def _add_pair(self, day, pair, with_pref: bool):
        pair = (day, pair)
        for group_id in self.groups.keys():
            if with_pref:
                tmp = sorted(self.groups[group_id].subjects.keys(),
                             key=lambda x: len(self.groups[group_id].subjects[x].teacher.pref_slots),
                             reverse=True)
            else:
                tmp = self.groups[group_id].subjects.keys()
            for subject_id in tmp:
                if self.groups[group_id].subjects[subject_id].lessons > 0:
                    # Практика
                    if self.groups[group_id].subjects[subject_id].type == 1:
                        if with_pref and self.groups[group_id].subjects[subject_id].teacher.pref_aud:
                            rooms=[]
                            for room_name in self.groups[group_id].subjects[subject_id].teacher.pref_aud:
                                rooms.append(room_name)
                        else:
                            rooms = self.rooms.keys()
                        for room_id in rooms:
                            if self.rooms[room_id].type == "Для практики" and self.is_free(pair, self.groups[group_id].subjects[subject_id].teacher.id, room_id, group_id):
                                self.schedule.append((pair, self.groups[group_id], self.groups[group_id].subjects[subject_id], self.rooms[room_id]))
                                self.teachers[self.groups[group_id].subjects[subject_id].teacher.id].pairs.append((pair, self.groups[group_id], self.groups[group_id].subjects[subject_id], self.rooms[room_id]))
                                self.groups[group_id].pairs.append((pair, self.groups[group_id], self.groups[group_id].subjects[subject_id], self.rooms[room_id]))
                                self.groups[group_id].subjects[subject_id].lessons -= 1
                                if self.groups[group_id].subjects[subject_id].lessons < 0:
                                    print("Перевыполнили план1")
                                self.rooms[room_id].pairs.append((pair, self.groups[group_id], self.groups[group_id].subjects[subject_id], self.rooms[room_id]))
                                break
                        break
                    # Лекция
                    else:
                        if with_pref and self.groups[group_id].subjects[subject_id].teacher.pref_aud:
                            rooms = []
                            for room_name in self.groups[group_id].subjects[subject_id].teacher.pref_aud:
                                rooms.append(room_name)
                        else:
                            rooms = self.rooms.keys()
                        for room_id in rooms:
                            if self.rooms[room_id].type == "Для лекций" and self.is_free(pair,
                                                                                         self.groups[group_id].subjects[
                                                                                             subject_id].teacher.id,
                                                                                         room_id, group_id):
                                available_groups = []
                                for other_group_id in self.groups.keys():
                                    tmp = False
                                    for sub_id in self.groups[other_group_id].subjects.keys():
                                        if sub_id == subject_id:
                                            tmp = True
                                            break
                                    if tmp:
                                        if not self.is_free(pair,
                                                            self.groups[other_group_id].subjects[subject_id].teacher.id,
                                                            room_id, other_group_id) or self.groups[other_group_id] == \
                                                self.groups[group_id]:
                                            continue
                                        if (self._check_lecture_compatibility(group_id, other_group_id) and
                                                self.groups[other_group_id].subjects[subject_id].lessons > 0):
                                            if len(available_groups) < 3:
                                                available_groups.append((other_group_id))

                                for other_group_id in available_groups:
                                    if self.groups[other_group_id].subjects[subject_id].lessons == 0:
                                        print("У другой группы 0 часов")
                                    self.schedule.append((pair, self.groups[other_group_id],
                                                          self.groups[other_group_id].subjects[subject_id],
                                                          self.rooms[room_id]))
                                    self.groups[other_group_id].subjects[subject_id].lessons -= 1
                                    if self.groups[other_group_id].subjects[subject_id].lessons < 0:
                                        print("Перевыполнили план2")
                                    self.groups[other_group_id].pairs.append((pair, self.groups[other_group_id],
                                                                              self.groups[other_group_id].subjects[
                                                                                  subject_id], self.rooms[room_id]))
                                if self.groups[group_id].subjects[subject_id].lessons == 0:
                                    print("У нашей группы 0 часов")
                                self.schedule.append((pair, self.groups[group_id],
                                                      self.groups[group_id].subjects[subject_id], self.rooms[room_id]))
                                self.teachers[self.groups[group_id].subjects[subject_id].teacher.id].pairs.append((pair,
                                                                                                                   self.groups[
                                                                                                                       group_id],
                                                                                                                   self.groups[
                                                                                                                       group_id].subjects[
                                                                                                                       subject_id],
                                                                                                                   self.rooms[
                                                                                                                       room_id]))
                                self.groups[group_id].pairs.append((pair, self.groups[group_id],
                                                                    self.groups[group_id].subjects[subject_id],
                                                                    self.rooms[room_id]))
                                self.rooms[room_id].pairs.append((pair, available_groups.append(group), subject, room))
                                self.groups[group_id].subjects[subject_id].lessons -= 1
                                if self.groups[group_id].subjects[subject_id].lessons < 0:
                                    print("Перевыполнили план3")
                                break
                        break

    def check(self):
        check = True
        res=''
        for group_id, group in self.groups.items():
            for subject_id, subject in group.subjects.items():
                if self.groups[group_id].subjects[subject_id].lessons != 0:
                    print("Несоответствие с учебным планом:")
                    res = f"У группы {self.groups[group_id].year}-{self.groups[group_id].classes}-{self.groups[group_id].id} {self.groups[group_id].direction} "
                    res += f"Не хватает {self.groups[group_id].subjects[subject_id].lessons} пар по предмету {self.groups[group_id].subjects[subject_id].name}, которые ведёт {self.groups[group_id].subjects[subject_id].teacher.name};"
                    print(res)
                    check = False
        for _ in range(len(res)):
            print("-", end="")
        print()
        return check


# # Создаем Faker-объект
# faker = faker.Faker()
# # Создаем новые комнаты
# rooms = {}
# for i in range(30):
#     room = Room(i, faker.random_element(["Для лекций", "Для практики"]))
#     rooms[i] = room
# # Создаем новые группы
# groups = {}
# for i in range(20):
#     group = Group(id=i, admission=faker.random_element([9, 11]), year=faker.random_int(2021, 2024),
#                   direction=faker.random_element(["Сисы", "Инфобез", "Прогеры"]), subjects={})
#     groups[i] = group
# # Создаем новых учителей
# teachers = {}
# for i in range(10):
#     teacher = Teacher(id=i, name=faker.name(), weekend_day=faker.random_int(1, 6),
#                       pref_slots=[(faker.random_int(1, 6), faker.random_int(1, 6)) for _ in
#                                   range(faker.random_int(0, 3))],
#                       pref_aud=list(set([faker.random_int(0,len(rooms)-1) for _ in range(0,6)])))

#     teachers[i] = teacher

# # Создаем новые предметы
# subjects = {}
# for i in range(10):
#     subject = Subject(id=i, hours=faker.random_int(2, 10), teacher=faker.random_element(teachers.values()),
#                       type=faker.random_int(0, 1), name=faker.random_element(
#             ["Математика", "Информатика", "Физика", "Химия", "Биология", "История", "География", "Литература",
#              "Иностранный язык"]))
#     subjects[i] = subject

# # Добавляем предметы группам
# for group_id, group in groups.items():
#     for _ in range(faker.random_int(2, 4)):
#         subject_id = faker.random_element(list(subjects.keys()))
#         while subject_id in group.subjects.keys():
#             subject_id = faker.random_element(list(subjects.keys()))
#         group.subjects[subject_id] = subjects[subject_id].__copy__()

# # Увеличиваем количество занятий по каждому предмету
# for subject_id, subject in subjects.items():
#     subject.lessons *= 2

rooms = {
    101: Room(101, "Для практики"),
    102: Room(102, "Для лекций"),
    103: Room(103, "Для практики"),
    104: Room(104, "Для практики"),
    105: Room(105, "Для лекций"),
    106: Room(106, "Для лекций"),
}
teachers = {
    1: Teacher(1, "Иванов Иван", 3, [(6, 3)], []),
    2: Teacher(2, "Петров Петр", 4, [], []),
    3: Teacher(3, "Цветкова Елена", 5, [(6, 3), (2, 1), (1, 1)], []),
    4: Teacher(4, "Чернова Анна", 6, [], []),
}
subjects = {
    1: Subject(1, 4, teachers[3], 0, "Администрирование"),
    2: Subject(2, 4, teachers[3], 0, "Сети и протоколы"),
    3: Subject(3, 4, teachers[1], 1, "Администрирование"),
    4: Subject(4, 4, teachers[1], 1, "Сети и протоколы"),
    5: Subject(5, 2, teachers[2], 1, "Алгебра"),
    6: Subject(6, 2, teachers[2], 0, "Алгебра"),
    7: Subject(7, 2, teachers[4], 0, "Программирование"),
    12: Subject(12, 2, teachers[4], 1, "Программирование"),
    8: Subject(8, 2, teachers[4], 1, "Алгоритмы"),
    9: Subject(9, 2, teachers[2], 0, "Алгоритмы"),
}
groups = {
    1: Group(1, 2022, 9, "Инфобезы", {}),
    2: Group(2, 2022, 9, "Сисадмины", {}),
    3: Group(3, 2023, 11, "Сисадмины", {}),
    4: Group(4, 2021, 11, "Комплексы", {}),
    5: Group(5, 2022, 9, "Программисты", {}),
    6: Group(6, 2022, 11, "Инфобезы", {}),
}

short_days=[5]
for i in range(len(teachers)):
    for j in short_days:
        for u in range(3,7):
            teachers[i+1].pref_slots.append((j,u))
for group in groups.values():
    for subject_id in list(subjects.keys()):
        group.subjects[subject_id] = subjects[subject_id].__copy__()

for i in range(5):
    groups[6].subjects.pop(i+1, None)
    groups[4].subjects.pop(i+2, None)
    groups[3].subjects.pop(i+1, None)

# Добавляем больше праздников
holidays = [i for i in range(7, 190, 7)]

# Проверка корректности предпочтений учителей
glcheck = True
for subject in subjects.values():
    if subject.teacher.pref_aud:
        check = False
        for room in subject.teacher.pref_aud:
            if rooms[room].type == "Для лекций":
                roomtype = 0
            else:
                roomtype = 1
            if roomtype == subject.type:
                check = True
                break
        if not check:
            glcheck = False
            print(f"У {subject.teacher.name} не выбрано ни одной аудитории, подходящей для предмета {subject.name}")


# if glcheck:
weekend = {
    1: "Понедельник",
    2: "Вторник",
    3: "Среда",
    4: "Четверг",
    5: "Пятница",
    6: "Суббота",
    7: "Воскресенье",
}

# Создаем расписание
schedule = Schedule(groups, rooms, holidays, teachers, 7)
if schedule.schedule is None or not schedule.check():
    print("Failed to generate a valid schedule")
else:
    # Print the schedule
    s = sorted(schedule.schedule, key=lambda x: x[0][0])
    tmp = s[0][0]
    for idx, pair in enumerate(s):
        res = ""
        if tmp != pair[0][0] or idx == 0:
            print()
            res += f"День: {pair[0][0]} ({weekend[pair[0][0]%7]})\n"
            tmp = pair[0][0]
        res += f"   Пара №{pair[0][1]} \n\tГруппа: {pair[1].direction}, {pair[1].classes} класс, {pair[1].year} года"
        res += f"\n\tПредмет: {pair[2].name} \n\tУчитель: {pair[2].teacher.name} \n\tАудитория: {pair[3].number}, {pair[3].type}"
        print(res)

print("--"*65, "\n")
for t in teachers.values():
    if t.pref_aud:
        print(f"Предпочтения преподавателя {t.name} по аудиториям - {t.pref_aud}")
    if t.pref_slots:
        print(f"\tПредпочтения по слотам для {t.name}:")
        for i in t.pref_slots:
            print(f"НЕТ {weekend[i[1]]} пары в {weekend[i[0]]} ")
        print()
    print()

