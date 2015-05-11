<html>
<head>

<title>GETit Student</title>

<meta name="viewport" content="width=device-width, user-scalable=no">

<style type="text/css">
    .onoffswitch {
    position: relative;
    width: 200px;
    -webkit-user-select:none; -moz-user-select:none; -ms-user-select: none;
    }
    .onoffswitch-checkbox {
    display: none;
    }
    .onoffswitch-label {
    display: block; overflow: hidden; cursor: pointer;
    border: 2px solid #999999; border-radius: 5px;
    }
    .onoffswitch-inner {
    width: 200%; margin-left: -100%;
    -moz-transition: margin 0.3s ease-in 0s; -webkit-transition: margin 0.3s ease-in 0s;
    -o-transition: margin 0.3s ease-in 0s; transition: margin 0.3s ease-in 0s;
    }
    .onoffswitch-inner:before, .onoffswitch-inner:after {
    float: left; width: 50%; height: 75px; padding: 0; line-height: 75px;
    font-size: 15px; color: white; font-family: Trebuchet, Arial, sans-serif; font-weight: bold;
    -moz-box-sizing: border-box; -webkit-box-sizing: border-box; box-sizing: border-box;
    }
    .onoffswitch-inner:before {
    content: "NOT GET IT";
    padding-left: 10px;
    background-color: #C71E31; color: #F2ECDA;
    }
    .onoffswitch-inner:after {
    content: "GET IT";
    padding-right: 10px;
    background-color: #463335; color: #F2ECDA;
    text-align: right;
    }
    .onoffswitch-switch {
    width: 50px; margin: 0px;
    background: #F2ECDA;
    border: 2px solid #999999; border-radius: 5px;
    position: absolute; top: 0; bottom: 0; right: 146px;
    -moz-transition: all 0.3s ease-in 0s; -webkit-transition: all 0.3s ease-in 0s;
    -o-transition: all 0.3s ease-in 0s; transition: all 0.3s ease-in 0s;
    background-image: -moz-linear-gradient(center top, rgba(0,0,0,0.1) 0%, rgba(0,0,0,0) 100%);
    background-image: -webkit-linear-gradient(center top, rgba(0,0,0,0.1) 0%, rgba(0,0,0,0) 100%);
    background-image: -o-linear-gradient(center top, rgba(0,0,0,0.1) 0%, rgba(0,0,0,0) 100%);
    background-image: linear-gradient(center top, rgba(0,0,0,0.1) 0%, rgba(0,0,0,0) 100%);
    }
    .onoffswitch-checkbox:checked + .onoffswitch-label .onoffswitch-inner {
    margin-left: 0;
    }
    .onoffswitch-checkbox:checked + .onoffswitch-label .onoffswitch-switch {
    right: 0px;
    }

	body,td,th {
		font-family: Arial, Helvetica, sans-serif;
	}
    .container{
        margin-right: auto;
        margin-left: auto;
        width: 200px;
    }
@media (max-width:30.063em) { /* smartphones, iPhone, portrait 480x320 phones */ 
	
}
</style>
</script>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
</head>
<body>
<div class="container">
<p>
    If you don't understand turn the switch to the right
</p>
<p>
    If you understand turn the switch to the left
</p>
<div class="onoffswitch">
	<!-- if student understands, checkbox is unchecked-->
<input type="checkbox" name="onoffswitch" 
	{{if eq .LastAnswer "no"}} checked {{end}}
	onclick="urlRequest()"
	class="onoffswitch-checkbox"
	id="myonoffswitch">
<label class="onoffswitch-label" for="myonoffswitch">
<div class="onoffswitch-inner"></div>
<div class="onoffswitch-switch"></div>
</label>
</div> 
</div>
</style>
	<script type="text/javascript">
	// Resend this URL to the server every few seconds so that they don't think we're lost
		var url = "student?answer={{.LastAnswer}}";
		var INTERVAL = 1000*60*5; // resend every few minutes
		var xmlHttp = null;
		if (window.location.search === '') { // first time, lets tell the server that we're real
			xmlHttp = new XMLHttpRequest();
    		xmlHttp.open( "GET", url, false );
    		xmlHttp.send( null );
		}

		function urlRequest(){
			var url = "student?studentid={{.StudentId}}&answer=";
			if (document.getElementById('myonoffswitch').checked === true) {
				url += "no"
			} else {
				url += "yes"
			}
    		xmlHttp = new XMLHttpRequest();
    		xmlHttp.open( "GET", url, false );
    		xmlHttp.send( null );
    		//console.log(xmlHttp)
    		/*
            d = document.open('text/html');
    		d.write(xmlHttp.responseText)
    		d.close();
            */
		}	
		// after a certain timeout, clear the window and reload everything from the server
		// it helps the server know which clients are alive
		window.setTimeout(urlRequest,INTERVAL);
	</script>
</body>
</html>