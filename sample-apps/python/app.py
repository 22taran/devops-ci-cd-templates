from flask import Flask, jsonify

app = Flask(__name__)


@app.route("/")
def home():
    return jsonify({"message": "Hello, World!"})


@app.route("/health")
def health():
    return jsonify({"status": "healthy"})


def add(a: int, b: int) -> int:
    return a + b


def greet(name: str = None) -> str:
    if not name or name.strip() == "":
        return "Hello, World!"
    return f"Hello, {name}!"


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)
