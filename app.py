from flask import Flask, render_template, request

app = Flask(__name__)


@app.route("/")
def hello():
    return "Hello World!"


@app.route("/pay")
def pay():
    print(request.get_data())
    return render_template("pay.html")


if __name__ == "__main__":
    app.run()
