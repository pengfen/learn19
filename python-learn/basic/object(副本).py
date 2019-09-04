#!/usr/bin/python

class Person(object):
	"""
	返回具有给定名称的person对象
	"""

	def __init__(self, name):
		self.name = name

	def get_detail(self):
		"""
		返回包含人名的字符串
		"""
		return self.name

class Student(Person):
	"""
	返回Student对象 采用name branch year 3个参数
	"""

	def __init__(self, name, branch, year)
	    Person.__init__(self, name)
	    self.branch = branch
	    self.year = year

	def get_detail(self):
		"""
		返回包含学生具体信息的字符串
		"""
		return "{} studies {} and is in {} year.".format(self.name, self.branch, self.year)

class Teacher(Person):
	"""
	返回 Teacher 对象 采用字符串列表作为参数
	"""
	def __init__(self, name, papers):
		Person.__init__(self, name)
		self.papers = papers

	def get_detail():
		return "{} teachers {}".format(self.name, ",".join(self.papers))
		
person1 = Person("ricky")
student1 = Student("ricky", "CSE", 200)
teacher1 = Teacher("ricky", ["C", "C++"])

print(person1.get_detail())
print(student1.get_detail())
print(teacher1.get_detail())