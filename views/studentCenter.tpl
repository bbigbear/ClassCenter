<!DOCTYPE html>
<html>
<head>
	{{template "header"}}
	<script src="http://cdn.static.runoob.com/libs/jquery/2.1.1/jquery.min.js"></script>
	<script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
</head>
<body>
<div class="navbar navbar-default" >
	<div class="container">
		{{template "navbar".}}
	</div>
</div>
<div class="container">
	<div class="row">
		<div class="col-sm-6">
		<div class="panel panel-primary" >
			<div class="panel-heading">
				<h3 class="panel-title">通知</h3>
			</div>
			<div class="panel-body" style="height:200px">
				2018年选课通知
			</div>
		</div>
		</div>
		<div class="col-sm-6">
		<div class="panel panel-warning">
			<div class="panel-heading">
				<h3 class="panel-title">重要消息</h3>
			</div>
			<div class="panel-body" style="height:200px">消息列表	
			</div>
		</div>
		</div>
		<div class="col-sm-6">
		<div class="panel panel-success">
			<div class="panel-heading">
				<h3 class="panel-title">课表</h3>
			</div>
			<div class="panel-body">我的课表	
			</div>
		</div>
		</div>	
		<div class="col-sm-6">
		<div class="panel panel-info">
			<div class="panel-heading">
				<h3 class="panel-title">成绩</h3>
			</div>
			<div class="panel-body">我的成绩
			</div>
		</div>
		</div>		
	</div>
</div>
</body>
</html>
