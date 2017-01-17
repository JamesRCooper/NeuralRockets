var _self = this;

var bars;
var socket;

function setup() {
  createCanvas(700, 400);
  bars = new Rockets();
}

function draw() {
  background(50, 89, 100);
  rect(200, 180, 300, 40);
  bars.display();
}

function Rockets(positions) {

  var ws = new WebSocket("ws://localhost:80/ws");

  ws.onopen = function() {
    console.log('Connected');
  }

  ws.onmessage = function(evt) {
    _self.poss = JSON.parse(evt.data);
  }

  ws.onclose = function () {
    console.log("Connection closed");
  };

  this.display = function() {
    if (!_self.poss) { return; }

    for (var i = 0; i < _self.poss.length; i++) {
      rocket = _self.poss[i];
      push();
      translate(rocket["X"], rocket["Y"]);
      rotate(createVector(1,0).rotate(rocket["A"] + 0.5 * PI).heading());
      rectMode(CENTER);
      rect(0, 0, 10, 50);
      pop();
    }
  }

  this.test = function() {
    console.log("Hello!")
  };
}
