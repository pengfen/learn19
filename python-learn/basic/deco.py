def my_decorator(func):
	def wrapper(*args, **kwargs):
		print("调用前")
		result = func(*args, **kwargs)
		print("调用后")
		return result
	return wrapper

@my_decorator
def add(a, b):
	return a + b