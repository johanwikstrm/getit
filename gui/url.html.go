<html>
<head>
	<title>GETit</title>

<link rel="stylesheet" type="text/css" href="getit.css">
<script src="//ajax.googleapis.com/ajax/libs/jquery/1.11.0/jquery.min.js"></script>
<style>
body{
    margin: 0;
	background: url(class_blurred.jpg)  no-repeat ;
	background-size: cover ;
}

.clickable:hover{
	opacity: 0.9;
	cursor: pointer;
}

.qr{
	opacity: 0.5;
	z-index: 2;
}

.startbutton{
	display: block;
	padding: 1em;
	background-color: #DDD;
	text-decoration: none;
	color: #463335;
	border-radius: 2px;
}

.bigQr{
	width: 500px;
	height: 500px;
	position: absolute;
	margin-left: 300px;
	top: 10px;
	opacity: 1;

}


</style>
</head>


<body >
	<div style="border-radius:0px;background-color:#f2ecda;opacity:0.5;position:absolute;right:130px;width:600px;height:683px;"></div>
	<br><br><br><br>
	<div style="border-radius:200px;background-color:#c71e31;opacity:0.8;position:absolute;left:60px;width:250px;height:250px;"></div>
	<div style="text-align:center;position:absolute;left:113px;width:150px;height:400px;font-family:century gothic;font-size:15px">
		<font color="white"><h3>1.</h3>Students go to {{.StudentUrl}}
		<br> OR scan this QR on their mobile device
		<br><br>
			<img id="qr" onclick="toggleQR()" class="clickable qr" 
			width="40px" height="40px" src="https://chart.googleapis.com/chart?cht=qr&chs=400x400&chl={{.StudentUrl}}"></img>
			<br><font size="2px"><i>click to enlarge!</i></font></font>
	</div>			
	
	<br><br><br><br><br>
	<div style="text-align:right;opacity:0.8;position:absolute;right:200px;width:300px;height:100px;font-family:century gothic;font-size:40px">
		Welcome to
	</div>	
	
	<br><br><br><br>
	<img border="0" src="getit-logo_transparent.png" alt="GETit" width="478" height="156" style="position:absolute;right:200px"></img>
	
	<br><br><br><br><br>
	<div style="border-radius:200px;background-color:#463335;opacity:0.8;position:absolute;left:60px;width:250px;height:250px;"></div>
	<div style="text-align:center;position:absolute;left:113px;width:150px;height:200px;font-family:century gothic;font-size:15px">	
		<font color="white"><h3>2.</h3>
		<br><a target="_blank" class="startbutton" href="http://{{.TeacherUrl}}">Teacher clicks here</a></center>
		<br>Then you are good to go!</font>
	</div>	
	
	<br><br><br><br>
	<div style="text-align:justify;opacity:0.8;position:absolute;right:200px;width:478px;height:150px;font-family:century gothic;font-size:25px">	
		GETit allows students to mediate their level of understanding with the lecturer without interrupting and 
		shortens the distance between students <br> and teachers in the classroom.
	</div>
	
	

	<br><br><br><br><br><br><br><br><br><br>
	<script>
		/* The script making the QR code bigger and smaller*/
		var qrSmall = true;
		function toggleQR(){
			var qr = $('#qr');
			if (qrSmall) {
				qr.addClass('bigQr');
			}else{
				qr.removeClass('bigQr');
			}
			qrSmall = !qrSmall;
		}
	</script>
	

	<div style="opacity:0.8;position:absolute;left:10px;width:300px;height:50px;font-family:century gothic;font-size:13px">
		This web app is a project from a course in <a href="kth.se/social-etc-etc">information visualization</a> at the <a href="http://kth.se/en">Royal Institute of Technology</a>.
		For more information, please contact Johan Wikström, <a href="mailto:jwikst@kth.se">jwikst@kth.se</a> or Ida Renström, <a href="idare@kth.se">idare@kth.se</a>
	</div>	
		
	</body>
</html>