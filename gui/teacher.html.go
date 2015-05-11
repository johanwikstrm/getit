<html>
<head>
<title>GETit - teacher interface</title>
<style>
/*
  red: c71e31
  grey: 463335
  white: f2ecda
*/
body {
  font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
  margin: 0px;
  padding:0px;  
  background-color: #f2ecda;
  color: #c71e31;
}

div.info{
  position: absolute;
  top: 0px;
  left: 0px;
  width: 200px;
  text-align: center;
  opacity: 0.7;
  size:1.5em;
}

.url{
  font-size: 2em;
}

#nstudents {
  font-size: 2em;
}

#switchview{
  position: absolute;
  height: 6em;
  width: 7em;
  left: 0px;
  bottom: 0px;
}

.understanding{
  font-size: 2em;
  color: #c71e31;
}

</style>
<script src="http://d3js.org/d3.v3.min.js"></script>
<script src="//ajax.googleapis.com/ajax/libs/jquery/1.11.0/jquery.min.js"></script>
</head>
<body onresize="setScale()">
<div class="info">
Students connected to<br/> <span class="url">{{.SuperShortUrl}}</span></br>
<div id ="nstudents"></div>
</div>
<div class="understanding" style="position: absolute;right: 20px;top: 10px;">100% -</div>
<div class="understanding" style="position: absolute;right: 20px;bottom: 10px;">0% -</div>

<button id="switchview" onclick="switchview()">Summary</button>


<script>
var m = 180; // number of samples per layer
var totalVotingData = [];
var votingData = Array.apply(null, new Array(m)).map(Number.prototype.valueOf,0.5);
var starttime = new Date().getTime();
var currentVotingStatus = 0.5;
var updateInterval = 1000; // milliseconds
var currentView = "realtime";

function updateVotingData() {
  // update data
  //console.log(votingData)
  votingData.shift();
  votingData.push(currentVotingStatus);
  layers0 = stack([votingData.map(
      function(d,i){ 
        return {x:i,y:d};
      })])
  layers1 = stack([totalVotingData]);
  //console.log(votingData)
  // update visualization
  transition();
}

function switchview(){
  if (currentView === 'realtime') {
    currentView = 'alltime';
    $('#switchview').text('Real time');
  }else{
    currentView = 'realtime';
    $('#switchview').text('Summary');
  }
}
  
var wsbase = window.location.hostname+':'+window.location.port+window.location.pathname;
wsbase = wsbase.substr(0,wsbase.lastIndexOf('/'));
var connection = new WebSocket('ws://'+wsbase+'/datafeed');
connection.onopen = function(){
   /* Send a small message to the console once the connection is established */
   console.log('Connection open!');
}
connection.onmessage = function(e){
   var server_message = JSON.parse(e.data); // TODO: catch errors
   var yes = server_message.yes;
   var no = server_message.no;
   var frac = yes/(yes+no)
   if (yes+no === 0) {
      frac = 0.5;
   }
   console.log(server_message,frac);
   document.getElementById('nstudents').innerHTML = yes+no+'';
   currentVotingStatus = frac;
   totalVotingData.push({x:new Date().getTime()-starttime , y:frac});
   console.log(totalVotingData)
   //votingData.push(server_message.yes/(server_message.yes+server_message.no));
   //
}
connection.onclose = function(){
   console.log('Connection closed');
}
connection.onerror = function(error){
   console.log('Error detected: ' + error);
}


var n = 1, // number of layers
    stack = d3.layout.stack().offset("silhouette‏"),
    //layers0 = stack([[{x:0,y:0.1},{x:1,y:0.3},{x:2,y:0.4}]])
    layers0 = stack([votingData.map(
      function(d,i){ 
        return {x:i,y:d};
      })])
    layers1 = stack([totalVotingData]);

// Update voting data periodically
setInterval(updateVotingData,updateInterval);


var width = window.innerWidth,
    height = window.innerHeight;

var x = d3.scale.linear()
    .domain([0, m - 1])
    .range([0, width*(m+1)/m]);


var y,area;

var svg = d3.select("body").append("svg")
    .attr("width", width)
    .attr("height", height);

function setScale(){
  console.log('setting scale')
  width = window.innerWidth;
  height = window.innerHeight;
  y = d3.scale.linear()
    .domain([0, 1])
    .range([height, 0]);
  area = d3.svg.area()
    .interpolate('basis')
    .x(function(d) { return x(d.x); })
    .y0(function(d) { return y(d.y0); })
    .y1(function(d) { return y(d.y0 + d.y); });
    
  svg.attr("width", width)
    .attr("height", height);
}

setScale();

svg.selectAll("path")
    .data(layers0)
  .enter().append("path")
    .attr("d", area)
    .style("fill", '#463335');

function transition() {
  if (currentView === 'realtime'){
    x = d3.scale.linear()
    .domain([0, m - 1])
    .range([0, width*(m+1)/m]);

    d3.selectAll("path")
      .data(layers0)
      .attr("d", area)
      .attr("transform", null)
      .style("fill", '#463335')
    .transition()
      .ease("linear")
      .duration(updateInterval*0.95)
      .attr("transform", "translate(" + x(-1) + ")");
  }else{
    x = d3.scale.linear()
    .domain([d3.min(totalVotingData,function(d){
      return d.x
    }),d3.max(totalVotingData,function(d){
      return d.x
    })])
    .range([0, width]);

    d3.selectAll("path")
      .data(layers1)
      .attr("d", area)
      .style("fill", '#463335')
      .attr("transform", null)
    }
}

</script>

</body>
</html>
