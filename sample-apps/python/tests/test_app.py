from app import add, greet


def test_add():
    assert add(2, 3) == 5


def test_add_negative():
    assert add(-1, 1) == 0


def test_greet_with_name():
    assert greet("DevOps") == "Hello, DevOps!"


def test_greet_default():
    assert greet() == "Hello, World!"


def test_greet_empty():
    assert greet("") == "Hello, World!"
